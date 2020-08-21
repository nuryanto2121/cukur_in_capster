package useauth

import (
	"context"
	"errors"
	iauth "nuryanto2121/cukur_in_capster/interface/auth"
	ifileupload "nuryanto2121/cukur_in_capster/interface/fileupload"
	iusers "nuryanto2121/cukur_in_capster/interface/user"
	"nuryanto2121/cukur_in_capster/models"
	util "nuryanto2121/cukur_in_capster/pkg/utils"
	"nuryanto2121/cukur_in_capster/redisdb"
	useemailauth "nuryanto2121/cukur_in_capster/usecase/email_auth"
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
func (u *useAuht) Login(ctx context.Context, dataLogin *models.LoginForm) (output interface{}, err error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeOut)
	defer cancel()

	DataCapster, err := u.repoAuth.GetByAccount(dataLogin.Account) //u.repoUser.GetByEmailSaUser(dataLogin.UserName)
	if err != nil {
		// return util.GoutputErrCode(http.StatusUnauthorized, "Your User/Email not valid.") //appE.ResponseError(util.GetStatusCode(err), fmt.Sprintf("%v", err), nil)
		return nil, errors.New("Your Account not valid.")
	}

	if !util.ComparePassword(DataCapster.Password, util.GetPassword(dataLogin.Password)) {
		return nil, errors.New("Your Password not valid.")
	}
	// DataFile, err := u.repoFile.GetBySaFileUpload(ctx, DataCapster.FileID)

	token, err := util.GenerateToken(DataCapster.CapsterID, DataCapster.OwnerID, DataCapster.BarberID)
	if err != nil {
		return nil, err
	}

	redisdb.AddSession(token, DataCapster.CapsterID, 0)

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
	response := map[string]interface{}{
		"token":        token,
		"data_capster": restUser,
	}

	return response, nil
}

func (u *useAuht) ForgotPassword(ctx context.Context, dataForgot *models.ForgotForm) (err error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeOut)
	defer cancel()

	DataCapster, err := u.repoAuth.GetByAccount(dataForgot.Account) //u.repoUser.GetByEmailSaUser(dataLogin.UserName)
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
		return errors.New("Password and Confirm Password not same.")
	}

	DataCapster, err := u.repoAuth.GetByAccount(dataReset.Account)
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
		return output, errors.New("Password and Confirm Password Not Valid")
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
	err = redisdb.AddSession(dataRegister.Account, GenCode, 2)
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
	data := redisdb.GetSession(dataVeriry.Account)
	if data == "" {
		return errors.New("Please Resend Code")
	}

	if data != dataVeriry.VerifyCode {
		return errors.New("Invalid Code.")
	}

	return nil
}
