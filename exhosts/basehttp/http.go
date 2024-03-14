package basehttp

import (
	"disbursement-service/models"
	"disbursement-service/utils"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/parnurzeal/gorequest"
	"net/http"
	"strconv"
	"time"
)

var debugClientHTTP bool

func init() {
	debugClientHTTP = beego.AppConfig.DefaultBool("debug.clienthttp", true)
}

// HTTPService struct
type HTTPService struct {
	GoRequest *gorequest.SuperAgent
}

// HTTPRequest func
func (service *HTTPService) HTTPRequest(jsonData interface{}) ([]byte, models.Response) {
	var resp *http.Response
	var bodyString string
	var errs []error

	logs.Info("Request url : " + service.GoRequest.Url)
	if jsonData != nil {
		resp, bodyString, errs = service.GoRequest.
			Send(jsonData).
			End()
	} else {
		resp, bodyString, errs = service.GoRequest.
			End()
	}

	if errs != nil {
		logs.Error("Failed connect service %v", errs[0])
		return []byte(bodyString), models.GetResponseObject(utils.FailedToConnectReceivedCode)
	}

	err := true
	if resp.StatusCode == 200 || resp.StatusCode == 201 {
		err = false
	}

	if err {
		errCode := strconv.Itoa(resp.StatusCode)
		logs.Error("Http status code not 200 : %+v", resp)
		return []byte(bodyString), models.GetResponseObject("ex-bank." + errCode)
	}

	return []byte(bodyString), models.ResponseSuccess(nil)
}

// GetAccountInfoHTTPRequest func
func GetAccountInfoHTTPRequest(service *HTTPService, jsonData interface{}, serviceName string) ([]byte, models.Response) {
	body, response := service.HTTPRequest(jsonData)
	if response.Code != utils.ErrNone {
		return body, response
	}

	var baseResponse []models.ExBankResponse
	if err := json.Unmarshal([]byte(body), &baseResponse); err != nil {
		logs.Error("Failed convert string body to struct for service "+serviceName, err)
		return body, models.GetResponseObject(utils.FailedToConnectReceivedCode)
	}

	return body, models.ResponseSuccess(nil)
}

// DisburseHTTPRequest func
func DisburseHTTPRequest(service *HTTPService, jsonData interface{}, serviceName string) ([]byte, models.Response) {
	body, response := service.HTTPRequest(jsonData)
	if response.Code != utils.ErrNone {
		return body, response
	}

	var baseResponse models.ExDisburseResponse
	if err := json.Unmarshal([]byte(body), &baseResponse); err != nil {
		logs.Error("Failed convert string body to struct for service "+serviceName, err)
		return body, models.GetResponseObject(utils.FailedToConnectReceivedCode)
	}

	return body, models.ResponseSuccess(nil)
}

// NewGetHTTPService factory method to initialize HTTPService specific for get
func NewGetHTTPService(url string, timeout string, retry int, sk string, ms int64) *HTTPService {
	duration, _ := time.ParseDuration(timeout)
	request := gorequest.New().
		SetDebug(debugClientHTTP).
		Get(url).
		Timeout(duration).
		Retry(retry, time.Second, http.StatusInternalServerError).
		Set("Content-Type", "application/json").
		Set("secretKet", sk).
		Set("timeStamp", strconv.Itoa(int(ms)))

	return &HTTPService{
		GoRequest: request,
	}
}

// NewPostHTTPService factory method to initialize HTTPService specific for post
func NewPostHTTPService(url string, timeout string, retry int, sk string, ms int64, signature string) *HTTPService {
	duration, _ := time.ParseDuration(timeout)
	request := gorequest.New().
		SetDebug(debugClientHTTP).
		Timeout(duration).
		Retry(retry, time.Second, http.StatusInternalServerError).
		Post(url).
		Set("Content-Type", "application/json").
		Set("secretKet", sk).
		Set("signature", signature).
		Set("milliseconds", strconv.Itoa(int(ms))).
		Set("signature", signature)

	return &HTTPService{
		GoRequest: request,
	}
}
