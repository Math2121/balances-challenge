package getbalance

import (
	"github.com/Math2121/balances-challenge/internal/entity"
	"github.com/Math2121/balances-challenge/internal/gateway"
)

type GetBalanceInput struct {
	AccountID string `json:"account_id"`
}

type GetBalanceOutput struct {
	Balance *entity.Balance `json:"balance"`
}

type GetBalanceUseCase struct {
	gateway gateway.BalanceGateway
}

func NewGetBalanceUseCase(gateway gateway.BalanceGateway) *GetBalanceUseCase {
	return &GetBalanceUseCase{gateway: gateway}
}

func (uc *GetBalanceUseCase) Execute(input GetBalanceInput) (*GetBalanceOutput, error) {
	balance, err := uc.gateway.GetBalance(input.AccountID)
	if err != nil {
		return nil, err
	}

	return &GetBalanceOutput{Balance: balance}, nil
}
