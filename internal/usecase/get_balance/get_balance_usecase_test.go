package getbalance

import (
	"testing"

	"github.com/Math2121/balances-challenge/internal/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type BalanceGatewayMock struct {
	mock.Mock
}

func (m *BalanceGatewayMock) Save(balance *entity.Balance) error {
	args := m.Called(balance)
	return args.Error(0)
}

func (m *BalanceGatewayMock) GetBalance(accountId string) (*entity.Balance, error) {
	args := m.Called(accountId)
	return args.Get(0).(*entity.Balance), args.Error(1)
}

func (m *BalanceGatewayMock) List() ([]entity.Balance, error) {
	args := m.Called()
	return args.Get(0).([]entity.Balance), args.Error(1)
}

func TestGetBalance(t *testing.T) {
	balanceGateway := &BalanceGatewayMock{}

	balanceGateway.On("GetBalance", "123").Return(&entity.Balance{ID: "123", AccountID: "123", Amount: 100.0}, nil)

	usecase := NewGetBalanceUseCase(balanceGateway)

	input := GetBalanceInput{
		AccountID: "123",
	}
	
	output, err := usecase.Execute(input)

	assert.Nil(t, err)
	assert.Equal(t, 100.0, output.Balance.Amount)
}
