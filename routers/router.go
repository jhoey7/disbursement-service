package routers

import (
	"disbursement-service/controllers"
	"github.com/astaxie/beego"
)

func init() {
	ns :=
		beego.NewNamespace("/1.0",
			beego.NSNamespace("/accounts",
				beego.NSRouter("/validate", &controllers.Account{}, "post:Validate"),
			),

			beego.NSNamespace("/disburse",
				beego.NSRouter("", &controllers.Disbursement{}, "post:Disburse"),
				beego.NSRouter("/callback", &controllers.Disbursement{}, "post:Callback"),
				beego.NSRouter("/callback", &controllers.Disbursement{}, "get:Callback"),
			),
		)

	beego.AddNamespace(ns)
}
