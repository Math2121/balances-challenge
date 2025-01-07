package database

import (
	"database/sql"
	"testing"

	"github.com/Math2121/balances-challenge/internal/entity"
	"github.com/stretchr/testify/suite"
)

type BalanceDbTestSuite struct {
	suite.Suite
	db        *sql.DB
	balanceDb *BalanceDb
}

func (s *BalanceDbTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Nil(err)
	s.db = db
	db.Exec("Create table balances (id varchar(255) not null, account_id varchar(255), amount int )")

	s.balanceDb = NewBalanceDb(db)
}

func (s *BalanceDbTestSuite) TearDownSuite() {
	defer s.db.Close()
	s.db.Exec("DROP TABLE balances")
}
func TestBalanceDbTestSuite(t *testing.T) {
	suite.Run(t, new(BalanceDbTestSuite))
}

func (s *BalanceDbTestSuite) TestSaveBalance() {
	balance := entity.NewBalance("12", 100.0)

	err := s.balanceDb.Save(balance)

	output, _ := s.balanceDb.GetBalance(balance.AccountID)

	s.Equal(balance, output)
	s.Nil(err)

}
