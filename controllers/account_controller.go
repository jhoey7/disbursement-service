package controllers

import (
	"disbursement-service/exhosts"
	"disbursement-service/repositories"
	"disbursement-service/services"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
)

// Account operation
type Account struct {
	beego.Controller
}

// Validate controller for validate account.
// @Title Validate controller for validate account.
// @Description Validate controller for validate account.
// @Success 200 {object} models.Response
// @routers /validate [post]
func (c Account) Validate() {
	identifier := time.Now().UnixNano()

	bRepo := repositories.NewBankRepository(orm.NewOrm())
	gai := exhosts.NewGetAccountInfo(identifier)
	svc := services.NewAccountService(bRepo, gai, identifier)
	resp := svc.Find(c.Ctx.Input.RequestBody)

	c.Data["json"] = resp
	c.ServeJSON()
	return
}
