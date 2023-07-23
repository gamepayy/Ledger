package v1

import (
	"net/http"

	database "gamepayy_ledger/database"

	"github.com/gin-gonic/gin"
)

type Ledger struct {
	Account string `json:"account"`
	Balance string `json:"balance"`
}

type TransferRequest struct {
	From     string `json:"from"`
	To       string `json:"to"`
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}

type LedgerChangeRequest struct {
	Account  string `json:"account"`
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}

// @Summary 	Gets an account's data
// @Produce  	json
// @Param 		account query string true "Account"
// @Success 	200 {object} Ledger
// @Failure 	400 {object} string "Bad request: no query found"
// @Failure 	404 {object} string "Account not found"
// @Router 		/user [get]
func GetLedger(c *gin.Context) {
	account := c.Query("account")
	if account == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request: no query found"})
		return
	}

	ledger, err := database.GetLedger(account)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Account not found"})
		return
	}

	c.JSON(http.StatusOK, ledger)
}

// @Summary 	Creates a new account with a ledger
// @Produce  	json
// @Param 		account body Ledger true "Account details"
// @Success 	200 {object} Ledger
// @Failure 	400 {object} string "Bad request: no body found"
// @Failure 	500 {object} string "Internal server error: error message"
// @Router 		/user/new [post]
func NewLedger(c *gin.Context) {
	var ledger database.Ledger

	if err := c.ShouldBindJSON(&ledger); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request: no body found"})
		return
	}

	_, err := database.NewLedger(ledger)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Account created successfully."})
}

// @Summary Deletes an account
// @Produce  json
// @Param account body string true "Account"
// @Success 200 {object} string "Account successfully deleted"
// @Failure 500 {object} string "Internal server error: error message"
// @Router /user/delete [delete]
func DeleteLedger(c *gin.Context) {
	var request struct {
		Account string `json:"account" binding:"required"`
	}

	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request: no body found"})
		return
	}

	if _, err := database.DeleteLedger(request.Account); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Account successfully deleted."})
}

// @Summary Transfers an amount from one account to another
// @Produce  json
// @Param request body TransferRequest true "Transfer details"
// @Success 200 {object} object
// @Failure 400 {object} object
// @Router /user/transfer [put]
func TransferLedger(c *gin.Context) {
	var request TransferRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request: " + err.Error()})
		return
	}

	result, err := database.TransferLedger(request.From, request.To, request.Amount, request.Currency)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to transfer: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

// @Summary Withdraws an amount from an account
// @Produce  json
// @Param request body LedgerChangeRequest true "Withdraw details"
// @Success 200 {object} object
// @Failure 400 {object} object
// @Router /user/withdraw [put]
func WithdrawLedger(c *gin.Context) {
	var request LedgerChangeRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request: " + err.Error()})
		return
	}

	_, err := database.WithdrawLedger(request.Account, request.Currency, request.Amount)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to withdraw: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Withdrawal successful."})
}

// @Summary Deposits an amount to an account
// @Produce  json
// @Param request body LedgerChangeRequest true "Deposit details"
// @Success 200 {object} object
// @Failure 400 {object} object
// @Router /user/deposit [put]
func DepositLedger(c *gin.Context) {
	var request LedgerChangeRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request: " + err.Error()})
		return
	}

	result, err := database.DepositLedger(request.Account, request.Currency, request.Amount)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to deposit: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}
