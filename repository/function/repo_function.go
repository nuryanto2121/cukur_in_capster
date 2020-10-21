package repofunction

import (
	"fmt"
	"nuryanto2121/cukur_in_capster/models"
	"nuryanto2121/cukur_in_capster/pkg/logging"
	"nuryanto2121/cukur_in_capster/pkg/postgresdb"
	util "nuryanto2121/cukur_in_capster/pkg/utils"
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
)

type FN struct {
	Claims util.Claims
}

func (fn *FN) GenTransactionNo(BarberCd string) (string, error) {
	var (
		result string
		conn   *gorm.DB
		logger = logging.Logger{}
		mSeqNo = &models.SsSequenceNo{}
		t      = time.Now()
		year   = t.Year()
		month  = int(t.Month())
		sYear  = strconv.Itoa(year)
		sMonth = strconv.Itoa(month)
		// abjad  = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	)
	if len(sMonth) == 1 {
		sMonth = fmt.Sprintf("0%s", sMonth)
	}
	if len(sYear) == 4 {
		sYear = sYear[len(sYear)-2:]
	}
	pref := fmt.Sprintf("%s%s", sMonth, sYear)
	conn = postgresdb.Conn

	// fmt.Printf("%v", prefixArr)
	// ss := prefixArr[0]
	// query := conn.Table("barber").Select("max(barber_cd)") //
	query := conn.Where("sequence_cd = ?", BarberCd).Find(mSeqNo)
	logger.Query(fmt.Sprintf("%v", query.QueryExpr()))
	err := query.Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			mSeqNo.Prefix = fmt.Sprintf("%s", pref)
			mSeqNo.SeqNo = 1
			mSeqNo.SequenceCd = BarberCd
			mSeqNo.UserInput = fn.Claims.CapsterID
			mSeqNo.UserEdit = fn.Claims.CapsterID
			queryC := conn.Create(mSeqNo)
			logger.Query(fmt.Sprintf("%v", queryC.QueryExpr()))
			err = queryC.Error
			if err != nil {
				return "", err
			}
			result = fmt.Sprintf("%s/%s/0001", BarberCd, pref)
			return result, nil
		}
		return "", err
	}
	seq_no := ""

	if mSeqNo.Prefix == pref {
		mSeqNo.SeqNo += 1
		seq_no = strconv.Itoa(10000 + mSeqNo.SeqNo)
		seq_no = seq_no[len(seq_no)-4:]
	} else {
		mSeqNo.Prefix = fmt.Sprintf("%s", pref)
		mSeqNo.SeqNo = 1
		seq_no = "0001"
	}
	query = conn.Model(models.SsSequenceNo{}).Where("sequence_id = ?", mSeqNo.SequenceID).Updates(mSeqNo)
	logger.Query(fmt.Sprintf("%v", query.QueryExpr()))
	err = query.Error
	if err != nil {
		return "", err
	}
	result = fmt.Sprintf("%s/%s/%s", BarberCd, pref, seq_no)

	return result, nil
}
func (fn *FN) GetBarberData() (result *models.Barber, err error) {
	var (
		logger  = logging.Logger{}
		mBarber = &models.Barber{}
		conn    *gorm.DB
	)
	conn = postgresdb.Conn
	query := conn.Where("barber_id = ? ", fn.Claims.BarberID).Find(mBarber)
	logger.Query(fmt.Sprintf("%v", query.QueryExpr()))
	err = query.Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, models.ErrNotFound
		}
		return nil, err
	}
	return mBarber, nil
}

func (fn *FN) GetCapsterData() (result *models.SsUser, err error) {
	var (
		logger   = logging.Logger{}
		mCapster = &models.SsUser{}
		conn     *gorm.DB
	)
	conn = postgresdb.Conn
	query := conn.Where("user_id = ? ", fn.Claims.CapsterID).Find(mCapster)
	logger.Query(fmt.Sprintf("%v", query.QueryExpr()))
	err = query.Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, models.ErrNotFound
		}
		return nil, err
	}
	return mCapster, nil
}

func (fn *FN) InTimeActiveBarber(data *models.Barber, orderDate time.Time) bool {
	var (
		P = fmt.Println
	)

	// timeStart := data.OperationStart.Date(timeOrder.Year(), timeOrder.Month(), timeOrder.Day())
	timeStart := time.Date(orderDate.Year(), orderDate.Month(), orderDate.Day(), data.OperationStart.Hour(),
		data.OperationStart.Minute(), data.OperationStart.Second(), data.OperationStart.Nanosecond(), data.OperationStart.Local().Location())

	timeEnd := time.Date(orderDate.Year(), orderDate.Month(), orderDate.Day(), data.OperationEnd.Hour(),
		data.OperationEnd.Minute(), data.OperationEnd.Second(), data.OperationEnd.Nanosecond(), data.OperationEnd.Local().Location())

	timeOrder := time.Date(orderDate.Year(), orderDate.Month(), orderDate.Day(), orderDate.Hour(),
		orderDate.Minute(), orderDate.Second(), orderDate.Nanosecond(), orderDate.Local().Location())

	P(timeStart)
	P(timeEnd)
	P(timeOrder)

	if timeOrder.Before(timeEnd) && timeOrder.After(timeStart) {
		return true
	} else {
		return false
	}

}

func (fn *FN) GetCountTrxProses() int {
	var (
		logger = logging.Logger{}
		result = 0
		conn   *gorm.DB
	)
	orderDate := time.Now()

	conn = postgresdb.Conn
	query := conn.Table("order_h").Select(`*`).Where(`
			status = 'P' AND order_date::date = now()::date AND barber_id = ? AND capster_id = ? 
		`, fn.Claims.BarberID, fn.Claims.CapsterID).Count(&result)
	logger.Query(fmt.Sprintf("%v", query.QueryExpr())) //cath to log query string
	err := query.Error
	if err != nil {
		if err == models.ErrNotFound {
			return 0
		}
		logger.Error(err)
	}
	return result
}
