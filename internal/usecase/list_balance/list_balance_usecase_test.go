package listbalance

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

func TestListAllBalances(t *testing.T) {
	balanceGatewayMock := &BalanceGatewayMock{}
	balanceGatewayMock.On("List").Return([]entity.Balance{{
		AccountID: "1", Amount: 100.0,
	}, {AccountID: "2", Amount: 200.0}}, nil)

	listBalance := NewListBalanceUseCase(balanceGatewayMock)

	output, err := listBalance.Execute()

	assert.Nil(t, err)
	assert.Equal(t, 2, len(output))
	assert.Equal(t, "1", output[0].AccountID)
	assert.Equal(t, 100.0, output[0].Amount)
	assert.Equal(t, "2", output[1].AccountID)
	assert.Equal(t, 200.0, output[1].Amount)

}
