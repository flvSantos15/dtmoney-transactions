package entity

import (
	"errors"
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type Transaction struct {
	ID string `json:"id" valid:"notnull" gorm:"type:uuid;primary_key"`
	Amount float64 `json:"amount" valid:"notnull" gorm:"type:decimal"`
	AmountType string `json:"amount_type" valid:"notnull"`
	CreateAt time.Time `json:"created_at" valid:"-"`
	UpdatedAt time.Time `json:"updated_at" valid:"-"`
}

func NewTransaction(amount float64, amountType string) (*Transaction, error) {
	transaction := &Transaction{
		ID: uuid.NewV4().String(),
		Amount: amount,
		AmountType: amountType,
		CreateAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := transaction.Validate()
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (transaction *Transaction) Validate() error {
	_, err := govalidator.ValidateStruct(transaction)

	if err != nil {
		return err
	}

	if transaction.Amount <= 0 {
		return errors.New("amount must be greater than zero")
	}

	return nil
}