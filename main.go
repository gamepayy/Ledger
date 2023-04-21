package main

import (
	"database/sql"
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
	ledger := database.Ledger{
		Account: "tester",
		Balance: *big.NewInt(100),
	}

	_, err = database.NewLedger(db, ledger)

}
