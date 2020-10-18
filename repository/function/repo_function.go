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