package routers

import (
	"jkl-finance/controllers"

	"github.com/beego/beego/v2/server/web"
)

func init() {
	web.Router("/", &controllers.AuthController{}, "get:Login;post:Login")

	web.Router("/pengajuan", &controllers.PengajuanController{}, "get:Index")
	web.Router("/pengajuan/new", &controllers.PengajuanController{}, "get:New")
	web.Router("/pengajuan/create", &controllers.PengajuanController{}, "post:Create")
	web.Router("/pengajuan/edit/:id", &controllers.PengajuanController{}, "get:Edit")
	web.Router("/pengajuan/update/:id", &controllers.PengajuanController{}, "post:Update")
	web.Router("/pengajuan/delete/:id", &controllers.PengajuanController{}, "post:Delete")

	web.Router("/dealer", &controllers.DealerController{}, "get:Index")
	web.Router("/dealer/new", &controllers.DealerController{}, "get:New")
	web.Router("/dealer/create", &controllers.DealerController{}, "post:Create")
	web.Router("/dealer/edit/:id", &controllers.DealerController{}, "get:Edit")
	web.Router("/dealer/update/:id", &controllers.DealerController{}, "post:Update")
	web.Router("/dealer/delete/:id", &controllers.DealerController{}, "post:Delete")

	web.Router("/backoffice", &controllers.BackofficeController{}, "get:Index")
	web.Router("/backoffice/doc/:id", &controllers.BackofficeController{}, "get:Doc")

	web.Router("/approval", &controllers.ApprovalController{}, "get:Index")
	web.Router("/approval/detail/:id", &controllers.ApprovalController{}, "get:Detail")
	web.Router("/approval/:id/approve", &controllers.ApprovalController{}, "post:Approve")
	web.Router("/approval/:id/reject", &controllers.ApprovalController{}, "post:Reject")
}
