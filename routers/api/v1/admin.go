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
func NewToken(c *gin.Context) {
	c.String(http.StatusOK, "pang")
}

// @Summary
// @Produce  json
// @Param
// @Success 200 {object} json
// @Failure 400 {object} json
// @Router /api/v1/ledger [get]
func DeleteToken(c *gin.Context) {
	c.String(http.StatusOK, "pang")
}

// @Summary
// @Produce  json
// @Param
// @Success 200 {object} json
// @Failure 400 {object} json
// @Router /api/v1/ledger [get]
func EditToken(c *gin.Context) {
	c.String(http.StatusOK, "pang")
}
