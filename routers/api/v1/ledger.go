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
// @Param 		account body string true "Account"
// @Accept json
// @Success 	200 {object} object
// @Failure 	400 {object} object
// @Router 		/api/v1/user [get]
func GetLedger(c *gin.Context) {

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

	ledger, err := database.GetLedger(jsonMap["account"].(string))
	if err != nil {
		c.String(http.StatusBadRequest, "failed to get: %v", err)
	}
	c.JSON(http.StatusOK, ledger)

}

// @Summary 	Creates a new account with a ledger
// @Produce  	json
// @Success 	200 {object} object
// @Failure 	400 {object} object
// @Router 		/api/v1/ledger [post]
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

// @Summary
// @Produce  json

// @Success 200 {object} object
// @Failure 400 {object} object
// @Router /api/v1/ledger [delete]
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

// @Summary
// @Produce  json

// @Success 200 {object} object
// @Failure 400 {object} object
// @Router /api/v1/ledger [put]
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

// @Summary
// @Produce  json

// @Success 200 {object} object
// @Failure 400 {object} object
// @Router /api/v1/ledger [put]
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

// @Summary
// @Produce  json

// @Success 200 {object} object
// @Failure 400 {object} object
// @Router /api/v1/ledger [put]
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
