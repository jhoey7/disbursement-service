package services

import (
	"disbursement-service/models"
	"disbursement-service/utils"
	"encoding/json"
	"github.com/astaxie/beego/logs"
)

// ConvertRequest function for conversion request
func ConvertRequest(body []byte, request interface{}, identifier int64) models.Response {

	err := json.Unmarshal(body, &request)
	if err != nil {
		logs.Error("[%d] Error unmarshal from http request body to expected object request : %v", identifier, err)
		logs.Error("[%d] Request: %s", identifier, string(body))

		return models.ResponseError("Invalid Request", utils.ErrReqInvalid)
	}

	return models.ResponseSuccess(nil)
}
