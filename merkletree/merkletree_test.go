package merkletreeGen

import (
	"fmt"
	"log"
	"testing"
)

func TestGenerator(t *testing.T) {
	accounts := [4]string{"0x123", "0x456", "0x789", "0xabc"}
	tokens := [4]string{"0x111", "0x222", "0x333", "0x444"}
	amounts := [4]string{"100", "200", "300", "400"}
	withdraws := []Withdraw{Withdraw{Account: accounts[0], Token: tokens[0], Amount: amounts[0]}, Withdraw{Account: accounts[1], Token: tokens[1], Amount: amounts[1]}, Withdraw{Account: accounts[2], Token: tokens[2], Amount: amounts[2]}, Withdraw{Account: accounts[3], Token: tokens[3], Amount: amounts[3]}}

	fmt.Println("Withdraws: ", withdraws)

	tree, err := generateTree(withdraws)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Tree: ", tree)
}
