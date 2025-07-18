package entity_test

import (
	"testing"

	"github.com/flvSantos15/dtmoney-transactions/internal/entity"
	"github.com/stretchr/testify/assert"
)

func TestCreateNewTransaction(t *testing.T) {
	transaction, err := entity.NewTransaction(100, "deposit")
	assert.Nil(t, err)
	assert.NotNil(t, transaction)
	assert.Equal(t, 100.0, transaction.Amount)
	assert.Equal(t, "deposit", transaction.AmountType)
}

func TestCreateNewTransactionWithZeroAmount(t *testing.T) {
	transaction, err := entity.NewTransaction(0, "deposit")
	assert.NotNil(t, err)
	assert.Nil(t, transaction)
}

func TestCreateNewTransactionWithNoTypeAmount(t *testing.T) {
	transaction, err := entity.NewTransaction(80, "")
	assert.NotNil(t, err)
	assert.Nil(t, transaction)
}

func TestCreateNewTransactionWithWrongTypeAmount(t * testing.T) {
	transaction, err := entity.NewTransaction(150, "testing")
	assert.NotNil(t, err)
	assert.Nil(t, transaction)
}
