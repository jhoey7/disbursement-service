package models

import "time"

// Bank is bank struct
type Bank struct {
	ID       int64     `orm:"pk;auto;column(id)" json:"-"`
	BankCode string    `orm:"column(bank_code)" json:"bankCode"`
	BankName string    `orm:"column(bank_name)" json:"bankName"`
	CreateTs time.Time `orm:"column(create_ts)" json:"createTs"`
	UpdateTs time.Time `orm:"column(update_ts)" json:"updateTs"`
}

// TableName for users
func (b *Bank) TableName() string {
	return "banks"
}

// ExBankResponse struct
type ExBankResponse struct {
	ID            string `json:"id"`
	AccountName   string `json:"accountName"`
	AccountNumber string `json:"accountNumber"`
}

func (xbr ExBankResponse) ToGwResponse(bank Bank) GWValidateResponse {
	return GWValidateResponse{
		ExBankResponse:  xbr,
		BankInformation: bank,
	}
}

// GWValidateResponse struct
type GWValidateResponse struct {
	ExBankResponse
	BankInformation Bank `json:"bank"`
}
