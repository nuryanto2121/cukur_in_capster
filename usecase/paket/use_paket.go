package usepaket

import (
	"context"
	"fmt"
	"math"
	ipaket "nuryanto2121/cukur_in_capster/interface/paket"
	"nuryanto2121/cukur_in_capster/models"
	util "nuryanto2121/cukur_in_capster/pkg/utils"
	"strings"
	"time"
)

type usePaket struct {
	repoPaket      ipaket.Repository
	contextTimeOut time.Duration
}

func NewUsePaket(a ipaket.Repository, timeout time.Duration) ipaket.Usecase {
	return &usePaket{repoPaket: a, contextTimeOut: timeout}
}

func (u *usePaket) GetDataBy(ctx context.Context, ID int) (result *models.Paket, err error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeOut)
	defer cancel()

	result, err = u.repoPaket.GetDataBy(ID)
	if err != nil {
		return result, err
	}
	return result, nil
}
func (u *usePaket) GetList(ctx context.Context, Claims util.Claims, queryparam models.ParamList) (result models.ResponseModelList, err error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeOut)
	defer cancel()

	// var tUser = models.Paket{}
	/*membuat Where like dari struct*/
	if queryparam.Search != "" {
		// queryparam.Search = fmt.Sprintf("paket_name iLIKE '%%%s%%' OR descs iLIKE '%%%s%%'", queryparam.Search, queryparam.Search)
		queryparam.Search = strings.ToLower(fmt.Sprintf("%%%s%%", queryparam.Search))
	}

	if queryparam.InitSearch != "" {
		queryparam.InitSearch += " AND owner_id = " + Claims.OwnerID
	} else {
		queryparam.InitSearch = " owner_id = " + Claims.OwnerID
	}
	result.Data, err = u.repoPaket.GetList(queryparam)
	if err != nil {
		return result, err
	}

	result.Total, err = u.repoPaket.Count(queryparam)
	if err != nil {
		return result, err
	}

	// d := float64(result.Total) / float64(queryparam.PerPage)
	result.LastPage = int(math.Ceil(float64(result.Total) / float64(queryparam.PerPage)))
	result.Page = queryparam.Page

	return result, nil
}

// func (u *usePaket) Create(ctx context.Context, data *models.Paket) (err error) {
// 	ctx, cancel := context.WithTimeout(ctx, u.contextTimeOut)
// 	defer cancel()

// 	err = u.repoPaket.Create(data)
// 	if err != nil {
// 		return err
// 	}
// 	return nil

// }
// func (u *usePaket) Update(ctx context.Context, ID int, data interface{}) (err error) {
// 	ctx, cancel := context.WithTimeout(ctx, u.contextTimeOut)
// 	defer cancel()

// 	err = u.repoPaket.Update(ID, data)
// 	return nil
// }
// func (u *usePaket) Delete(ctx context.Context, ID int) (err error) {
// 	ctx, cancel := context.WithTimeout(ctx, u.contextTimeOut)
// 	defer cancel()

// 	err = u.repoPaket.Delete(ID)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
