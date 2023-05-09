package v1

import (
	"encoding/json"
	"gamepayy_ledger/database"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary
// @Param name body string true "Username" default(user) name body string true "Username" default(user)
// @Success 200 {object} object
// @Failure 400 {object} object
// @Router /api/v1/ledger [get]
func NewToken(c *gin.Context) {

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

	token := &database.Token{
		Address:  jsonMap["address"].(string),
		Decimals: jsonMap["decimals"].(string),
		Symbol:   jsonMap["symbol"].(string),
		Name:     jsonMap["name"].(string),
	}

	_, err = database.NewToken(*token)
	if err != nil {
		c.String(http.StatusBadRequest, "failed to create: %v", err)
	}

	c.JSON(http.StatusOK, true)
}

func GetToken(c *gin.Context) {

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

	response, err := database.GetToken(jsonMap["address"].(string))
	if err != nil {
		c.String(http.StatusBadRequest, "Token not found: %v", err)
	}

	c.JSON(http.StatusOK, response)
}

// @Summary
// @Param name body string true "Username" default(user)
// @Success 200 {object} object
// @Failure 400 {object} object
// @Router /api/v1/ledger [get]
func DeleteToken(c *gin.Context) {

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

	response, err := database.DeleteToken(jsonMap["address"].(string))
	if err != nil {
		c.String(http.StatusBadRequest, "failed to delete: %v", err)
	}

	c.JSON(http.StatusOK, response)

}

// @Summary
// @Success 200 {object} object
// @Failure 400 {object} object
// @Router /api/v1/ledger [get]
func UpdateToken(c *gin.Context) {

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

	token := database.Token{
		Address:  jsonMap["address"].(string),
		Decimals: jsonMap["decimals"].(string),
		Symbol:   jsonMap["symbol"].(string),
		Name:     jsonMap["name"].(string),
	}

	_, err = database.UpdateToken(token)
	if err != nil {
		c.String(http.StatusBadRequest, "failed to create: %v", err)
	}

	c.JSON(http.StatusOK, true)
}
