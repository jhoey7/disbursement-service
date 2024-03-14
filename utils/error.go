package utils

const (
	ErrNone          = 200
	ErrorCodeDefault = 509
	ErrReqInvalid    = 400
	ErrNotFound      = 404
	ErrDisburse      = 501

	ErrorMessageDefault         = "Don't worry, we're handling this. You can try again after few minutes."
	ErrorMessageBankNotFound    = "Bank code not found."
	ErrorMessageDisburse        = "Failed to disburse"
	FailedToConnectReceivedCode = "failed.connect"
	InvalidJSONReceivedCode     = "invalid.json"
)
