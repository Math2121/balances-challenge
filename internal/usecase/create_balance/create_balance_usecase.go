package createbalance

import (
	"github.com/Math2121/balances-challenge/internal/entity"
	"github.com/Math2121/balances-challenge/internal/gateway"
)

type CreateBalanceUseCase struct {
	Gateway gateway.BalanceGateway
}

type CreateBalanceInput struct {
	AccountID string
	Amount    float64
}

func NewCreateBalanceUseCase(gateway gateway.BalanceGateway) *CreateBalanceUseCase {
	return &CreateBalanceUseCase{Gateway: gateway}
}
func (c *CreateBalanceUseCase) Execute(input CreateBalanceInput) error {
	balance := entity.Balance{
		AccountID: input.AccountID,
		Amount:    input.Amount,
	}

	err := c.Gateway.Save(&balance)
	if err != nil {
		return err
	}
	return nil
}
