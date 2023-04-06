package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary
// @Produce  json
// @Param
// @Success 200 {object} json
// @Failure 400 {object} json
// @Router /api/v1/ledger [get]
func GetLedger(c *gin.Context) {
	c.String(http.StatusOK, "pang")
}

// @Summary
// @Produce  json
// @Param
// @Success 200 {object} json
// @Failure 400 {object} json
// @Router /api/v1/ledger [post]
func NewLedger(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

// @Summary
// @Produce  json
// @Param
// @Success 200 {object} json
// @Failure 400 {object} json
// @Router /api/v1/ledger [delete]
func DeleteLedger(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

// @Summary
// @Produce  json
// @Param
// @Success 200 {object} json
// @Failure 400 {object} json
// @Router /api/v1/ledger [put]
func TransferLedger(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

// @Summary
// @Produce  json
// @Param
// @Success 200 {object} json
// @Failure 400 {object} json
// @Router /api/v1/ledger [put]
func DepositLedger(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

// @Summary
// @Produce  json
// @Param
// @Success 200 {object} json
// @Failure 400 {object} json
// @Router /api/v1/ledger [put]
func WithdrawLedger(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
