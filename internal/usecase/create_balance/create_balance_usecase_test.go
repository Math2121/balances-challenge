package createbalance

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

func TestCreateBalance(t *testing.T) {
	balanceGateway := BalanceGatewayMock{}
	balanceGateway.On("Save", mock.Anything).Return(nil)

	input := CreateBalanceInput{
		Amount:      100.0,
        AccountID: "123",
	}

	usecase := NewCreateBalanceUseCase(&balanceGateway)

	err := usecase.Execute(input)

	assert.Nil(t,err)
	balanceGateway.AssertExpectations(t)
}
