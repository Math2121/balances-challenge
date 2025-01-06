package entity

import (
	"github.com/google/uuid"
)

type Balance struct {
	ID            string  `json:"id"`
	AccountID   string  `json:"account_id"`
	Amount        float64 `json:"amount"`
}

func NewBalance(accountID string, amount float64) *Balance {
	return &Balance{
		ID:            uuid.New().String(),
		AccountID:   accountID,
		Amount:        amount,
	}
}
