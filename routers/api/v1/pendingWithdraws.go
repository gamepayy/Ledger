package v1

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	database "gamepayy_ledger/database"

	"github.com/gin-gonic/gin"
)

type Withdraw struct {
	Account        string    `json:"account"`
	Token          string    `json:"token"`
	Amount         string    `json:"amount"`
	Pending        bool      `json:"pending"`
	IssueTimestamp time.Time `json:"issueTimestamp"`
}

// @Summary 	Inserts a pending withdraw into the database
// @Produce  	json
// @Param 		account body string true "Account"
// @Param 		token body string true "Token"
// @Param 		amount body string true "Amount"
// @Accept json
// @Success 	200 {object} object
// @Failure 	400 {object} object
// @Router 		/withdraws/new [post]
func InsertWithdraw(c *gin.Context) {
	body := c.Request.Body

	if body == nil {
		c.String(http.StatusBadRequest, "no body found")
	}

	bodyBytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.String(http.StatusBadRequest, "failed to read body: %v", err)
	}

	jsonMap := make(map[string]interface{})
	err = json.Unmarshal(bodyBytes, &jsonMap)
	if err != nil {
		c.String(http.StatusBadRequest, "failed to unmarshal: %v", err)
	}

	withdraws, err := database.InsertWithdraw(jsonMap["account"].(string), jsonMap["token"].(string), jsonMap["amount"].(string))
	if err != nil {
		c.String(http.StatusBadRequest, "failed to get: %v", err)
	}
	c.JSON(http.StatusOK, withdraws)
}

// @Summary 	Removes a pending withdraw from the database
// @Produce  	json
// @Param 		account body string true "Account"
// @Param 		token body string true "Token"
// @Accept json
// @Success 	200 {object} object
// @Failure 	400 {object} object
// @Router 		/withdraws/delete [delete]
func DeleteWithdraw(c *gin.Context) {
	body := c.Request.Body

	if body == nil {
		c.String(http.StatusBadRequest, "no body found")
	}

	bodyBytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.String(http.StatusBadRequest, "failed to read body: %v", err)
	}

	jsonMap := make(map[string]interface{})
	err = json.Unmarshal(bodyBytes, &jsonMap)
	if err != nil {
		c.String(http.StatusBadRequest, "failed to unmarshal: %v", err)
	}

	withdraws, err := database.DeleteWithdraw(jsonMap["account"].(string), jsonMap["token"].(string))
	if err != nil {
		c.String(http.StatusBadRequest, "failed to get: %v", err)
	}
	c.JSON(http.StatusOK, withdraws)
}

// @Summary 	Removes all finished withdraws from the database and adds them to the finished withdraws table
// @Produce  	json
// @Accept json
// @Success 	200 {object} object
// @Failure 	400 {object} object
// @Router 		/withdraws/clean [delete]
func DeleteProcessedWithdraws(c *gin.Context) {
	body := c.Request.Body

	if body == nil {
		c.String(http.StatusBadRequest, "no body found")
	}

	bodyBytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.String(http.StatusBadRequest, "failed to read body: %v", err)
	}

	jsonMap := make(map[string]interface{})
	err = json.Unmarshal(bodyBytes, &jsonMap)
	if err != nil {
		c.String(http.StatusBadRequest, "failed to unmarshal: %v", err)
	}

	withdraws, err := database.DeleteProcessedWithdraws()
	if err != nil {
		c.String(http.StatusBadRequest, "failed to get: %v", err)
	}
	c.JSON(http.StatusOK, withdraws)
}

// @Summary 	Sets a pending withdraw to finished
// @Produce  	json
// @Param 		account body string true "Account"
// @Param 		token body string true "Token"
// @Accept json
// @Success 	200 {object} object
// @Failure 	400 {object} object
// @Router 		/withdraws/process [put]
func ProcessWithdraw(c *gin.Context) {
	body := c.Request.Body

	if body == nil {
		c.String(http.StatusBadRequest, "no body found")
	}

	bodyBytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.String(http.StatusBadRequest, "failed to read body: %v", err)
	}

	jsonMap := make(map[string]interface{})
	err = json.Unmarshal(bodyBytes, &jsonMap)
	if err != nil {
		c.String(http.StatusBadRequest, "failed to unmarshal: %v", err)
	}

	withdraws, err := database.ProcessWithdraw(jsonMap["account"].(string), jsonMap["token"].(string))
	if err != nil {
		c.String(http.StatusBadRequest, "failed to get: %v", err)
	}
	c.JSON(http.StatusOK, withdraws)
}

// @Summary 	Gets an account's withdraws data
// @Produce  	json
// @Param 		account body string true "Account"
// @Param 		token body string true "Token"
// @Accept json
// @Success 	200 {object} object
// @Failure 	400 {object} object
// @Router 		/withdraws [get]
func GetWithdraws(c *gin.Context) {
	body := c.Request.Body

	if body == nil {
		c.String(http.StatusBadRequest, "no body found")
	}

	bodyBytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.String(http.StatusBadRequest, "failed to read body: %v", err)
	}

	jsonMap := make(map[string]interface{})
	err = json.Unmarshal(bodyBytes, &jsonMap)
	if err != nil {
		c.String(http.StatusBadRequest, "failed to unmarshal: %v", err)
	}

	withdraws, err := database.GetWithdraws(jsonMap["account"].(string), jsonMap["token"].(string))
	if err != nil {
		c.String(http.StatusBadRequest, "failed to get: %v", err)
	}
	c.JSON(http.StatusOK, withdraws)

}
