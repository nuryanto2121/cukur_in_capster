package useauth

import (
	"context"
	"errors"
	iauth "nuryanto2121/cukur_in_capster/interface/auth"
	ifileupload "nuryanto2121/cukur_in_capster/interface/fileupload"
	iusers "nuryanto2121/cukur_in_capster/interface/user"
	"nuryanto2121/cukur_in_capster/models"
	"nuryanto2121/cukur_in_capster/pkg/setting"
	util "nuryanto2121/cukur_in_capster/pkg/utils"
	"nuryanto2121/cukur_in_capster/redisdb"
	useemailauth "nuryanto2121/cukur_in_capster/usecase/email_auth"
	"strconv"
	"time"
)

type useAuht struct {
	repoAuth       iusers.Repository
	repoFile       ifileupload.Repository
	contextTimeOut time.Duration
}

func NewUserAuth(a iusers.Repository, b ifileupload.Repository, timeout time.Duration) iauth.Usecase {
	return &useAuht{repoAuth: a, repoFile: b, contextTimeOut: timeout}
}

func (u *useAuht) Logout(ctx context.Context, Claims util.Claims, Token string) (err error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeOut)
	defer cancel()

	redisdb.TurncateList(ctx, Token)
	redisdb.TurncateList(ctx, Claims.CapsterID+"_fcm")

	return nil
}

func (u *useAuht) Login(ctx context.Context, dataLogin *models.LoginForm) (output interface{}, err error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeOut)
	defer cancel()

	var (
		DataOwner        = models.SsUser{}
		DataCapster      = models.LoginCapster{}
		isBarber    bool = true
		response         = map[string]interface{}{}
		expireToken      = setting.FileConfigSetting.JWTExpired
		canChange   bool = false
	)

	if dataLogin.Type == "owner" {
		DataOwner, err = u.repoAuth.GetByAccount(dataLogin.Account, false) //u.repoUser.GetByEmailSaUser(dataLogin.UserName)
		if DataOwner.UserType == "" && err == models.ErrNotFound {
			return nil, errors.New("Email anda belum terdaftar.")
		} else {
			if DataOwner.UserType == "capster" {
				DataCapster, err = u.repoAuth.GetByCapster(dataLogin.Account)
				if err != nil {
					return nil, errors.New("Email anda belum terdaftar.")
				}
				if !DataCapster.IsActive {
					return nil, errors.New("Account anda belum aktif. Silahkan hubungi pemilik Barber")
				}

				if DataCapster.BarberID == 0 {
					return nil, errors.New("Anda belum terhubung dengan Barber, Silahkan hubungi pemilik Barber")
				}

				if !DataCapster.BarberIsActive {
					return nil, errors.New("Saat ini barber anda sedang tidak aktif.")
				}
				isBarber = false
			} else {
				DataCapster, err = u.repoAuth.GetByCapster(dataLogin.Account)
				if DataCapster.Email != "" && DataCapster.Email == dataLogin.Account {
					canChange = true
				}

			}

			if !DataOwner.IsActive {
				return nil, errors.New("Account andan belum aktif. Silahkan hubungi pemilik Barber")
			}
		}
	} else {
		isBarber = false
		DataCapster, err = u.repoAuth.GetByCapster(dataLogin.Account)
		if err != nil {
			return nil, errors.New("Email anda belum terdaftar.")
		}
		if !DataCapster.IsActive {
			return nil, errors.New("Account anda belum aktif. Silahkan hubungi pemilik Barber")
		}

		if DataCapster.BarberID == 0 {
			return nil, errors.New("Anda belum terhubung dengan Barber, Silahkan hubungi pemilik Barber")
		}

		if !DataCapster.BarberIsActive {
			return nil, errors.New("Saat ini barber anda sedang tidak aktif.")
		}

		DataOwner, err = u.repoAuth.GetByAccount(dataLogin.Account, true)
		if DataOwner.Email != "" && DataOwner.Email == dataLogin.Account {
			canChange = true
		}

	}

	if isBarber {
		if !util.ComparePassword(DataOwner.Password, util.GetPassword(dataLogin.Password)) {
			return nil, errors.New("Password yang anda masukkan salah. Silahkan coba lagi.")
		}
		DataFile, err := u.repoFile.GetBySaFileUpload(ctx, DataOwner.FileID)

		token, err := util.GenerateTokenBarber(DataOwner.UserID, dataLogin.Account, DataOwner.UserType)
		if err != nil {
			return nil, err
		}

		redisdb.AddSession(ctx, token, DataOwner.UserID, time.Duration(expireToken)*time.Hour)

		if dataLogin.FcmToken != "" {
			redisdb.AddSession(ctx, strconv.Itoa(DataOwner.UserID)+"_fcm", dataLogin.FcmToken, time.Duration(expireToken)*time.Hour)
		}

		restUser := map[string]interface{}{
			"owner_id":   DataOwner.UserID,
			"owner_name": DataOwner.Name,
			"email":      DataOwner.Email,
			"telp":       DataOwner.Telp,
			"file_id":    DataOwner.FileID,
			"file_name":  DataFile.FileName,
			"file_path":  DataFile.FilePath,
		}
		response = map[string]interface{}{
			"token":      token,
			"data_owner": restUser,
			"user_type":  "barber",
			"can_change": canChange,
		}

	} else {
		if !util.ComparePassword(DataCapster.Password, util.GetPassword(dataLogin.Password)) {
			return nil, errors.New("Password yang anda masukkan salah. Silahkan coba lagi.")
		}

		token, err := util.GenerateToken(DataCapster.CapsterID, DataCapster.OwnerID, DataCapster.BarberID)
		if err != nil {
			return nil, err
		}
		redisdb.AddSession(ctx, token, DataCapster.CapsterID, time.Duration(expireToken)*time.Hour)

		if dataLogin.FcmToken != "" {
			redisdb.AddSession(ctx, strconv.Itoa(DataCapster.CapsterID)+"_fcm", dataLogin.FcmToken, time.Duration(expireToken)*time.Hour)
		}

		restUser := map[string]interface{}{
			"owner_id":     DataCapster.OwnerID,
			"owner_name":   DataCapster.OwnerName,
			"barber_id":    DataCapster.BarberID,
			"barber_name":  DataCapster.BarberName,
			"capster_id":   DataCapster.CapsterID,
			"email":        DataCapster.Email,
			"telp":         DataCapster.Telp,
			"capster_name": DataCapster.CapsterName,
			"file_id":      DataCapster.FileID,
			"file_name":    DataCapster.FileName,
			"file_path":    DataCapster.FilePath,
		}
		response = map[string]interface{}{
			"token":        token,
			"data_capster": restUser,
			"user_type":    "capster",
			"can_change":   canChange,
		}
	}

	return response, nil
}

func (u *useAuht) ForgotPassword(ctx context.Context, dataForgot *models.ForgotForm) (err error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeOut)
	defer cancel()

	DataCapster, err := u.repoAuth.GetByCapster(dataForgot.Account) //u.repoUser.GetByEmailSaUser(dataLogin.UserName)
	if err != nil {
		// return util.GoutputErrCode(http.StatusUnauthorized, "Your User/Email not valid.") //appE.ResponseError(util.GetStatusCode(err), fmt.Sprintf("%v", err), nil)
		return errors.New("Your Account not valid.")
	}
	if DataCapster.CapsterName == "" {
		return errors.New("Your Account not valid.")
	}

	return nil
}

func (u *useAuht) ResetPassword(ctx context.Context, dataReset *models.ResetPasswd) (err error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeOut)
	defer cancel()

	if dataReset.Passwd != dataReset.ConfirmPasswd {
		return errors.New("Password dan confirm password harus sama.")
	}

	DataCapster, err := u.repoAuth.GetByCapster(dataReset.Account)
	if err != nil {
		return err
	}

	DataUser, err := u.repoAuth.GetDataBy(DataCapster.CapsterID)
	if err != nil {
		return err
	}
	DataUser.Password, _ = util.Hash(dataReset.Passwd)
	// email, err := util.ParseEmailToken(dataReset.TokenEmail)
	// if err != nil {
	// 	email = dataReset.TokenEmail
	// }

	// dataUser, err := u.repoUser.GetByEmailSaUser(email)

	// dataUser.Password, _ = util.Hash(dataReset.Passwd)

	err = u.repoAuth.Update(DataUser.UserID, &DataUser)
	if err != nil {
		return err
	}

	return nil
}

func (u *useAuht) Register(ctx context.Context, dataRegister models.RegisterForm) (output interface{}, err error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeOut)
	defer cancel()

	var User models.SsUser

	User.Name = dataRegister.Name
	if dataRegister.Passwd != dataRegister.ConfirmPasswd {
		return output, errors.New("Password dan confirm password harus sama.")
	}
	User.Password, _ = util.Hash(dataRegister.Passwd)
	User.UserType = dataRegister.UserType
	User.UserEdit = dataRegister.Name
	User.UserInput = dataRegister.Name

	//check email or telp
	if util.CheckEmail(dataRegister.Account) {
		User.Email = dataRegister.Account
	} else {
		User.Telp = dataRegister.Account
	}
	err = u.repoAuth.Create(&User)
	if err != nil {
		return output, err
	}

	GenCode := util.GenerateNumber(4)

	// send generate code
	mailService := &useemailauth.Register{
		Email:      User.Email,
		Name:       User.Name,
		GenerateNo: GenCode,
	}

	err = mailService.SendRegister()
	if err != nil {
		return output, err
	}

	//store to redis
	err = redisdb.AddSession(ctx, dataRegister.Account, GenCode, 24*time.Hour)
	if err != nil {
		return output, err
	}
	out := map[string]interface{}{
		"gen_code": GenCode,
	}
	return out, nil
}

func (u *useAuht) Verify(ctx context.Context, dataVeriry models.VerifyForm) (err error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeOut)
	defer cancel()
	data := redisdb.GetSession(ctx, dataVeriry.Account)
	if data == "" {
		return errors.New("Please Resend Code")
	}

	if data != dataVeriry.VerifyCode {
		return errors.New("Invalid Code.")
	}

	return nil
}

// PathFCM implements iauth.Usecase.
func (u *useAuht) PathFCM(ctx context.Context, dataPatch models.PathFCM) (err error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeOut)
	defer cancel()

	//check active user
	CapsterID, _ := strconv.Atoi(dataPatch.CapsterID)
	_, err = u.repoAuth.GetDataBy(CapsterID)
	if err != nil {
		return err
	}

	if dataPatch.FcmToken != "" {
		redisdb.AddSession(ctx, dataPatch.CapsterID+"_fcm", dataPatch.FcmToken, time.Duration(setting.FileConfigSetting.JWTExpired)*time.Hour)
	}
	return nil
}
