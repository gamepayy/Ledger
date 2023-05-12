package v1

import (
	"encoding/json"
	"gamepayy_ledger/database"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Creates a new token
// @Produce json
// @Param address body string true "Address"
// @Param decimals body string true "Decimals"
// @Param symbol body string true "Symbol"
// @Param name body string true "Name"
// @Success 200 {object} object
// @Failure 400 {object} object
// @Router /token/new [post]
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

// @Summary 	Gets a token's data
// @Produce  	json
// @Param 		address query string true "Token"
// @Success 	200 {object} object
// @Failure 	400 {object} object
// @Router 		/token [get]
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

// @Summary Deletes a token
// @Param address query string true "Token"
// @Success 200 {object} object
// @Failure 400 {object} object
// @Router /token/delete [delete]
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

// @Summary Updates a token's data
// @Produce json
// @Param address body string true "Token"
// @Param decimals body string true "Decimals"
// @Param symbol body string true "Symbol"
// @Param name body string true "Name"
// @Success 200 {object} object
// @Failure 400 {object} object
// @Router /token/update [put]
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
