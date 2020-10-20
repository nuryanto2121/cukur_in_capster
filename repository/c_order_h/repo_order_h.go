package repoorderh

import (
	"fmt"
	iorder_h "nuryanto2121/cukur_in_capster/interface/c_order_h"
	"nuryanto2121/cukur_in_capster/models"
	"nuryanto2121/cukur_in_capster/pkg/logging"
	"nuryanto2121/cukur_in_capster/pkg/setting"
	"strings"

	"github.com/jinzhu/gorm"
)

type repoOrderH struct {
	Conn *gorm.DB
}

func NewRepoOrderH(Conn *gorm.DB) iorder_h.Repository {
	return &repoOrderH{Conn}
}

func (db *repoOrderH) GetDataBy(ID int) (result *models.OrderH, err error) {
	var (
		logger = logging.Logger{}
		mOrder = &models.OrderH{}
	)

	query := db.Conn.Where("order_id = ? ", ID).Find(mOrder)
	logger.Query(fmt.Sprintf("%v", query.QueryExpr()))
	err = query.Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return result, models.ErrNotFound
		}
		return result, err
	}
	return mOrder, nil
}
func (db *repoOrderH) GetList(queryparam models.ParamList) (result []*models.OrderList, err error) {

	var (
		pageNum  = 0
		pageSize = setting.FileConfigSetting.App.PageSize
		sWhere   = ""
		logger   = logging.Logger{}
		orderBy  = queryparam.SortField
		query    *gorm.DB
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
			sWhere += " and lower(customer_name) LIKE ?" //+ queryparam.Search
		} else {
			sWhere += "lower(customer_name) LIKE ?" //queryparam.Search
		}
		query = db.Conn.Table("v_order_h").Select(`
		*
		`).Where(sWhere, queryparam.Search).Offset(pageNum).Limit(pageSize).Order(orderBy).Find(&result)
	} else {
		query = db.Conn.Table("v_order_h").Select(`
		*
		`).Where(sWhere).Offset(pageNum).Limit(pageSize).Order(orderBy).Find(&result)
	}

	// end where

	// query := db.Conn.Where(sWhere).Offset(pageNum).Limit(pageSize).Order(orderBy).Find(&result)

	logger.Query(fmt.Sprintf("%v", query.QueryExpr())) //cath to log query string
	err = query.Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, err
		}
		return nil, err
	}
	return result, nil
}
func (db *repoOrderH) SumPriceDetail(queryparam models.ParamList) (result float32, err error) {
	type Results struct {
		Price float32 `json:"price"`
	}
	var (
		sWhere = ""
		logger = logging.Logger{}
		op     = &Results{}
		query  *gorm.DB
	)

	result = 0

	// WHERE
	if queryparam.InitSearch != "" {
		sWhere = queryparam.InitSearch
	}

	sWhere = strings.ReplaceAll(sWhere, "barber_id", "v_order_h.barber_id")
	if queryparam.Search != "" {
		if sWhere != "" {
			sWhere += " and lower(customer_name) LIKE ?" //+ queryparam.Search
		} else {
			sWhere += "lower(customer_name) LIKE ?" //queryparam.Search
		}
		query = db.Conn.Table("v_order_h").Select(`
		coalesce(sum(price ),0) as price
		`).Where(sWhere, queryparam.Search).First(&op)
	} else {
		query = db.Conn.Table("v_order_h").Select(`
		coalesce(sum(price ),0) as price
		`).Where(sWhere).First(&op)
	}

	logger.Query(fmt.Sprintf("%v", query.QueryExpr())) //cath to log query string
	err = query.Error
	if err != nil {
		return 0, err
	}

	return op.Price, nil
}
func (db *repoOrderH) Create(data *models.OrderH) error {
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
func (db *repoOrderH) Update(ID int, data map[string]interface{}) error {
	var (
		logger = logging.Logger{}
		err    error
	)
	query := db.Conn.Model(models.OrderH{}).Where("order_id = ?", ID).Updates(data)
	logger.Query(fmt.Sprintf("%v", query.QueryExpr())) //cath to log query string
	err = query.Error
	if err != nil {
		return err
	}
	return nil
}
func (db *repoOrderH) Delete(ID int) error {
	var (
		logger = logging.Logger{}
		err    error
	)
	// query := db.Conn.Where("order_id = ?", ID).Delete(&models.OrderH{})
	query := db.Conn.Exec("Delete From order_h WHERE order_id = ?", ID)
	logger.Query(fmt.Sprintf("%v", query.QueryExpr())) //cath to log query string
	err = query.Error
	if err != nil {
		return err
	}
	return nil
}
func (db *repoOrderH) Count(queryparam models.ParamList) (result int, err error) {
	var (
		sWhere = ""
		logger = logging.Logger{}
		query  *gorm.DB
	)
	result = 0

	// WHERE
	if queryparam.InitSearch != "" {
		sWhere = queryparam.InitSearch
	}

	if queryparam.Search != "" {
		if sWhere != "" {
			sWhere += " and lower(customer_name) LIKE ?" //+ queryparam.Search
		} else {
			sWhere += "lower(customer_name) LIKE ?" //queryparam.Search
		}

		query = db.Conn.Table("v_order_h").Select(`
		*
		`).Where(sWhere, queryparam.Search).Count(&result)
	} else {
		query = db.Conn.Table("v_order_h").Select(`
		*
		`).Where(sWhere).Count(&result)
	}
	// end where

	logger.Query(fmt.Sprintf("%v", query.QueryExpr())) //cath to log query string
	err = query.Error
	if err != nil {
		return 0, err
	}

	return result, nil
}
