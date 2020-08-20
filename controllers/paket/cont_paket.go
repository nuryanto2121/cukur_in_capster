package contpaket

import (
	"context"
	"fmt"
	"net/http"
	ipakets "nuryanto2121/cukur_in_capster/interface/paket"
	midd "nuryanto2121/cukur_in_capster/middleware"
	"nuryanto2121/cukur_in_capster/models"
	app "nuryanto2121/cukur_in_capster/pkg"
	tool "nuryanto2121/cukur_in_capster/pkg/tools"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ContPaket struct {
	usePaket ipakets.Usecase
}

func NewContPaket(e *echo.Echo, a ipakets.Usecase) {
	controller := &ContPaket{
		usePaket: a,
	}

	r := e.Group("/api.v1/capster/paket")
	r.Use(midd.JWT)
	r.GET("/:id", controller.GetDataBy)
	r.GET("", controller.GetList)

}

// GetDataByID :
// @Summary GetById
// @Security ApiKeyAuth
// @Tags Paket
// @Produce  json
// @Param OS header string true "OS Device"
// @Param Version header string true "OS Device"
// @Param id path string true "ID"
// @Success 200 {object} tool.ResponseModel
// @Router /api.v1/capster/paket/{id} [get]
func (u *ContPaket) GetDataBy(e echo.Context) error {
	ctx := e.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	var (
		// logger = logging.Logger{}
		appE = tool.Res{R: e} // wajib
		id   = e.Param("id")  //kalo bukan int => 0
		// valid  validation.Validation                 // wajib
	)
	ID, err := strconv.Atoi(id)
	if err != nil {
		return appE.Response(http.StatusBadRequest, fmt.Sprintf("%v", err), nil)
	}

	data, err := u.usePaket.GetDataBy(ctx, ID)
	if err != nil {
		return appE.Response(http.StatusInternalServerError, fmt.Sprintf("%v", err), nil)
	}

	return appE.Response(http.StatusOK, "Ok", data)
}

// GetList :
// @Summary GetList Paket
// @Security ApiKeyAuth
// @Tags Paket
// @Produce  json
// @Param OS header string true "OS Device"
// @Param Version header string true "OS Device"
// @Param page query int true "Page"
// @Param perpage query int true "PerPage"
// @Param search query string false "Search"
// @Param initsearch query string false "InitSearch"
// @Param sortfield query string false "SortField"
// @Success 200 {object} models.ResponseModelList
// @Router /api.v1/capster/paket [get]
func (u *ContPaket) GetList(e echo.Context) error {
	ctx := e.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	var (
		// logger = logging.Logger{}
		appE = tool.Res{R: e} // wajib
		//valid      validation.Validation // wajib
		paramquery   = models.ParamList{} // ini untuk list
		responseList = models.ResponseModelList{}
		err          error
	)

	httpCode, errMsg := app.BindAndValid(e, &paramquery)
	// logger.Info(util.Stringify(paramquery))
	if httpCode != 200 {
		return appE.ResponseErrorList(http.StatusBadRequest, errMsg, responseList)
	}
	// claims, err := app.GetClaims(e)
	// if err != nil {
	// 	return appE.ResponseErrorList(http.StatusBadRequest, fmt.Sprintf("%v", err), responseList)
	// }
	// if !claims.IsAdmin {
	// paramquery.InitSearch = " owner_id = " + claims.UserID
	// }

	responseList, err = u.usePaket.GetList(ctx, paramquery)
	if err != nil {
		// return e.JSON(http.StatusBadRequest, err.Error())
		return appE.ResponseErrorList(tool.GetStatusCode(err), fmt.Sprintf("%v", err), responseList)
	}

	// return e.JSON(http.StatusOK, ListDataPaket)
	return appE.ResponseList(http.StatusOK, "", responseList)
}
