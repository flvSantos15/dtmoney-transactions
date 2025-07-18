package database

import (
	"database/sql"

	"github.com/flvSantos15/dtmoney-transactions/internal/entity"
)

type TransactionDB struct {
	db *sql.DB
}

func NewTransactionDB(db *sql.DB) *TransactionDB {
	return &TransactionDB{db: db}
}

func (transationDB *TransactionDB) CreateTransaction(transaction *entity.Transaction) (*entity.Transaction, error) {
	_, err := transationDB.db.Exec(`
		INSERT INTO transactions (id, amount, amount_type, created_at, updated_at)
		VALUES(?, ?, ?, ?, ?)`,
		transaction.ID, transaction.Amount, transaction.AmountType, transaction.CreateAt, transaction.UpdatedAt)
	
	if err != nil {
		return nil, err
	}

	return transaction, nil
}