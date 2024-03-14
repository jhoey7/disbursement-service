package exhosts

import (
	"disbursement-service/exhosts/basehttp"
	"disbursement-service/models"
	"disbursement-service/utils"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"time"
)

var (
	ExBankHost               = beego.AppConfig.DefaultString("bank.url", "https://65f2112b034bdbecc7645157.mockapi.io/")
	SecretKey                = beego.AppConfig.DefaultString("bank.secret.key", "8c78707182475d89b9a7313ac9143758")
	ExBankTimeout            = beego.AppConfig.DefaultString("bank.timeout", "30s")
	ExBankRetryBadReq        = beego.AppConfig.DefaultInt("bank.retrybadreq", 3)
	ExBankServiceName        = "bankservice"
	ExBankValidateAccountURL = beego.AppConfig.DefaultString("bank.validate.account.url", "validate-account")
	ExBankDisburseRL         = beego.AppConfig.DefaultString("bank.disburse.url", "disbursement")
)

// BankHost struct
type BankHost struct {
	HTTPRequest func(service *basehttp.HTTPService, jsondata interface{},
		serviceName string) ([]byte, models.Response)
	Identifier int64
}

// NewGetAccountInfo factory method to initialize BankHost with identifier.
func NewGetAccountInfo(i int64) *BankHost {
	return &BankHost{
		HTTPRequest: basehttp.GetAccountInfoHTTPRequest,
		Identifier:  i,
	}
}

// NewDisbursement factory method to initialize BankHost with identifier.
func NewDisbursement(i int64) *BankHost {
	return &BankHost{
		HTTPRequest: basehttp.DisburseHTTPRequest,
		Identifier:  i,
	}
}

// GetAccountInfo host function for get account info.
func (service *BankHost) GetAccountInfo(hp models.HTTPParams) (models.Response, []models.ExBankResponse) {
	qp := hp.URLValues()
	endpoint := ExBankHost + ExBankValidateAccountURL + "?" + qp.Encode()

	httpReq := basehttp.NewGetHTTPService(endpoint, ExBankTimeout, ExBankRetryBadReq, SecretKey, time.Now().UnixMilli())
	body, response := service.HTTPRequest(httpReq, nil, ExBankServiceName+"-GetAccountInfo")

	var bankRes []models.ExBankResponse
	if response.Code != utils.ErrNone {
		logs.Warn(response.ErrorMessage)
		return response, bankRes
	}

	if json.Unmarshal(body, &bankRes) != nil {
		logs.Warn("Response : %s", string(body))
		logs.Warn("Failed unmarshal get account info response")
		return models.ResponseError(utils.InvalidJSONReceivedCode, utils.ErrReqInvalid), bankRes
	}

	return response, bankRes
}

// Disburse host function for disburse.
func (service *BankHost) Disburse(req models.DisburseRequest, extID string) (models.Response, models.ExDisburseResponse) {
	endpoint := ExBankHost + ExBankDisburseRL
	payload := req.ToExDisburseReq(extID)
	bodyByte, _ := json.Marshal(payload)
	signature := models.GetSignature(bodyByte, SecretKey)
	httpReq := basehttp.NewPostHTTPService(endpoint, ExBankTimeout, ExBankRetryBadReq, SecretKey, time.Now().Unix(), signature)
	body, response := service.HTTPRequest(httpReq, payload, ExBankServiceName+"-Disburse")

	var disburseRes models.ExDisburseResponse
	if response.Code != utils.ErrNone {
		logs.Warn(response.ErrorMessage)
		return response, disburseRes
	}

	if json.Unmarshal(body, &disburseRes) != nil {
		logs.Warn("Response : %s", string(body))
		logs.Warn("Failed unmarshal disburse response")
		return models.ResponseError(utils.InvalidJSONReceivedCode, utils.ErrReqInvalid), disburseRes
	}

	return response, disburseRes
}
