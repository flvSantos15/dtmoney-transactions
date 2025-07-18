package web

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/flvSantos15/dtmoney-transactions/internal/usecase/create_transaction"
)

type WebTransactionHandler struct {
	CreateTransactionUseCase create_transaction.CreateTransactionUseCase
}

func NewTransactionHandler(createTransactionUseCase create_transaction.CreateTransactionUseCase) *WebTransactionHandler {
	return &WebTransactionHandler{
		CreateTransactionUseCase: createTransactionUseCase,
	}
}

func (handler *WebTransactionHandler) CreateTransaction(writer http.ResponseWriter, req *http.Request) {
	var dto create_transaction.CreateTransactionInputDTO
	err := json.NewDecoder(req.Body).Decode(&dto)
	if err != nil {
		fmt.Println(err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = handler.CreateTransactionUseCase.Execute(dto)
	if err != nil {
		fmt.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")

	writer.WriteHeader(http.StatusCreated)
}