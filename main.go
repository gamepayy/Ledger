package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/big"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"

	database "gamepayy_ledger/database"
)

func main() {

	// load
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("err loading: %v", err)
	}

	db, err := sql.Open("mysql", os.Getenv("DSN"))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("failed to ping: %v", err)
	}

	// insert a new ledger
	newLedger := database.Ledger{
		Account: "tester22",
		Balance: "1000000000000000000000000000000",
	}

	_, err = database.NewLedger(db, newLedger)

	if err != nil {
		log.Fatalf("failed to insert: %v", err)
	}

	withdraw := new(big.Int)
	withdraw, ok := withdraw.SetString("1000000000000000000000000000", 10)
	if !ok {
		log.Fatalf("failed to set string")
	}
	// withdraw from a ledger
	_, err = database.WithdrawLedger(db, "tester12", "token", withdraw)
	if err != nil {
		log.Fatalf("failed to withdraw: %v", err)
	}

	// get a ledger
	ledger, err := database.GetLedger(db, "tester12")
	if err != nil {
		log.Fatalf("failed to get: %v", err)
	}
	fmt.Println(ledger)

}
