package database

import (
	"database/sql"
	"fmt"
	"log"
	"math/big"

	_ "github.com/go-sql-driver/mysql"
)

type Ledger struct {
	Account string `json:"account"`
	Balance string `json:"balance"`
}

func NewLedger(db *sql.DB, ledger Ledger) (bool, error) {
	// insert
	stmt, err := db.Prepare("INSERT INTO users (address, balance) VALUES (?, ?)")
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(ledger.Account, ledger.Balance)
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

func GetLedger(db *sql.DB, account string) (Ledger, error) {
	// get
	stmt, err := db.Prepare("SELECT balance FROM users WHERE address = ?")
	if err != nil {
		return Ledger{}, err
	}
	defer stmt.Close()

	var ledger Ledger

	err = stmt.QueryRow(account).Scan(&ledger.Balance)
	if err != nil {
		return Ledger{}, err
	}

	return ledger, nil
}

func WithdrawLedger(db *sql.DB, account string, token string, amount *big.Int) (bool, error) {

	// check if user has enough balance
	balanceCheck, err := HasEnoughBalance(db, account, amount)
	if err != nil {
		return false, err
	}

	if !balanceCheck.IsPossible {
		return false, fmt.Errorf("not enough balance")
	}

	// withdraw
	stmt, err := db.Prepare("UPDATE users SET balance = ? WHERE address = ?")
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(balanceCheck.NewBalance, account)
	if err != nil {
		return false, err
	}

	affect, err := res.RowsAffected()
	if err != nil {
		return false, err
	}

	log.Printf("Withdrawal affected %d rows", affect)
	return true, nil
}

/*
HasEnoughBalance checks if the user has enough balance

*/

type BalanceCheck struct {
	NewBalance string `json:"balance"`
	IsPossible bool   `json:"is_possible"`
}

func HasEnoughBalance(db *sql.DB, account string, amount *big.Int) (BalanceCheck, error) {

	Check := BalanceCheck{
		NewBalance: "0",
		IsPossible: false,
	}
	// get
	stmt, err := db.Prepare("SELECT * FROM users WHERE address = ?")
	if err != nil {
		return Check, err
	}
	defer stmt.Close()

	var ledger Ledger
	row := stmt.QueryRow(account)
	row.Scan(&ledger.Account, &ledger.Balance)
	if err != nil {
		return Check, err
	}

	ledgerBalance := new(big.Int)
	ledgerBalance, ok := ledgerBalance.SetString(ledger.Balance, 10)

	if !ok {
		return Check, fmt.Errorf("failed to convert string to big.Int")
	}

	if ledgerBalance.Cmp(amount) < 0 {
		return Check, nil
	}
	Check.IsPossible = true
	Check.NewBalance = (ledgerBalance.Sub(ledgerBalance, amount)).String()

	return Check, nil
}
