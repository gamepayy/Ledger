package main

import (
	router "gamepayy_ledger/routers"

	_ "github.com/go-sql-driver/mysql"
)

// @title GP API
// @version 1.0
// @description This API is used to manage the GP ledger. It is used to create, read, update and delete ledgers, to create, read, update and delete pending withdraws, and to create, read, update and delete tokens to the system.
// @host localhost:8080
// @BasePath /api/v1
func main() {
	router := router.InitRouter()
	router.Run("0.0.0.0:8080")
}
