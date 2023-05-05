package database

import (
	"fmt"
	"log"
	"math/big"

	_ "github.com/go-sql-driver/mysql"
)

type Ledger struct {
	Account string `json:"account"`
	Balance string `json:"balance"`
}

func NewLedger(ledger Ledger) (bool, error) {

	db, err := connectDB()

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

func GetLedger(account string) (Ledger, error) {

	db, err := connectDB()

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

	ledger.Account = account

	return ledger, nil
}

func DeleteLedger(account string) (string, error) {
	db, err := connectDB()
	fmt.Println(account)
	stmt, err := db.Prepare("DELETE FROM users WHERE address = ?")
	if err != nil {
		return "DB statement preparation error ", err
	}
	defer stmt.Close()

	_, err = stmt.Exec(account)
	if err != nil {
		return "DB deletion error ", err
	}

	return "DB deletion success.", nil

}

func DepositLedger(account string, token string, amount *big.Int) (bool, error) {

	db, err := connectDB()

	// sum balance
	balanceCheck, err := SumBalance(account, amount)
	// deposit
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

	log.Printf("Deposit affected %d rows", affect)
	return true, nil
}

func WithdrawLedger(account string, token string, amount *big.Int) (bool, error) {

	db, err := connectDB()

	// check if user has enough balance
	balanceCheck, err := HasEnoughBalance(account, amount)
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

func HasEnoughBalance(account string, amount *big.Int) (BalanceCheck, error) {

	db, err := connectDB()

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

func SumBalance(account string, amount *big.Int) (BalanceCheck, error) {
	db, err := connectDB()

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
	Check.NewBalance = (ledgerBalance.Add(ledgerBalance, amount)).String()

	return Check, nil
}
