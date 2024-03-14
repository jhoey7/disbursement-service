package models

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"time"
)

// Disbursement is bank struct
type Disbursement struct {
	ID            int64     `orm:"pk;auto;column(id)" json:"-"`
	AccountNumber string    `orm:"column(account_number)" json:"accountNumber"`
	AccountName   string    `orm:"column(account_name)" json:"accountName"`
	Amount        int       `orm:"column(amount)" json:"amount"`
	ReceiptEmail  string    `orm:"column(receipt_email)" json:"receiptEmail"`
	Remark        string    `orm:"column(remark)" json:"remark"`
	RefNumber     string    `orm:"column(ref_number)" json:"refNumber"`
	ExternalID    string    `orm:"column(external_id)" json:"externalId"`
	Status        string    `orm:"column(status)" json:"status"`
	CreateTs      time.Time `orm:"column(create_ts)" json:"createTs"`
	UpdateTs      time.Time `orm:"column(update_ts)" json:"updateTs"`
}

// TableName for users
func (b *Disbursement) TableName() string {
	return "disbursements"
}

// DisburseRequest struct
type DisburseRequest struct {
	Amount        int    `json:"amount"`
	Remark        string `json:"remark"`
	AccountNumber string `json:"accountNumber"`
	AccountName   string `json:"accountName"`
	ReceiptEmail  string `json:"receiptEmail"`
	ExternalID    string `json:"externalId"`
	BankCode      string `json:"bankCode"`
}

// ToGetAccountReq func
func (dr DisburseRequest) ToGetAccountReq() HTTPParams {
	hp := new(HTTPParams)

	hp.BuildParam("accountNumber", dr.AccountNumber)
	hp.BuildParam("accountName", dr.AccountName)

	return *hp
}

// ToInsertRequest func
func (dr DisburseRequest) ToInsertRequest(refID string) Disbursement {
	return Disbursement{
		AccountNumber: dr.AccountNumber,
		AccountName:   dr.AccountName,
		Amount:        dr.Amount,
		ReceiptEmail:  dr.ReceiptEmail,
		Remark:        dr.Remark,
		RefNumber:     refID,
		Status:        "PENDING",
		CreateTs:      time.Now(),
	}
}

// ToExDisburseReq func
func (dr DisburseRequest) ToExDisburseReq(externalID string) ExDisburseRequest {
	return ExDisburseRequest{
		AccountNumber: dr.AccountNumber,
		AccountName:   dr.AccountName,
		Amount:        dr.Amount,
		ReceiptEmail:  dr.ReceiptEmail,
		Remark:        dr.Remark,
		ExternalId:    externalID,
	}
}

// ExDisburseRequest struct
type ExDisburseRequest struct {
	Amount        int    `json:"amount"`
	Remark        string `json:"remark"`
	AccountNumber string `json:"accountNumber"`
	AccountName   string `json:"accountName"`
	ReceiptEmail  string `json:"receiptEmail"`
	ExternalId    string `json:"externalId"`
}

// ExDisburseResponse struct
type ExDisburseResponse struct {
	CreatedAt     string `json:"createdAt"`
	Amount        int    `json:"amount"`
	Remark        string `json:"remark"`
	AccountNumber string `json:"accountNumber"`
	AccountName   string `json:"accountName"`
	ReceiptEmail  string `json:"receiptEmail"`
	ExternalID    string `json:"externalId"`
	ID            string `json:"id"`
}

// DisburseCallbackReq struct
type DisburseCallbackReq struct {
	CreatedAt     string `json:"createdAt"`
	Amount        int    `json:"amount"`
	Remark        string `json:"remark"`
	AccountNumber string `json:"accountNumber"`
	AccountName   string `json:"accountName"`
	ReceiptEmail  string `json:"receiptEmail"`
	ExternalID    string `json:"externalId"`
	ID            string `json:"id"`
	Status        string `json:"status"`
}

// GetSignature func to get signature
func GetSignature(bodyReq []byte, appKey string) string {
	body := MD5Encode(bodyReq)
	return ComputeHmac256(body, appKey)
}

// MD5Encode func to convert data into md5
func MD5Encode(vMar []byte) string {
	return fmt.Sprintf("%x", md5.Sum(vMar))
}

// ComputeHmac256 func to convert data into HMAC SHA-256
func ComputeHmac256(message string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}
