package models

import (
	"disbursement-service/utils"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
	"io/ioutil"
	"path"
	"runtime"
)

// Response struct
type Response struct {
	Content        interface{} `json:"content"`
	SuccessMessage string      `json:"successMessage"`
	Code           int         `json:"code"`
	ErrorMessage   string      `json:"errorMessage"`
}

// ResponseSuccess function to generate non pointer success response
func ResponseSuccess(result interface{}) Response {
	return Response{
		Code:           utils.ErrNone,
		Content:        result,
		SuccessMessage: "Success",
	}
}

// ObjectError model
type ObjectError struct {
	ErrorCode int    `json:"errorCode"`
	ErrorDesc string `json:"errorDesc"`
}

// ResponseError function to generate non pointer error response
func ResponseError(err string, code int) Response {
	return Response{
		Code:         code,
		ErrorMessage: err,
	}
}

// JSONErrorCode : mapping error code model
type JSONErrorCode struct {
	ReceiveCode string `json:"receiveCode"`
	ObjectError `json:"objectError"`
}

// JSONErrorCodes : array of JSONErrorCode
type JSONErrorCodes struct {
	MappingErrorCodes []JSONErrorCode `json:"mappingErrorCodes"`
}

// MErrorCodes for err codes mapping
var (
	MErrorCodes = make(map[string]ObjectError)
)

// RegisterErrorCode init errorCode list on memory
func RegisterErrorCode() bool {
	logs.Info("Register error code from json file")

	var b []byte
	var err error
	b, err = ioutil.ReadFile("errorcode.json") // just pass the file name
	if err != nil {
		_, file, _, _ := runtime.Caller(0)
		dir := path.Join(path.Dir(file), "..")
		errorCodeFile := fmt.Sprintf("%s/errorcode.json", dir)
		b, err = ioutil.ReadFile(errorCodeFile) // just pass the file name
		if err != nil {
			logs.Error("Failed to read file error code json ", err)
		}
	}

	errorCodeJSON := JSONErrorCodes{}

	if json.Unmarshal(b, &errorCodeJSON) != nil {
		logs.Error("Unmarshal [%v] or JSONErrorCode Failed : [%d]", err)
		return false
	}

	for _, element := range errorCodeJSON.MappingErrorCodes {
		MErrorCodes[element.ReceiveCode] = element.ObjectError
	}

	return true
}

// GetResponseObject mapping error code with Response pointer return struct
func GetResponseObject(key string) Response {
	if (ObjectError{}) != MErrorCodes[key] {
		return ResponseError(MErrorCodes[key].ErrorDesc, MErrorCodes[key].ErrorCode)
	}

	logs.Warn("Err Code Not Mapping")
	return ResponseError(utils.ErrorMessageDefault, utils.ErrorCodeDefault)
}
