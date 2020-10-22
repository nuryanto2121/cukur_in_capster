package useorder

import (
	"context"
	"errors"
	"fmt"
	"math"
	repofunction "nuryanto2121/cukur_in_capster/repository/function"
	"strings"

	iorderd "nuryanto2121/cukur_in_capster/interface/c_order_d"
	iorderh "nuryanto2121/cukur_in_capster/interface/c_order_h"
	"nuryanto2121/cukur_in_capster/models"
	util "nuryanto2121/cukur_in_capster/pkg/utils"
	"strconv"
	"time"

	"github.com/mitchellh/mapstructure"
)

type useOrder struct {
	repoOrderH     iorderh.Repository
	repoOrderD     iorderd.Repository
	contextTimeOut time.Duration
}

func NewUserMOrder(a iorderh.Repository, b iorderd.Repository, timeout time.Duration) iorderh.Usecase {
	return &useOrder{
		repoOrderH:     a,
		repoOrderD:     b,
		contextTimeOut: timeout}
}

func (u *useOrder) GetDataBy(ctx context.Context, Claims util.Claims, ID int) (interface{}, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeOut)
	defer cancel()

	dataHeader, err := u.repoOrderH.GetDataBy(ID)
	if err != nil {
		return nil, err
	}

	dataDetail, err := u.repoOrderD.GetDataBy(ID)
	if err != nil {
		return nil, err
	}
	fn := &repofunction.FN{
		Claims: Claims,
	}
	dataBarber, err := fn.GetBarberData()
	if err != nil {
		return nil, err
	}
	var total_price float32 = 0

	for _, dataDetail := range dataDetail {
		total_price += dataDetail.Price
	}

	response := map[string]interface{}{
		"from_apps":     dataHeader.FromApps,
		"barber_name":   dataBarber.BarberName,
		"customer_name": dataHeader.CustomerName,
		"email":         dataHeader.Email,
		"telp":          dataHeader.Telp,
		"order_date":    dataHeader.OrderDate,
		"status":        dataHeader.Status,
		"order_id":      dataHeader.OrderID,
		"data_detail":   dataDetail,
		"total_price":   total_price,
	}

	return response, nil
}
func (u *useOrder) GetList(ctx context.Context, Claims util.Claims, queryparam models.ParamList) (result models.ResponseModelList, err error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeOut)
	defer cancel()

	if queryparam.Search != "" {
		queryparam.Search = strings.ToLower(fmt.Sprintf("%%%s%%", queryparam.Search))
	}

	if queryparam.InitSearch != "" {
		queryparam.InitSearch += fmt.Sprintf(" AND capster_id = %s", Claims.CapsterID)
	} else {
		queryparam.InitSearch = fmt.Sprintf("capster_id = %s", Claims.CapsterID)
	}
	result.Data, err = u.repoOrderH.GetList(queryparam)
	if err != nil {
		return result, err
	}

	result.Total, err = u.repoOrderH.Count(queryparam)
	if err != nil {
		return result, err
	}

	// d := float64(result.Total) / float64(queryparam.PerPage)
	result.LastPage = int(math.Ceil(float64(result.Total) / float64(queryparam.PerPage)))
	result.Page = queryparam.Page

	return result, nil
}
func (u *useOrder) GetSumPrice(ctx context.Context, Claims util.Claims, queryparam models.ParamList) (result float32, err error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeOut)
	defer cancel()

	if queryparam.Search != "" {
		queryparam.Search = strings.ToLower(fmt.Sprintf("%%%s%%", queryparam.Search))
	}

	if queryparam.InitSearch != "" {
		queryparam.InitSearch += fmt.Sprintf(" AND status ='F' AND capster_id = %s", Claims.CapsterID)
	} else {
		queryparam.InitSearch = fmt.Sprintf(" status ='F' AND capster_id = %s", Claims.CapsterID)
	}
	result, err = u.repoOrderH.SumPriceDetail(queryparam)
	if err != nil {
		return result, err
	}
	return result, nil
}
func (u *useOrder) Create(ctx context.Context, Claims util.Claims, data *models.OrderPost) (err error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeOut)
	defer cancel()
	var (
		mOrder models.OrderH
	)
	// mapping to struct model saRole
	err = mapstructure.Decode(data, &mOrder)
	if err != nil {
		return err
	}
	fn := &repofunction.FN{
		Claims: Claims,
	}
	dataBarber, err := fn.GetBarberData()
	if err != nil {
		if err == models.ErrNotFound {
			return errors.New("Data Barber tidak ditemukan, Hubungi pemilik barber")
		}
		return err
	}

	dataCapster, err := fn.GetCapsterData()
	if err != nil {
		if err == models.ErrNotFound {
			return errors.New("Data Capster tidak ditemukan, Hubungi pemilik barber")
		}
		return err
	}

	if !dataCapster.IsActive {
		return errors.New("Tidak bisa transaksi, Akun anda tidak aktif")
	}

	if !dataBarber.IsActive {
		return errors.New("Tidak bisa transaksi,Barber sedang tidak aktif")
	}

	if !fn.InTimeActiveBarber(dataBarber, data.OrderDate) {
		return errors.New("Mohon maaf , waktu di luar jam oprasional")
	}

	mOrder.OrderDate = data.OrderDate
	mOrder.BarberID, _ = strconv.Atoi(Claims.BarberID)
	mOrder.Status = "N"
	mOrder.FromApps = false
	mOrder.CapsterID, _ = strconv.Atoi(Claims.CapsterID)
	mOrder.OrderNo, err = fn.GenTransactionNo(dataBarber.BarberCd)
	if err != nil {
		return err
	}

	mOrder.UserInput = Claims.CapsterID
	mOrder.UserEdit = Claims.CapsterID
	err = u.repoOrderH.Create(&mOrder)
	if err != nil {
		return err
	}

	for _, dataDetail := range data.Pakets {
		var mDetail models.OrderD
		err = mapstructure.Decode(dataDetail, &mDetail)
		if err != nil {
			return err
		}
		mDetail.BarberID = mOrder.BarberID
		mDetail.OrderID = mOrder.OrderID
		mDetail.UserEdit = Claims.CapsterID
		mDetail.UserInput = Claims.CapsterID
		err = u.repoOrderD.Create(&mDetail)
		if err != nil {
			return err
		}
	}

	return nil

}
func (u *useOrder) Update(ctx context.Context, Claims util.Claims, ID int, data models.OrderStatus) (err error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeOut)
	defer cancel()

	fn := &repofunction.FN{
		Claims: Claims,
	}

	if data.Status == "P" {
		CntProgress := fn.GetCountTrxProses()
		if CntProgress > 0 {
			return errors.New("Anda sedang dalam proses transaksi")
		}
	}

	var dataUpdate = map[string]interface{}{
		"status": data.Status,
	}

	err = u.repoOrderH.Update(ID, dataUpdate)
	if err != nil {
		return err
	}

	return nil
}
func (u *useOrder) Delete(ctx context.Context, Claims util.Claims, ID int) (err error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeOut)
	defer cancel()

	err = u.repoOrderH.Delete(ID)
	if err != nil {
		return err
	}
	return nil
}
