package reposysusers

import (
	"fmt"
	iusers "nuryanto2121/cukur_in_capster/interface/user"
	"nuryanto2121/cukur_in_capster/models"
	"nuryanto2121/cukur_in_capster/pkg/logging"
	"nuryanto2121/cukur_in_capster/pkg/setting"

	"github.com/jinzhu/gorm"
	"github.com/labstack/gommon/log"
)

type repoSysUser struct {
	Conn *gorm.DB
}

func NewRepoSysUser(Conn *gorm.DB) iusers.Repository {
	return &repoSysUser{Conn}
}
func (db *repoSysUser) GetByAccount(Account string, IsOwner bool) (result models.SsUser, err error) {
	var query *gorm.DB

	if IsOwner {
		query = db.Conn.Where("(email iLIKE ? OR telp = ?) AND user_type = 'owner' ", Account, Account).First(&result)
	} else {
		query = db.Conn.Where("(email iLIKE ? OR telp = ?)", Account, Account).First(&result)
	}

	log.Info(fmt.Sprintf("%v", query.QueryExpr()))
	// logger.Query(fmt.Sprintf("%v", query.QueryExpr()))
	err = query.Error

	if err != nil {
		//
		if err == gorm.ErrRecordNotFound {
			return result, models.ErrNotFound
		}
		return result, err
	}

	return result, err
}
func (db *repoSysUser) GetByCapster(Account string) (result models.LoginCapster, err error) {

	query := db.Conn.Table("ss_user su").Select(`su.user_id as capster_id, su."name" as capster_name,su."password",su.email ,
	su.telp ,su.file_id ,sf.file_name ,sf.file_path ,b.barber_id ,b.barber_name,su.user_input as owner_id ,so."name" as owner_name,su.is_active,
	b.is_active as barber_is_active
	`).Joins(`
	left join barber_capster bc on su.user_id = bc.capster_id`).Joins(`
	left join barber b on b.barber_id = bc.barber_id `).Joins(`
	left join sa_file_upload sf on sf.file_id =su.file_id`).Joins(`
	left join ss_user so on so.user_id::varchar = su.user_input `).Where(`
	su.user_type = 'capster' AND (su.email iLike ? OR su.telp = ?)`, Account, Account).First(&result)

	// query := db.Conn.Where("email = ?", Account).Or("telp = ?", Account).First(&result)
	// query := db.Conn.Table("ss_user su").Select(`su.user_id as capster_id, su."name" as capster_name,su."password",su.email ,
	// 					su.telp ,su.file_id ,sf.file_name ,sf.file_path ,b.barber_id ,b.barber_name,b.owner_id ,so."name" as owner_name`).Joins(`
	// 					inner join barber_capster bc on su.user_id = bc.capster_id`).Joins(`
	// 					inner join barber b on b.barber_id = bc.barber_id `).Joins(`
	// 					left join sa_file_upload sf on sf.file_id =su.file_id`).Joins(`
	// 					inner join ss_user so on so.user_id = b.owner_id `).Where(`
	// 					su.user_name = ?`, Account).First(&result)
	log.Info(fmt.Sprintf("%v", query.QueryExpr()))
	// logger.Query(fmt.Sprintf("%v", query.QueryExpr()))
	err = query.Error

	if err != nil {
		//
		if err == gorm.ErrRecordNotFound {
			return result, models.ErrNotFound
		}
		return result, err
	}

	return result, err
}

func (db *repoSysUser) GetDataBy(ID int) (result *models.SsUser, err error) {
	var sysUser = &models.SsUser{}
	query := db.Conn.Where("user_id = ? ", ID).Find(sysUser)
	log.Info(fmt.Sprintf("%v", query.QueryExpr()))
	err = query.Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, models.ErrNotFound
		}
		return nil, err
	}
	return sysUser, nil
}

func (db *repoSysUser) GetList(queryparam models.ParamList) (result []*models.SsUser, err error) {

	var (
		pageNum  = 0
		pageSize = setting.FileConfigSetting.App.PageSize
		sWhere   = ""
		// logger   = logging.Logger{}
		orderBy = queryparam.SortField
	)
	// pagination
	if queryparam.Page > 0 {
		pageNum = (queryparam.Page - 1) * queryparam.PerPage
	}
	if queryparam.PerPage > 0 {
		pageSize = queryparam.PerPage
	}
	//end pagination

	// Order
	if queryparam.SortField != "" {
		orderBy = queryparam.SortField
	}
	//end Order by

	// WHERE
	if queryparam.InitSearch != "" {
		sWhere = queryparam.InitSearch
	}

	if queryparam.Search != "" {
		if sWhere != "" {
			sWhere += " and " + queryparam.Search
		} else {
			sWhere += queryparam.Search
		}
	}

	// end where
	if pageNum >= 0 && pageSize > 0 {
		query := db.Conn.Where(sWhere).Offset(pageNum).Limit(pageSize).Order(orderBy).Find(&result)
		fmt.Printf("%v", query.QueryExpr()) //cath to log query string
		err = query.Error
	} else {
		query := db.Conn.Where(sWhere).Order(orderBy).Find(&result)
		fmt.Printf("%v", query.QueryExpr()) //cath to log query string
		err = query.Error
	}

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, err
		}
		return nil, err
	}
	return result, nil
}
func (db *repoSysUser) Create(data *models.SsUser) error {
	var (
		logger = logging.Logger{}
		err    error
	)
	query := db.Conn.Create(data)
	logger.Query(fmt.Sprintf("%v", query.QueryExpr())) //cath to log query string
	err = query.Error
	if err != nil {
		return err
	}
	return nil
}
func (db *repoSysUser) Update(ID int, data interface{}) error {
	var (
		logger = logging.Logger{}
		err    error
	)
	query := db.Conn.Model(models.SsUser{}).Where("user_id = ?", ID).Updates(data)
	logger.Query(fmt.Sprintf("%v", query.QueryExpr())) //cath to log query string
	err = query.Error
	if err != nil {
		return err
	}
	return nil
}
func (db *repoSysUser) Delete(ID int) error {
	var (
		logger = logging.Logger{}
		err    error
	)
	query := db.Conn.Where("user_id = ?", ID).Delete(&models.SsUser{})
	logger.Query(fmt.Sprintf("%v", query.QueryExpr())) //cath to log query string
	err = query.Error
	if err != nil {
		return err
	}
	return nil
}
func (db *repoSysUser) Count(queryparam models.ParamList) (result int, err error) {
	var (
		sWhere = ""
		logger = logging.Logger{}
	)
	result = 0

	// WHERE
	if queryparam.InitSearch != "" {
		sWhere = queryparam.InitSearch
	}

	if queryparam.Search != "" {
		if sWhere != "" {
			sWhere += " and " + queryparam.Search
		}
	}
	// end where

	query := db.Conn.Model(&models.SsUser{}).Where(sWhere).Count(&result)
	logger.Query(fmt.Sprintf("%v", query.QueryExpr())) //cath to log query string
	err = query.Error
	if err != nil {
		return 0, err
	}

	return result, nil
}
