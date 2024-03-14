package controllers

import (
	"disbursement-service/exhosts"
	"disbursement-service/repositories"
	"disbursement-service/services"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
)

// Disbursement operation
type Disbursement struct {
	beego.Controller
}

// Disburse controller for disbursement.
// @Title Disburse controller for disbursement.
// @Description Disburse controller for disbursement.
// @Success 200 {object} models.Response
// @routers  [post]
func (c Disbursement) Disburse() {
	identifier := time.Now().UnixNano()

	o := orm.NewOrm()
	bRepo := repositories.NewBankRepository(o)
	dRepo := repositories.NewDisbursementRepository(o)
	gai := exhosts.NewGetAccountInfo(identifier)
	disburse := exhosts.NewDisbursement(identifier)
	svc := services.NewDisburseService(bRepo, gai, dRepo, disburse, o, identifier)
	resp := svc.Disburse(c.Ctx.Input.RequestBody)

	c.Data["json"] = resp
	c.ServeJSON()
	return
}

// Callback controller for callback notification.
// @Title Disburse controller for  callback notification.
// @Description Disburse controller for  callback notification.
// @Success 200 {object} models.Response
// @routers  [post, get]
func (c Disbursement) Callback() {
	identifier := time.Now().UnixNano()

	o := orm.NewOrm()
	dRepo := repositories.NewDisbursementRepository(o)
	svc := services.NewCallbackService(dRepo, o, identifier)
	resp := svc.Callback(c.Ctx.Input.RequestBody)

	c.Data["json"] = resp
	c.ServeJSON()
	return
}
