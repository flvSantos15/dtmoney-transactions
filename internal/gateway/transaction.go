package gateway

import (
	"github.com/flvSantos15/dtmoney-transactions/internal/entity"
)

type TransactionGateway interface {
	Create(transaction *entity.Transaction) error
}