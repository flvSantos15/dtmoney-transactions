package create_transaction

import (
	"testing"

	"github.com/flvSantos15/dtmoney-transactions/internal/entity"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type TransactionGatewayMock struct {
	mock.Mock
}

func (mock *TransactionGatewayMock) Create(transaction *entity.Transaction) error {
	args := mock.Called(transaction)
	return args.Error(0)
}

func TestCreateTransactionUseCase_Execute(t *testing.T) {
	mk := &TransactionGatewayMock{}
	mk.On("Create", mock.Anything).Return(nil)
	usecase := NewTransactionUseCase(mk)

	amount := decimal.RequireFromString("100.00")

	err := usecase.Execute(CreateTransactionInputDTO{
		Amount: amount,
		AmountType: entity.Deposit,
	})
	assert.Nil(t, err)
	mk.AssertExpectations(t)
	mk.AssertNumberOfCalls(t, "Create", 1)
}