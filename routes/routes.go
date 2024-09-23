package routes

import (
	"nuryanto2121/cukur_in_capster/pkg/postgresdb"
	// sqlxposgresdb "nuryanto2121/cukur_in_capster/pkg/postgresqlxdb"
	"nuryanto2121/cukur_in_capster/pkg/setting"

	_saauthcont "nuryanto2121/cukur_in_capster/controllers/auth"
	_authuse "nuryanto2121/cukur_in_capster/usecase/auth"

	_saFilecont "nuryanto2121/cukur_in_capster/controllers/fileupload"
	_repoFile "nuryanto2121/cukur_in_capster/repository/ss_fileupload"
	_useFile "nuryanto2121/cukur_in_capster/usecase/ss_fileupload"

	_saPaketcont "nuryanto2121/cukur_in_capster/controllers/paket"
	_repoPaket "nuryanto2121/cukur_in_capster/repository/paket"
	_usePaket "nuryanto2121/cukur_in_capster/usecase/paket"

	_contUser "nuryanto2121/cukur_in_capster/controllers/user"
	_repoUser "nuryanto2121/cukur_in_capster/repository/ss_user"
	_useUser "nuryanto2121/cukur_in_capster/usecase/ss_user"

	_contOrder "nuryanto2121/cukur_in_capster/controllers/c_order"
	_repoOrderd "nuryanto2121/cukur_in_capster/repository/c_order_d"
	_repoOrder "nuryanto2121/cukur_in_capster/repository/c_order_h"
	_useOrder "nuryanto2121/cukur_in_capster/usecase/c_order"

	_repoNotif "nuryanto2121/cukur_in_capster/repository/notification"
	_useNotif "nuryanto2121/cukur_in_capster/usecase/notification"

	"time"

	"github.com/labstack/echo/v4"
)

// Echo :
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

	repoPaket := _repoPaket.NewRepoPaket(postgresdb.Conn)
	usePaket := _usePaket.NewUsePaket(repoPaket, timeoutContext)
	_saPaketcont.NewContPaket(e.E, usePaket)

	repoNotif := _repoNotif.NewRepoNotification(postgresdb.Conn)
	useNotif := _useNotif.NewUseNotification(repoNotif, timeoutContext)
	repoOrderD := _repoOrderd.NewRepoOrderD(postgresdb.Conn)
	repoOrder := _repoOrder.NewRepoOrderH(postgresdb.Conn)
	useOrder := _useOrder.NewUserMOrder(repoOrder, repoOrderD, useNotif, timeoutContext)
	_contOrder.NewContOrder(e.E, useOrder)

	//_saauthcont
	// repoAuth := _repoAuth.NewRepoOptionDB(postgresdb.Conn)
	useAuth := _authuse.NewUserAuth(repoUser, repoFile, timeoutContext)
	_saauthcont.NewContAuth(e.E, useAuth)

}
