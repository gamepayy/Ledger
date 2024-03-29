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
	if err != nil {
		return false, err
	}

	// insert
	stmt, err := db.Prepare("INSERT INTO users (address, balance) VALUES (?, ?)")
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(ledger.Account, ledger.Balance)
	if err != nil {
		return false, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return false, err
	}

	log.Printf("New ledger created with insertion id: %d", id)
	return true, nil
}

func GetLedger(account string) (Ledger, error) {

	db, err := connectDB()
	if err != nil {
		return Ledger{}, err
	}

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
	if err != nil {
		return "Error: ", err
	}

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

func DepositLedger(account, token, amount string) (bool, error) {

	db, err := connectDB()
	if err != nil {
		return false, err
	}

	bigAmount := new(big.Int)
	bigAmount, _ = bigAmount.SetString(amount, 10)

	// sum balance
	balanceCheck, err := sumBalance(account, bigAmount)
	if err != nil {
		return false, err
	}

	if token == "0x" {

	} else {
		// inserts into tokens table
	}
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

func TransferLedger(from, to, amount, token string) (bool, error) {

	db, err := connectDB()
	if err != nil {
		return false, err
	}

	bigAmount := new(big.Int)
	bigAmount, ok := bigAmount.SetString(amount, 10)
	if !ok {
		return false, fmt.Errorf("failed to convert string to big.Int")
	}

	// check if user has enough balance
	balanceCheck, err := hasEnoughBalance(from, bigAmount)
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

	_, err = stmt.Exec(balanceCheck.NewBalance, from)
	if err != nil {
		return false, err
	}

	receiverBalance, _ := sumBalance(to, bigAmount)
	// deposit
	stmt, err = db.Prepare("UPDATE users SET balance = ? WHERE address = ?")
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(receiverBalance.NewBalance, to)
	if err != nil {
		return false, err
	}

	return true, nil
}

func WithdrawLedger(account, token, amount string) (bool, error) {

	db, err := connectDB()
	if err != nil {
		return false, err
	}

	bigAmount := new(big.Int)
	bigAmount, _ = bigAmount.SetString(amount, 10)

	// check if user has enough balance
	balanceCheck, err := hasEnoughBalance(account, bigAmount)
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

	// add to pending withdrawals
	_, err = InsertWithdraw(account, token, amount)
	if err != nil {
		return false, err
	}

	return true, nil
}

type BalanceCheck struct {
	NewBalance string `json:"balance"`
	IsPossible bool   `json:"is_possible"`
}

func hasEnoughBalance(account string, amount *big.Int) (BalanceCheck, error) {

	db, err := connectDB()
	if err != nil {
		return BalanceCheck{}, err
	}

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
	Check.NewBalance = (big.NewInt(0).Sub(ledgerBalance, amount)).String()

	return Check, nil
}

func sumBalance(account string, amount *big.Int) (BalanceCheck, error) {
	db, err := connectDB()
	if err != nil {
		return BalanceCheck{}, err
	}

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
	Check.IsPossible = true
	Check.NewBalance = (big.NewInt(0).Add(ledgerBalance, amount)).String()

	return Check, nil
}
