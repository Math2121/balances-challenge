package gateway

import "github.com/Math2121/balances-challenge/internal/entity"

type BalanceGateway interface {
	List() ([]entity.Balance, error)
	Save(balance *entity.Balance) error
	GetBalance(accountId string) (*entity.Balance, error)
}
