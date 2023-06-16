package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func SaveHashedPassword(password []byte) ([]byte, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	fmt.Println("Hash to store:", string(hashedPassword))
	fmt.Println(hashedPassword)
	return hashedPassword, nil
}

func CheckPassword(hashedPassword, password []byte) (bool, error) {

	err := bcrypt.CompareHashAndPassword(hashedPassword, password)
	if err != nil {
		fmt.Println("Password was incorrect")
		return false, err
	}
	fmt.Println("Password was correct")
	return true, nil
}

func main() {
	hashed, err := SaveHashedPassword([]byte("aaaaa"))
	if err != nil {
		fmt.Println(err)
	}
	CheckPassword(hashed, []byte("aaaaaa"))
}
