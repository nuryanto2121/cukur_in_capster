package routes

import (
	"nuryanto2121/dynamic_rest_api_go/pkg/postgresdb"
	// sqlxposgresdb "nuryanto2121/dynamic_rest_api_go/pkg/postgresqlxdb"
	"nuryanto2121/dynamic_rest_api_go/pkg/setting"

	_saauthcont "nuryanto2121/dynamic_rest_api_go/controllers/auth"
	_authuse "nuryanto2121/dynamic_rest_api_go/usecase/auth"

	_saFilecont "nuryanto2121/dynamic_rest_api_go/controllers/fileupload"
	_repoFile "nuryanto2121/dynamic_rest_api_go/repository/ss_fileupload"
	_useFile "nuryanto2121/dynamic_rest_api_go/usecase/ss_fileupload"

	_contUser "nuryanto2121/dynamic_rest_api_go/controllers/user"
	_repoUser "nuryanto2121/dynamic_rest_api_go/repository/ss_user"
	_useUser "nuryanto2121/dynamic_rest_api_go/usecase/ss_user"

	_contOrder "nuryanto2121/dynamic_rest_api_go/controllers/c_order"
	_repoOrderd "nuryanto2121/dynamic_rest_api_go/repository/c_order_d"
	_repoOrder "nuryanto2121/dynamic_rest_api_go/repository/c_order_h"
	_useOrder "nuryanto2121/dynamic_rest_api_go/usecase/c_order"

	"time"

	"github.com/labstack/echo/v4"
)

//Echo :
type EchoRoutes struct {
	E *echo.Echo
}

func (e *EchoRoutes) InitialRouter() {
	timeoutContext := time.Duration(setting.FileConfigSetting.Server.ReadTimeout) * time.Second

	repoUser := _repoUser.NewRepoSysUser(postgresdb.Conn)
	useUser := _useUser.NewUserSysUser(repoUser, timeoutContext)
	_contUser.NewContUser(e.E, useUser)

	repoFile := _repoFile.NewRepoFileUpload(postgresdb.Conn)
	useFile := _useFile.NewSaFileUpload(repoFile, timeoutContext)
	_saFilecont.NewContFileUpload(e.E, useFile)

	repoOrderD := _repoOrderd.NewRepoOrderD(postgresdb.Conn)
	repoOrder := _repoOrder.NewRepoOrderH(postgresdb.Conn)
	useOrder := _useOrder.NewUserMOrder(repoOrder, repoOrderD, timeoutContext)
	_contOrder.NewContOrder(e.E, useOrder)

	//_saauthcont
	// repoAuth := _repoAuth.NewRepoOptionDB(postgresdb.Conn)
	useAuth := _authuse.NewUserAuth(repoUser, repoFile, timeoutContext)
	_saauthcont.NewContAuth(e.E, useAuth)

}
