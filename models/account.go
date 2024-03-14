package models

// GetAccountInfoRequest struct
type GetAccountInfoRequest struct {
	AccountNumber string `json:"accountNumber"`
	AccountName   string `json:"accountName"`
	BankCode      string `json:"bankCode"`
}

func (r GetAccountInfoRequest) ToHTTPParams() HTTPParams {
	hp := new(HTTPParams)

	hp.BuildParam("accountNumber", r.AccountNumber)
	hp.BuildParam("accountName", r.AccountName)

	return *hp
}
