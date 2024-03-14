package repositories

import (
	"disbursement-service/models"
	"github.com/astaxie/beego/orm"
)

// DisbursementRepository struct
type DisbursementRepository struct {
	db orm.Ormer
}

// NewDisbursementRepository is func for initiate DisbursementRepository
func NewDisbursementRepository(o orm.Ormer) DisbursementRepository {
	return DisbursementRepository{db: o}
}

// Insert func for insert user
func (repo DisbursementRepository) Insert(user models.Disbursement) error {
	_, err := repo.db.Insert(&user)
	return err
}

// UpdateColumns function for update voucher detail data using certain columns
func (repo DisbursementRepository) UpdateColumns(d models.Disbursement, cols ...string) error {
	_, err := repo.db.Update(&d, cols...)
	return err
}

func (repo DisbursementRepository) FinPendingDisburseByReferenceNo(refNumber string) (models.Disbursement, error) {
	disburse := models.Disbursement{}
	err := repo.db.QueryTable("disbursements").
		Filter("status", "PENDING").
		Filter("ref_number", refNumber).
		One(&disburse)

	return disburse, err
}
