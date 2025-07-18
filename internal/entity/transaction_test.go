package entity_test

import (
	"testing"

	"github.com/flvSantos15/dtmoney-transactions/internal/entity"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestCreateNewTransaction(t *testing.T) {
	amount := decimal.RequireFromString("100.00")
	transaction, err := entity.NewTransaction(amount, entity.Deposit)
	assert.Nil(t, err)
	assert.NotNil(t, transaction)
	assert.Equal(t,amount, transaction.Amount)
	assert.Equal(t, entity.Deposit, transaction.AmountType)
}

func TestCreateNewTransactionWithZeroAmount(t *testing.T) {
	amount := decimal.RequireFromString("0.00")
	transaction, err := entity.NewTransaction(amount, entity.Deposit)
	assert.NotNil(t, err)
	assert.Nil(t, transaction)
}

func TestCreateNewTransactionWithNoTypeAmount(t *testing.T) {
	amount := decimal.RequireFromString("80.00")
	transaction, err := entity.NewTransaction(amount, "")
	assert.NotNil(t, err)
	assert.Nil(t, transaction)
}

func TestCreateNewTransactionWithWrongTypeAmount(t * testing.T) {
	amount := decimal.RequireFromString("150.00")
	transaction, err := entity.NewTransaction(amount, "testing")
	assert.NotNil(t, err)
	assert.Nil(t, transaction)
}
