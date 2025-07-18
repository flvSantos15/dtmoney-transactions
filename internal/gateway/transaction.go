package gateway

import (
	"github.com/flvSantos15/dtmoney-transactions/internal/entity"
	"github.com/shopspring/decimal"
)

type TransactionInputDTO struct {
	Amount decimal.Decimal `json:"amount"`
	AmountType entity.TransactionType `json:"amount_type"`
}

type TransactionGateway interface {
	Create(transaction *entity.Transaction) error
}