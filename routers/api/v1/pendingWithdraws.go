package v1

import (
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
// @Param body body Withdraw true "Withdraw"
// @Success 	200 {object} bool "True"
// @Failure 	400 {object} string "Bad request: error message"
// @Failure 	500 {object} string "Internal server error: error message"
// @Router 		/withdraws/new [post]
func InsertWithdraw(c *gin.Context) {
	var withdraw Withdraw
	if err := c.ShouldBindJSON(&withdraw); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request: " + err.Error()})
		return
	}

	if _, err := database.InsertWithdraw(withdraw.Account, withdraw.Token, withdraw.Amount); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, true)
}

// @Summary 	Removes a pending withdraw from the database
// @Produce  	json
// @Param body body Withdraw true "Withdraw"
// @Success 	200 {object} bool "True"
// @Failure 	400 {object} string "Bad request: error message"
// @Failure 	500 {object} string "Internal server error: error message"
// @Router 		/withdraws/delete [delete]
func DeleteWithdraw(c *gin.Context) {
	var request struct {
		Account string `json:"account" binding:"required"`
		Token   string `json:"token" binding:"required"`
	}

	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request: no body found"})
		return
	}

	if _, err := database.DeleteWithdraw(request.Account, request.Token); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, true)
}

// @Summary 	Removes all finished withdraws from the database and adds them to the finished withdraws table
// @Produce  	json
// @Success 	200 {object} bool "True"
// @Failure 	500 {object} string "Internal server error: error message"
// @Router 		/withdraws/clean [delete]
func DeleteProcessedWithdraws(c *gin.Context) {
	if _, err := database.DeleteProcessedWithdraws(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, true)
}

// @Summary 	Sets a pending withdraw to finished
// @Produce  	json
// @Param body body Withdraw true "Withdraw"
// @Success 	200 {object} bool "True"
// @Failure 	400 {object} string "Bad request: error message"
// @Failure 	500 {object} string "Internal server error: error message"
// @Router 		/withdraws/process [put]
func ProcessWithdraw(c *gin.Context) {
	var request struct {
		Account string `json:"account" binding:"required"`
		Token   string `json:"token" binding:"required"`
	}

	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request: no body found"})
		return
	}

	if _, err := database.ProcessWithdraw(request.Account, request.Token); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, true)
}

// @Summary 	Gets an account's withdraws data
// @Produce  	json
// @Param account query string true "Account"
// @Success 	200 {array} Withdraw
// @Failure 	400 {object} string "Bad request: error message"
// @Failure 	500 {object} string "Internal server error: error message"
// @Router 		/withdraws [get]
func GetWithdraws(c *gin.Context) {
	account := c.Query("account")

	if account == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request: no account provided"})
		return
	}

	withdraws, err := database.GetWithdraws(account)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, withdraws)
}
