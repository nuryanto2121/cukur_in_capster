package contauth

import (
	"context"
	"fmt"
	"net/http"
	iauth "nuryanto2121/cukur_in_capster/interface/auth"
	midd "nuryanto2121/cukur_in_capster/middleware"
	"nuryanto2121/cukur_in_capster/models"
	app "nuryanto2121/cukur_in_capster/pkg"
	"nuryanto2121/cukur_in_capster/pkg/logging"
	tool "nuryanto2121/cukur_in_capster/pkg/tools"
	util "nuryanto2121/cukur_in_capster/pkg/utils"

	_ "nuryanto2121/cukur_in_capster/docs"

	"github.com/labstack/echo/v4"
)

type ContAuth struct {
	useAuth iauth.Usecase
}

func NewContAuth(e *echo.Echo, useAuth iauth.Usecase) {
	cont := &ContAuth{
		useAuth: useAuth,
		// useSaClient:     useSaClient,
		// useSaUser:       useSaUser,
		// useSaFileUpload: useSaFileUpload,
	}

	// e.POST("/capster/auth/register", cont.Register)
	r := e.Group("/capster/auth")
	r.Use(midd.Versioning)
	r.POST("/login", cont.Login)
	r.POST("/forgot", cont.ForgotPassword)
	r.POST("/change_password", cont.ChangePassword)
	r.POST("/verify", cont.Verify)
	r.POST("/register", cont.Register)

	L := e.Group("/capster/auth")
	L.Use(midd.JWT)
	L.Use(midd.Versioning)
	L.POST("/logout", cont.Logout)
	L.PATCH("/fcm", cont.pathFCM)
}

// Logout :
// @Summary logout
// @Security ApiKeyAuth
// @Tags Auth
// @Produce json
// @Param OS header string true "OS Device"
// @Param Version header string true "OS Device"
// @Success 200 {object} tool.ResponseModel
// @Router /capster/auth/logout [post]
func (u *ContAuth) Logout(e echo.Context) error {
	ctx := e.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	var (
		appE = tool.Res{R: e} // wajib
	)

	claims, err := app.GetClaims(e)
	if err != nil {
		return appE.Response(http.StatusBadRequest, fmt.Sprintf("%v", err), nil)
	}
	Token := e.Request().Header.Get("Authorization")
	err = u.useAuth.Logout(ctx, claims, Token)
	if err != nil {
		return appE.ResponseError(http.StatusUnauthorized, fmt.Sprintf("%v", err), nil)
	}

	return appE.Response(http.StatusOK, "Ok", nil)
}

// Login :
// @Summary Login
// @Tags Auth
// @Produce json
// @Param OS header string true "OS Device"
// @Param Version header string true "OS Device"
// @Param req body models.LoginForm true "req param #changes are possible to adjust the form of the registration form from frontend"
// @Success 200 {object} tool.ResponseModel
// @Router /capster/auth/login [post]
func (u *ContAuth) Login(e echo.Context) error {
	ctx := e.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	var (
		logger = logging.Logger{} // wajib
		appE   = tool.Res{R: e}   // wajib
		// client sa_models.SaClient

		form = models.LoginForm{}
		// dataFiles = sa_models.SaFileOutput{}
	)

	// validasi and bind to struct
	httpCode, errMsg := app.BindAndValid(e, &form)
	logger.Info(util.Stringify(form))
	if httpCode != 200 {
		return appE.ResponseError(http.StatusBadRequest, errMsg, nil)
	}

	out, err := u.useAuth.Login(ctx, &form)
	if err != nil {
		// return appE.Response(out)
		// return appE.ResponseError(util.GetStatusCode(err), fmt.Sprintf("%v", err), nil)
		return appE.ResponseError(http.StatusUnauthorized, fmt.Sprintf("%v", err), nil)
	}

	return appE.Response(http.StatusOK, "Ok", out)
}

// ChangePassword :
// @Summary Change Password
// @Tags Auth
// @Produce json
// @Param OS header string true "OS Device"
// @Param Version header string true "OS Device"
// @Param req body models.ResetPasswd true "req param #changes are possible to adjust the form of the registration form from frontend"
// @Success 200 {object} tool.ResponseModel
// @Router /capster/auth/change_password [post]
func (u *ContAuth) ChangePassword(e echo.Context) error {
	ctx := e.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	var (
		logger = logging.Logger{} // wajib
		appE   = tool.Res{R: e}   // wajib
		// client sa_models.SaClient

		form = models.ResetPasswd{}
	)
	httpCode, errMsg := app.BindAndValid(e, &form)
	logger.Info(util.Stringify(form))
	if httpCode != 200 {
		return appE.ResponseError(http.StatusBadRequest, errMsg, nil)
	}
	err := u.useAuth.ResetPassword(ctx, &form)
	if err != nil {
		return appE.ResponseError(http.StatusUnauthorized, fmt.Sprintf("%v", err), nil)
	}

	return appE.Response(http.StatusOK, "Ok", "Please Login")
}

// Register :
// @Summary Register
// @Tags Auth
// @Produce json
// @Param OS header string true "OS Device"
// @Param Version header string true "OS Device"
// @Param req body models.RegisterForm true "req param #changes are possible to adjust the form of the registration form from frontend"
// @Success 200 {object} tool.ResponseModel
// @Router /capster/auth/register [post]
func (u *ContAuth) Register(e echo.Context) error {
	ctx := e.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	var (
		logger = logging.Logger{} // wajib
		appE   = tool.Res{R: e}   // wajib
		// client sa_models.SaClient

		form = models.RegisterForm{}
	)

	// validasi and bind to struct
	httpCode, errMsg := app.BindAndValid(e, &form)
	logger.Info(util.Stringify(form))
	if httpCode != 200 {
		return appE.ResponseError(http.StatusBadRequest, errMsg, nil)
	}

	data, err := u.useAuth.Register(ctx, form)
	if err != nil {
		return appE.ResponseError(http.StatusUnauthorized, fmt.Sprintf("%v", err), nil)
	}
	return appE.Response(http.StatusOK, "Ok", data)
}

// Verify :
// @Summary Verify
// @Tags Auth
// @Produce json
// @Param OS header string true "OS Device"
// @Param Version header string true "OS Device"
// @Param req body models.VerifyForm true "req param #changes are possible to adjust the form of the registration form from frontend"
// @Success 200 {object} tool.ResponseModel
// @Router /capster/auth/verify [post]
func (u *ContAuth) Verify(e echo.Context) error {
	ctx := e.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	var (
		logger = logging.Logger{} // wajib
		appE   = tool.Res{R: e}   // wajib
		// client sa_models.SaClient

		form = models.VerifyForm{}
	)

	// validasi and bind to struct
	httpCode, errMsg := app.BindAndValid(e, &form)
	logger.Info(util.Stringify(form))
	if httpCode != 200 {
		return appE.ResponseError(http.StatusBadRequest, errMsg, nil)
	}

	err := u.useAuth.Verify(ctx, form)
	if err != nil {
		return appE.ResponseError(http.StatusUnauthorized, fmt.Sprintf("%v", err), nil)
	}
	return appE.Response(http.StatusOK, "Ok", nil)
}

// ForgotPassword :
// @Summary Forgot Password
// @Tags Auth
// @Produce json
// @Param OS header string true "OS Device"
// @Param Version header string true "OS Device"
// @Param req body models.ForgotForm true "req param #changes are possible to adjust the form of the registration form from frontend"
// @Success 200 {object} tool.ResponseModel
// @Router /capster/auth/forgot [post]
func (u *ContAuth) ForgotPassword(e echo.Context) error {
	ctx := e.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	var (
		logger = logging.Logger{} // wajib
		appE   = tool.Res{R: e}   // wajib
		// client sa_models.SaClient

		form = models.ForgotForm{}
	)
	// validasi and bind to struct
	httpCode, errMsg := app.BindAndValid(e, &form)
	logger.Info(util.Stringify(form))
	if httpCode != 200 {
		return appE.ResponseError(http.StatusBadRequest, errMsg, nil)
	}

	err := u.useAuth.ForgotPassword(ctx, &form)
	if err != nil {
		return appE.ResponseError(http.StatusUnauthorized, fmt.Sprintf("%v", err), nil)
	}

	return appE.Response(http.StatusOK, "Check Your Email", nil)

}

// path-fcm :
// @Summary path-fcm
// @Security ApiKeyAuth
// @Tags Auth
// @Produce json
// @Param OS header string true "OS Device"
// @Param Version header string true "Version"
// @Param req body models.PathFCM true "req param #changes are possible to adjust the form of the registration form from frontend"
// @Success 200 {object} tool.ResponseModel
// @Router /user/auth/fcm [patch]
func (u *ContAuth) pathFCM(e echo.Context) error {
	ctx := e.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	var (
		logger = logging.Logger{} // wajib
		appE   = tool.Res{R: e}   // wajib
		// client sa_models.SaClient

		form = models.PathFCM{}
	)

	// validasi and bind to struct
	httpCode, errMsg := app.BindAndValid(e, &form)
	logger.Info(util.Stringify(form))
	if httpCode != 200 {
		return appE.ResponseError(http.StatusBadRequest, errMsg, nil)
	}

	claims, err := app.GetClaims(e)
	if err != nil {
		return appE.Response(http.StatusBadRequest, fmt.Sprintf("%v", err), nil)
	}

	form.CapsterID = claims.CapsterID
	err = u.useAuth.PathFCM(ctx, form)
	if err != nil {
		return appE.ResponseError(http.StatusUnauthorized, fmt.Sprintf("%v", err), nil)
	}
	return appE.Response(http.StatusOK, "Ok", nil)
}
