package create_transaction

import (
	"github.com/flvSantos15/dtmoney-transactions/internal/entity"
	"github.com/flvSantos15/dtmoney-transactions/internal/gateway"
	"github.com/shopspring/decimal"
)

type CreateTransactionInputDTO struct {
	Amount decimal.Decimal `json:"amount"`
	AmountType entity.TransactionType `json:"amount_type"`
}

type CreateTransactionUseCase struct {
	TransactionGateway gateway.TransactionGateway
}

func NewTransactionUseCase(transaction gateway.TransactionGateway) *CreateTransactionUseCase {
	return &CreateTransactionUseCase{
		TransactionGateway: transaction,
	}
}

func (usecase *CreateTransactionUseCase) Execute(input CreateTransactionInputDTO) error {
	transaction, err := entity.NewTransaction(input.Amount, input.AmountType)
	if err != nil {
		return err
	}
	err = usecase.TransactionGateway.Create(transaction)
	if err != nil {
		return err
	}

	return nil
}
