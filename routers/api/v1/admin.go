package v1

import (
	"net/http"

	database "gamepayy_ledger/database"

	"github.com/gin-gonic/gin"
)

// Token holds token information
// swagger:model Token
type Token struct {
	Address  string `json:"address"`
	Decimals string `json:"decimals"`
	Symbol   string `json:"symbol"`
	Name     string `json:"name"`
}

// @Summary Creates a new token
// @Produce json
// @Param body body Token true "Token"
// @Success 200 {object} bool "True"
// @Failure 400 {object} string "Bad request: error message"
// @Failure 500 {object} string "Internal server error: error message"
// @Router /token/new [post]
func NewToken(c *gin.Context) {
	var token database.Token
	if err := c.ShouldBindJSON(&token); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request: " + err.Error()})
		return
	}

	if _, err := database.NewToken(token); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, true)
}

// @Summary 	Gets a token's data
// @Produce  	json
// @Param 		address query string true "Token"
// @Success 	200 {object} Token
// @Failure 	400 {object} string "Bad request: no query found"
// @Failure 	404 {object} string "Token not found"
// @Router 		/token [get]
func GetToken(c *gin.Context) {
	address := c.Query("address")

	if address == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request: no query found"})
		return
	}

	token, err := database.GetToken(address)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Token not found"})
		return
	}

	c.JSON(http.StatusOK, token)
}

// @Summary Deletes a token
// @Param address query string true "Token"
// @Success 200 {object} string "DB deletion success"
// @Failure 400 {object} string "Bad request: error message"
// @Failure 500 {object} string "Internal server error: error message"
// @Router /token/delete [delete]
func DeleteToken(c *gin.Context) {
	var request struct {
		Address string `json:"address" binding:"required"`
	}

	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request: no body found"})
		return
	}

	if _, err := database.DeleteToken(request.Address); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, "DB deletion success")
}

// @Summary Updates a token's data
// @Produce json
// @Param body body Token true "Token"
// @Success 200 {object} bool "True"
// @Failure 400 {object} string "Bad request: error message"
// @Failure 500 {object} string "Internal server error: error message"
// @Router /token/update [put]
func UpdateToken(c *gin.Context) {
	var token database.Token
	if err := c.ShouldBindJSON(&token); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request: " + err.Error()})
		return
	}

	if _, err := database.UpdateToken(token); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, true)
}
