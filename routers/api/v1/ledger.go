package v1

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	database "gamepayy_ledger/database"

	"github.com/gin-gonic/gin"
)

type Ledger struct {
	Account string `json:"account"`
	Balance string `json:"balance"`
}

// @Summary 	Gets an account's data
// @Produce  	json
// @Param 		account query string true "Account"
// @Success 	200 {object} object
// @Failure 	400 {object} object
// @Router 		/user [get]
func GetLedger(c *gin.Context) {

	query := c.Request.URL.Query()

	if query == nil {
		c.String(http.StatusBadRequest, "no query found")
	}

	ledger, err := database.GetLedger(query["account"][0])
	if err != nil {
		c.String(http.StatusBadRequest, "failed to get: %v", err)
	}
	c.JSON(http.StatusOK, ledger)

}

// @Summary 	Creates a new account with a ledger
// @Produce  	json
// @Success 	200 {object} object
// @Failure 	400 {object} object
// @Router 		/user/new [post]
func NewLedger(c *gin.Context) {

	body := c.Request.Body

	if body == nil {
		c.String(http.StatusBadRequest, "no body found")
		return
	}

	bodyBytes, _ := ioutil.ReadAll(c.Request.Body)

	jsonMap := make(map[string]interface{})
	err := json.Unmarshal(bodyBytes, &jsonMap)

	if err != nil {
		c.String(http.StatusBadRequest, "failed to unmarshal: %v", err)
		return
	}

	ledger := &database.Ledger{
		Account: jsonMap["account"].(string),
		Balance: jsonMap["balance"].(string),
	}

	result, err := database.NewLedger(*ledger)
	if err != nil {
		c.String(http.StatusBadRequest, "failed to create: %v", err)
	}
	c.JSON(http.StatusOK, result)
}

// @Summary Deletes an account
// @Produce  json
// @Param name body string true "Username" default(user)
// @Success 200 {object} object
// @Failure 400 {object} object
// @Router /user/delete [delete]
func DeleteLedger(c *gin.Context) {

	body := c.Request.Body

	if body == nil {
		c.String(http.StatusBadRequest, "no body found")
	}

	bodyBytes, _ := ioutil.ReadAll(c.Request.Body)

	jsonMap := make(map[string]interface{})
	err := json.Unmarshal(bodyBytes, &jsonMap)
	if err != nil {
		c.String(http.StatusBadRequest, "failed to unmarshal: %v", err)
	}

	result, err := database.DeleteLedger(jsonMap["account"].(string))

	if err != nil {
		c.String(http.StatusBadRequest, "failed to delete: %v", err)
	}
	c.JSON(http.StatusOK, result)
}

// @Summary Transfers an amount from one account to another
// @Produce  json
// @Param name body string true "from" default(user)
// @Param name body string true "to" default(user)
// @Param name body string true "amount" default(1)
// @Param name body string true "token" default(0x)
// @Success 200 {object} object
// @Failure 400 {object} object
// @Router /user/transfer [put]
func TransferLedger(c *gin.Context) {

	body := c.Request.Body
	if body == nil {
		c.String(http.StatusBadRequest, "no body found")
		return
	}

	bodyBytes, _ := ioutil.ReadAll(c.Request.Body)
	jsonMap := make(map[string]interface{})
	err := json.Unmarshal(bodyBytes, &jsonMap)
	if err != nil {
		c.String(http.StatusBadRequest, "failed to unmarshal: %v", err)
		return
	}

	fmt.Println("amount", jsonMap["amount"].(string))

	result, err := database.TransferLedger(jsonMap["from"].(string), jsonMap["to"].(string), jsonMap["amount"].(string), jsonMap["currency"].(string))
	if err != nil {
		c.String(http.StatusBadRequest, "failed to transfer: %v", err)
	}
	c.JSON(http.StatusOK, result)
}

// @Summary Withdraws an amount from an account
// @Produce  json
// @Param name body string true "account" default(user)
// @Param amount body string true "amount" default(1)
// @Param token body string true "token" default(0x)
// @Success 200 {object} object
// @Failure 400 {object} object
// @Router /user/withdraw [put]
func WithdrawLedger(c *gin.Context) {

	body := c.Request.Body
	if body == nil {
		c.String(http.StatusBadRequest, "no body found")
		return
	}

	bodyBytes, _ := ioutil.ReadAll(c.Request.Body)
	jsonMap := make(map[string]interface{})
	err := json.Unmarshal(bodyBytes, &jsonMap)
	if err != nil {
		c.String(http.StatusBadRequest, "failed to unmarshal: %v", err)
		return
	}

	result, err := database.WithdrawLedger(jsonMap["account"].(string), jsonMap["currency"].(string), jsonMap["amount"].(string))
	if err != nil {
		c.String(http.StatusBadRequest, "failed to withdraw: %v", err)
	}
	c.JSON(http.StatusOK, result)
}

// @Summary Deposits an amount to an account
// @Produce  json
// @Param name body string true "account" default(user)
// @Param amount body string true "amount" default(1)
// @Param token body string true "token" default(0x)
// @Success 200 {object} object
// @Failure 400 {object} object
// @Router /user/deposit [put]
func DepositLedger(c *gin.Context) {

	body := c.Request.Body
	if body == nil {
		c.String(http.StatusBadRequest, "no body found")
		return
	}

	bodyBytes, _ := ioutil.ReadAll(c.Request.Body)
	jsonMap := make(map[string]interface{})
	err := json.Unmarshal(bodyBytes, &jsonMap)
	if err != nil {
		c.String(http.StatusBadRequest, "failed to unmarshal: %v", err)
		return
	}

	result, err := database.DepositLedger(jsonMap["account"].(string), jsonMap["currency"].(string), jsonMap["amount"].(string))
	if err != nil {
		c.String(http.StatusBadRequest, "failed to deposit: %v", err)
	}
	c.JSON(http.StatusOK, result)
}
