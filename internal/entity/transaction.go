package entity

import (
	"errors"
	"time"

	uuid "github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type TransactionType string

const (
	Withdraw TransactionType = "withdraw"
	Deposit  TransactionType = "deposit"
)

type Transaction struct {
	ID uuid.UUID `json:"id" gorm:"type:binary(16);primaryKey"`
	Amount decimal.Decimal `json:"amount" gorm:"type:decimal(10,2)"`
	AmountType TransactionType `json:"amount_type" gorm:"type:enum('withdraw', 'deposit')"`
	CreateAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewTransaction(amount decimal.Decimal, amountType TransactionType) (*Transaction, error) {
	id := uuid.New()
	transaction := &Transaction{
		ID: id,
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
	if transaction.Amount.LessThanOrEqual(decimal.Zero) {
		return errors.New("amount must be greater than zero")
	}

	if transaction.AmountType == "" {
		return errors.New("type of amount must be deposit or withdraw")
	}

	if transaction.AmountType != "deposit" && transaction.AmountType != "withdraw" {
		return errors.New("type of amount must be deposit or withdraw")
	}

	return nil
}