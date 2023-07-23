package database

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Withdraw struct {
	Id      int
	Account string
	Token   string
	Amount  string
	Pending string
}

// insert
func InsertWithdraw(account, token, amount string) (bool, error) {
	db, err := connectDB()
	if err != nil {
		return false, err
	}

	// insert
	stmt, err := db.Prepare("INSERT INTO pending_withdrawals (address, token_address, amount, pending) VALUES (?, ?, ?, ?)")
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(account, token, amount, true)
	if err != nil {
		return false, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return false, err
	}

	log.Printf("New pending withdraw created with insertion id: %d", id)
	return true, nil

}

// delete
func DeleteWithdraw(account, token string) (bool, error) {

	db, err := connectDB()
	if err != nil {
		return false, err
	}

	// insert
	stmt, err := db.Prepare("DELETE FROM pending_withdrawals WHERE address = ? AND token_address = ? AND pending = 1")
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(account, token)
	if err != nil {
		return false, err
	}

	_, err = res.LastInsertId()
	if err != nil {
		return false, err
	}

	log.Printf("Deleted withdraw succesfully")
	return true, nil
}

// delete
func DeleteProcessedWithdraws() (bool, error) {

	db, err := connectDB()
	if err != nil {
		return false, err
	}

	// insert
	stmt, err := db.Prepare("DELETE FROM pending_withdrawals WHERE pending = false")
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	res, err := stmt.Exec()
	if err != nil {
		return false, err
	}

	_, err = res.LastInsertId()
	if err != nil {
		return false, err
	}

	log.Printf("Deleted withdraws succesfully")
	return true, nil
}

// put
func ProcessWithdraw(account, token string) (bool, error) {

	db, err := connectDB()
	if err != nil {
		return false, err
	}

	// insert
	stmt, err := db.Prepare("UPDATE pending_withdrawals set pending = false WHERE address = ? AND token_address = ?")
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(account, token)
	if err != nil {
		return false, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return false, err
	}

	log.Printf("Withdraw processed with id: %d", id)
	return true, nil

}

// get
func GetWithdraws(account string) ([]Withdraw, error) {

	var withdraws []Withdraw
	db, err := connectDB()
	if err != nil {
		return []Withdraw{}, err
	}

	// insert
	stmt, err := db.Prepare("SELECT * FROM pending_withdrawals WHERE address = ?")
	if err != nil {
		return []Withdraw{}, err
	}
	defer stmt.Close()
	res, err := stmt.Query(account)
	if err != nil {
		return []Withdraw{}, err
	}

	for res.Next() {
		var row Withdraw
		err = res.Scan(&row.Id, &row.Account, &row.Token, &row.Amount, &row.Pending)
		if err != nil {
			return []Withdraw{}, err
		} else {
			withdraws = append(withdraws, row)
		}
	}

	return withdraws, nil

}
