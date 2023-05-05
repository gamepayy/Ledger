package database

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Token struct {
	Address  string `json:"address"`
	Name     string `json:"name"`
	Symbol   string `json:"symbol"`
	Decimals string `json:"decimals"`
}

func NewToken(token Token) (bool, error) {

	db, _ := connectDB()

	// insert
	stmt, err := db.Prepare("INSERT INTO tokens (address, name, symbol, decimals) VALUES (?, ?, ?, ?)")
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(token.Address, token.Name, token.Symbol, token.Decimals)
	if err != nil {

		fmt.Println(err)
		return false, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		fmt.Println(err)
		return false, err
	}

	log.Printf("New token created with insertion id: %d", id)
	return true, nil
}

func GetToken(tokenAddress string) (Token, error) {

	db, _ := connectDB()

	stmt, err := db.Prepare("SELECT * FROM tokens WHERE address = ?")
	if err != nil {
		return Token{}, err
	}
	defer stmt.Close()

	var token Token

	err = stmt.QueryRow(tokenAddress).Scan(&token.Address, &token.Name, &token.Symbol, &token.Decimals)
	if err != nil {
		return Token{}, err
	}

	return token, nil
}

func DeleteToken(token string) (string, error) {
	db, _ := connectDB()
	fmt.Println(token)
	stmt, err := db.Prepare("DELETE FROM tokens WHERE address = ?")
	if err != nil {
		return "DB statement preparation error ", err
	}
	defer stmt.Close()

	_, err = stmt.Exec(token)
	if err != nil {
		return "DB deletion error", err
	}

	return "DB deletion success", nil

}

func UpdateToken(token Token) (string, error) {

	db, _ := connectDB()

	// update
	stmt, err := db.Prepare("UPDATE tokens SET name = ?, symbol = ?, decimals = ? WHERE address = ?")
	if err != nil {
		return "Token update error ", err
	}
	defer stmt.Close()

	_, err = stmt.Exec(token.Name, token.Symbol, token.Decimals, token.Address)
	if err != nil {
		return "Token update error ", err
	}

	return "Token succesfully updated", nil
}
