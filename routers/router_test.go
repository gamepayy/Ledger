package routers

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func EndpointTest(t *testing.T, router *gin.Engine, requestType, endpoint string, responseCode int, body io.Reader, expectedResponse any) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(requestType, endpoint, body)

	router.ServeHTTP(w, req)

	assert.Equal(t, responseCode, w.Code)
	assert.Equal(t, expectedResponse, w.Body.String())

}

func TestAllRoutes(t *testing.T) {
	router := InitRouter()

	accountJson := `{"account":"tester16113","balance":"5"}`
	postBody := strings.NewReader(`{"account":"tester16113","balance":"5"}`)
	body := strings.NewReader(`{"account":"tester16113"}`)

	EndpointTest(t, router, "POST", "/api/v1/user/new", 200, postBody, `true`)
	EndpointTest(t, router, "GET", "/api/v1/user", 200, body, accountJson)

	time.Sleep(5 * time.Second)

	body = strings.NewReader(`{"account":"tester16113"}`)
	EndpointTest(t, router, "DELETE", "/api/v1/user/delete", 200, body, `"DB deletion success."`)

	body = strings.NewReader(`{"address":"0x","name":"Ether","symbol":"ETH","decimals":"18"}`)
	EndpointTest(t, router, "POST", "/api/v1/token/new", 200, body, `true`)

	body = strings.NewReader(`{"address":"0x"}`)
	EndpointTest(t, router, "GET", "/api/v1/token", 200, body, `{"address":"0x","name":"Ether","symbol":"ETH","decimals":"18"}`)

	body = strings.NewReader(`{"address":"0x","name":"Ether","symbol":"MATIC","decimals":"18" }`)
	EndpointTest(t, router, "PUT", "/api/v1/token/update", 200, body, `true`)

	body = strings.NewReader(`{"address":"0x"}`)
	EndpointTest(t, router, "GET", "/api/v1/token", 200, body, `{"address":"0x","name":"Ether","symbol":"MATIC","decimals":"18"}`)

	time.Sleep(5 * time.Second)
	body = strings.NewReader(`{"address":"0x"}`)
	EndpointTest(t, router, "DELETE", "/api/v1/token/delete", 200, body, `"DB deletion success"`)

	EndpointTest(t, router, "PUT", "/api/v1/user/transfer", 200, nil, `{"message":"pong"}`)
	EndpointTest(t, router, "PUT", "/api/v1/user/deposit", 200, nil, `{"message":"pong"}`)
	EndpointTest(t, router, "PUT", "/api/v1/user/withdraw", 200, nil, `{"message":"pong"}`)

}
