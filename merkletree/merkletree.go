package merkletreeGen

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"log"
	"sync"

	"encoding/hex"

	"github.com/cbergoon/merkletree"
)

type Withdraw struct {
	Account        string
	Token          string
	Amount         string
	WithdrawString string
}

// CalculateHash hashes the values of a TestContent
func (w Withdraw) CalculateHash() ([]byte, error) {
	h := sha256.New()
	w.WithdrawString = fmt.Sprintf("%s%s%s", w.Account, w.Token, w.Amount)

	if _, err := h.Write([]byte(w.WithdrawString)); err != nil {
		return nil, err
	}

	return h.Sum(nil), nil
}

// Equals tests for equality of two Contents
func (w Withdraw) Equals(other merkletree.Content) (bool, error) {
	otherTC, ok := other.(Withdraw)
	if !ok {
		return false, errors.New("value is not of type TestContent")
	}
	return w.WithdrawString == otherTC.WithdrawString, nil
}

func generateTree(withdraws []Withdraw) (*merkletree.MerkleTree, error) {
	//Build list of Content to build tree
	var list []merkletree.Content

	wg := sync.WaitGroup{}
	wg.Add(len(withdraws))
	c := sync.Mutex{}
	for _, withdraw := range withdraws {

		go func() {

			c.Lock()
			list = append(list, Withdraw{WithdrawString: withdraw.Account + withdraw.Token + withdraw.Amount})
			c.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println("List: ", list)

	//Create a new Merkle Tree from the list of Content
	t, err := merkletree.NewTree(list)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(list[0].CalculateHash())

	//Get the Merkle Root of the tree
	mr := t.MerkleRoot()
	log.Println(mr)
	// convert mr to hash string
	mrStr := hex.EncodeToString(mr)
	log.Println(mrStr)

	//Verify the entire tree (hashes for each node) is valid
	vt, err := t.VerifyTree()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Verify Tree: ", vt)

	//Verify all  contents are in the tree
	wg = sync.WaitGroup{}
	wg.Add(len(list))
	for _, content := range list {
		contentCpy := content
		go func() {
			vc, err := t.VerifyContent(contentCpy)
			if err != nil {
				log.Fatal(err)
			}
			log.Println("Verify Content: ", vc)
			wg.Done()
		}()
	}
	wg.Wait()

	//String representation

	fmt.Println("Tree: ", t)
	fmt.Println("Tree.Root: ", t.Root)
	fmt.Println("Tree.Root.Hash: ", t.Root.Hash)
	fmt.Println("Tree Leafs: ", t.Leafs)
	fmt.Println("Tree Root Parent Hash: ", t.Leafs[0].Parent.Hash)
	return t, nil
}
