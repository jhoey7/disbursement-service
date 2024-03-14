package services

import (
	"disbursement-service/models"
	"disbursement-service/utils"
	"errors"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"time"
)

// DisburseProcessor interface for insert disbursement to database
type DisburseProcessor interface {
	Insert(user models.Disbursement) error
	FinPendingDisburseByReferenceNo(refNumber string) (models.Disbursement, error)
	UpdateColumns(d models.Disbursement, cols ...string) error
}

// ExDisburseCreator interface for disburse to bank destination
type ExDisburseCreator interface {
	Disburse(req models.DisburseRequest, extID string) (models.Response, models.ExDisburseResponse)
}

// DisburseService struct
type DisburseService struct {
	Identifier        int64
	db                orm.Ormer
	bankFinder        BankFinder
	accountFinder     AccountFinder
	disburseProcessor DisburseProcessor
	exDisburseCreator ExDisburseCreator
}

// NewDisburseService func for initialize DisburseService
func NewDisburseService(bf BankFinder, af AccountFinder, dp DisburseProcessor,
	edc ExDisburseCreator, o orm.Ormer, i int64) DisburseService {
	return DisburseService{
		bankFinder:        bf,
		accountFinder:     af,
		disburseProcessor: dp,
		exDisburseCreator: edc,
		Identifier:        i,
		db:                o,
	}
}

// NewCallbackService func for initialize DisburseService
func NewCallbackService(dp DisburseProcessor, o orm.Ormer, i int64) DisburseService {
	return DisburseService{
		disburseProcessor: dp,
		Identifier:        i,
		db:                o,
	}
}

// Disburse func for disburse process
func (svc DisburseService) Disburse(b []byte) models.Response {
	request := models.DisburseRequest{}
	res := ConvertRequest(b, &request, svc.Identifier)
	if res.Code != utils.ErrNone {
		logs.Warn("[%d] Failed to convert request: %+v", svc.Identifier, res)
		return models.ResponseError(res.ErrorMessage, utils.ErrReqInvalid)
	}
	logs.Info("disburse request: %+v", request)

	bank, err := svc.bankFinder.FindByCode(request.BankCode)
	if err != nil && !errors.Is(err, orm.ErrNoRows) {
		logs.Warn("[%d] Failed to find bank by code: %s", svc.Identifier, err.Error())
		return models.ResponseError(utils.ErrorMessageDefault, utils.ErrorCodeDefault)
	}

	if bank.BankCode == "" {
		logs.Warn("[%d] Bank not exist exist: %s", svc.Identifier, request.BankCode)
		return models.ResponseError(utils.ErrorMessageBankNotFound, utils.ErrNotFound)
	}

	resp, _ := svc.accountFinder.GetAccountInfo(request.ToGetAccountReq())
	if resp.Code != utils.ErrNone {
		logs.Warn("[%d] Failed to find account whit bank code: %s", svc.Identifier, request.BankCode)
		return models.ResponseError(resp.ErrorMessage, resp.Code)
	}

	refID := fmt.Sprintf("%X", time.Now().UnixNano())
	disburseData := request.ToInsertRequest(refID)

	errDisburse := false
	resp, disbursementResp := svc.exDisburseCreator.Disburse(request, refID)
	if resp.Code != utils.ErrNone {
		errDisburse = true
		disburseData.Status = "FAILED"
		logs.Warn("[%d] Failed to disburse to destination: %s", svc.Identifier, resp.ErrorMessage)
	}

	svc.db.Begin()
	err = svc.disburseProcessor.Insert(disburseData)
	if err != nil {
		svc.db.Rollback()
		logs.Warn("[%d] Failed to insert disbursement: %s", svc.Identifier, err.Error())
		return models.ResponseError(utils.ErrorMessageDefault, utils.ErrorCodeDefault)
	}

	svc.db.Commit()

	if errDisburse {
		return models.ResponseError(utils.ErrorMessageDisburse, utils.ErrDisburse)
	}

	return models.ResponseSuccess(disbursementResp)
}

// Callback func for receive callback / notification from third party
func (svc DisburseService) Callback(b []byte) models.Response {
	request := models.DisburseCallbackReq{}
	res := ConvertRequest(b, &request, svc.Identifier)
	if res.Code != utils.ErrNone {
		logs.Warn("[%d] Failed to convert request: %+v", svc.Identifier, res)
		return models.ResponseError(res.ErrorMessage, utils.ErrReqInvalid)
	}
	logs.Info("callback request: %+v", request)

	disburse, err := svc.disburseProcessor.FinPendingDisburseByReferenceNo(request.ExternalID)
	if err != nil {
		logs.Warn("[%d] Failed to find pending disbursement: %s", svc.Identifier, err.Error())
		return models.ResponseError(utils.ErrorMessageDefault, utils.ErrorCodeDefault)
	}

	disburse.ExternalID = request.ID
	disburse.Status = request.Status
	disburse.UpdateTs = time.Now()

	svc.db.Begin()
	err = svc.disburseProcessor.UpdateColumns(disburse, "external_id", "status", "update_ts")
	if err != nil {
		svc.db.Rollback()
		logs.Warn("[%d] Failed to update disbursement: %s", svc.Identifier, err.Error())
		return models.ResponseError(utils.ErrorMessageDefault, utils.ErrorCodeDefault)
	}

	svc.db.Commit()

	return models.ResponseSuccess(request)
}
