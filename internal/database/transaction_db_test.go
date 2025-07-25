package database

import (
	"github.com/flvSantos15/dtmoney-transactions/internal/entity"
	"github.com/stretchr/testify/suite"
)

type TransactionDBTestSuite struct {
	suite.Suite
	transaction *entity.Transaction
	transactionDB *TransactionDB
}

func (suite *TransactionDBTestSuite) TestCreate() {
	transaction := suite.transaction
	err := suite.transactionDB.Create(transaction)
	suite.Nil(err)
}