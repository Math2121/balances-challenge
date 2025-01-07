package database

import (
	"database/sql"

	"github.com/Math2121/balances-challenge/internal/entity"
)

type BalanceDb struct {
	db *sql.DB
}

func NewBalanceDb(db *sql.DB) *BalanceDb {
	return &BalanceDb{db: db}
}

func (db *BalanceDb) GetBalance(accountID string) (*entity.Balance, error) {
	query := "SELECT * FROM balances WHERE id = $1"
	var balance entity.Balance
	err := db.db.QueryRow(query, accountID).Scan(&balance)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &balance, err

}

func (db *BalanceDb) Save(entity *entity.Balance) error {
	query := "INSERT INTO balances (id, account_id, amount) VALUES (?, ?, ?)"
	stmt, err := db.db.Prepare(query)

	if stmt == nil {
		return err
	}
	defer stmt.Close()

	_, err = db.db.Exec(query, entity.ID, entity.AccountID, entity.Amount)
	if err != nil {
		return err
	}
	return nil

}

func (db *BalanceDb) List() ([]entity.Balance, error) {

	query := "SELECT * FROM balances"
	rows, err := db.db.Query(query)
	if err!= nil {
        return nil, err
    }
	defer rows.Close()
	var balances []entity.Balance

	for rows.Next() {
		var balance entity.Balance
        err := rows.Scan(&balance.ID, &balance.AccountID, &balance.Amount)
        if err!= nil {
            return nil, err
        }
        balances = append(balances, balance)
	}
	
	return balances, nil

}
