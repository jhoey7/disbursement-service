package services

import (
	"disbursement-service/models"
	"disbursement-service/utils"
	"errors"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

// BankFinder find bank
type BankFinder interface {
	FindByCode(bankCode string) (models.Bank, error)
}

// AccountFinder find account to third party ( bank )
type AccountFinder interface {
	GetAccountInfo(hp models.HTTPParams) (models.Response, []models.ExBankResponse)
}

type AccountService struct {
	Identifier    int64
	bankFinder    BankFinder
	accountFinder AccountFinder
}

func NewAccountService(bf BankFinder, af AccountFinder, i int64) AccountService {
	return AccountService{
		bankFinder:    bf,
		accountFinder: af,
		Identifier:    i,
	}
}

func (svc AccountService) Find(b []byte) models.Response {
	request := models.GetAccountInfoRequest{}
	res := ConvertRequest(b, &request, svc.Identifier)
	if res.Code != utils.ErrNone {
		logs.Warn("[%d] Failed to convert request: %+v", svc.Identifier, res)
		return models.ResponseError(res.ErrorMessage, utils.ErrReqInvalid)
	}
	logs.Info("findAccount request: %+v", request)

	bank, err := svc.bankFinder.FindByCode(request.BankCode)
	if err != nil && !errors.Is(err, orm.ErrNoRows) {
		logs.Warn("[%d] Failed to find bank by code: %s", svc.Identifier, err.Error())
		return models.ResponseError(utils.ErrorMessageDefault, utils.ErrorCodeDefault)
	}

	fmt.Println(bank)

	if bank.BankCode == "" {
		logs.Warn("[%d] Bank not exist exist: %s", svc.Identifier, request.BankCode)
		return models.ResponseError(utils.ErrorMessageBankNotFound, utils.ErrNotFound)
	}

	resp, account := svc.accountFinder.GetAccountInfo(request.ToHTTPParams())
	if resp.Code != utils.ErrNone {
		logs.Warn("[%d] Bank not exist exist: %s", svc.Identifier, request.BankCode)
		return models.ResponseError(resp.ErrorMessage, resp.Code)
	}

	return models.ResponseSuccess(account[0].ToGwResponse(bank))
}
