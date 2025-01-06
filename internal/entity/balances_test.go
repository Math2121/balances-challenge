package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateEntity(t *testing.T) {

	balance := Balance{
		AccountID: "123",
		Amount:    100.0,
	}

	assert.Equal(t, "123", balance.AccountID)
	assert.Equal(t, 100.0, balance.Amount)

}
