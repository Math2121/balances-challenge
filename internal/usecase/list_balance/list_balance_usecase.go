package listbalance

import (
	"github.com/Math2121/balances-challenge/internal/entity"
	"github.com/Math2121/balances-challenge/internal/gateway"
)

type ListBalanceUseCase struct {
	gateway gateway.BalanceGateway
}

func NewListBalanceUseCase(gateway gateway.BalanceGateway) *ListBalanceUseCase {
    return &ListBalanceUseCase{gateway: gateway}
}

func (u *ListBalanceUseCase) Execute() ([]entity.Balance, error) {

	return u.gateway.List()
}