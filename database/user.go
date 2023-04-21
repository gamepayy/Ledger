package database

import (
	"database/sql"
	"fmt"
	"log"
	"math/big"

	_ "github.com/go-sql-driver/mysql"
)

type Ledger struct {
	Account string  `json:"account"`
	Balance big.Int `json:"balance"`
}

func NewLedger(db *sql.DB, ledger Ledger) (bool, error) {
	// insert
	stmt, err := db.Prepare("INSERT INTO users (address, balance) VALUES (?, ?)")
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(ledger.Account, ledger.Balance.String())
	if err != nil {

		fmt.Println(err)
		return false, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		fmt.Println(err)
		return false, err
	}

	log.Printf("New ledger created with insertion id: %d", id)
	return true, nil
}
