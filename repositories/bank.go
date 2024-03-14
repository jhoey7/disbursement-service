package repositories

import (
	"disbursement-service/models"
	"github.com/astaxie/beego/orm"
)

// BankRepository struct
type BankRepository struct {
	db orm.Ormer
}

// NewBankRepository is func for initiate BankRepository
func NewBankRepository(o orm.Ormer) BankRepository {
	return BankRepository{db: o}
}

// FindByCode is func for find bank by code
func (repo BankRepository) FindByCode(bankCode string) (models.Bank, error) {
	bank := models.Bank{}
	err := repo.db.QueryTable("banks").
		Filter("bank_code", bankCode).
		One(&bank)

	return bank, err
}
