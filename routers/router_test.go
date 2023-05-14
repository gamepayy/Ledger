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

	accountJson := `{"account":"tester16113","balance":"500000000000"}`
	postBody := strings.NewReader(`{"account":"tester16113","balance":"500000000000"}`)
	body := strings.NewReader(`{"account":"tester16113"}`)

	EndpointTest(t, router, "POST", "/api/v1/user/new", 200, postBody, `true`)
	EndpointTest(t, router, "GET", "/api/v1/user", 200, body, accountJson)

	time.Sleep(1 * time.Second)

	body = strings.NewReader(`{"address":"0x","name":"Ether","symbol":"ETH","decimals":"18"}`)
	EndpointTest(t, router, "POST", "/api/v1/token/new", 200, body, `true`)

	body = strings.NewReader(`{"address":"0x"}`)
	EndpointTest(t, router, "GET", "/api/v1/token", 200, body, `{"address":"0x","name":"Ether","symbol":"ETH","decimals":"18"}`)

	body = strings.NewReader(`{"address":"0x","name":"Ether","symbol":"MATIC","decimals":"18" }`)
	EndpointTest(t, router, "PUT", "/api/v1/token/update", 200, body, `true`)

	body = strings.NewReader(`{"address":"0x"}`)
	EndpointTest(t, router, "GET", "/api/v1/token", 200, body, `{"address":"0x","name":"Ether","symbol":"MATIC","decimals":"18"}`)

	time.Sleep(1 * time.Second)
	body = strings.NewReader(`{"address":"0x"}`)
	EndpointTest(t, router, "DELETE", "/api/v1/token/delete", 200, body, `"DB deletion success"`)

	body = strings.NewReader(`{"account":"tester1111155555"}`)
	EndpointTest(t, router, "DELETE", "/api/v1/user/delete", 200, body, `"DB deletion success."`)

	body = strings.NewReader(`{"account":"tester1111155555", "balance":"0"}`)
	EndpointTest(t, router, "POST", "/api/v1/user/new", 200, body, `true`)

	body = strings.NewReader(`{"from":"tester16113","to":"tester1111155555","amount":"50000000000","currency":"0x"}`)
	EndpointTest(t, router, "PUT", "/api/v1/user/transfer", 200, body, `true`)

	time.Sleep(1 * time.Second)
	accountJson = `{"account":"tester16113","balance":"450000000000"}`
	body = strings.NewReader(`{"account":"tester16113"}`)
	EndpointTest(t, router, "GET", "/api/v1/user", 200, body, accountJson)

	accountJson = `{"account":"tester1111155555","balance":"50000000000"}`
	body = strings.NewReader(`{"account":"tester1111155555"}`)
	EndpointTest(t, router, "GET", "/api/v1/user", 200, body, accountJson)

	body = strings.NewReader(`{"account":"tester16113"}`)
	EndpointTest(t, router, "DELETE", "/api/v1/user/delete", 200, body, `"DB deletion success."`)

	body = strings.NewReader(`{"account":"tester1111155555","amount":"50000000000","currency":"0x"}`)
	EndpointTest(t, router, "PUT", "/api/v1/user/deposit", 200, body, `true`)

	accountJson = `{"account":"tester1111155555","balance":"100000000000"}`
	body = strings.NewReader(`{"account":"tester1111155555"}`)
	EndpointTest(t, router, "GET", "/api/v1/user", 200, body, accountJson)

	accountJson = `{"account":"tester1111155555","balance":"50000000000"}`
	body = strings.NewReader(`{"account":"tester1111155555","amount":"50000000000","currency":"0x"}`)
	EndpointTest(t, router, "PUT", "/api/v1/user/withdraw", 200, body, `true`)

	accountJson = `{"account":"tester1111155555","balance":"50000000000"}`
	body = strings.NewReader(`{"account":"tester1111155555"}`)
	EndpointTest(t, router, "GET", "/api/v1/user", 200, body, accountJson)

}

// missing tests for pending withdraws api

func TestPendingWithdrawsRoutes(t *testing.T) {

	router := InitRouter()

	accountJson := `{"account":"tester16113","balance":"500000000000"}`
	postBody := strings.NewReader(`{"account":"tester16113","balance":"500000000000"}`)
	body := strings.NewReader(`{"account":"tester16113"}`)
	EndpointTest(t, router, "DELETE", "/api/v1/user/delete", 200, body, `"DB deletion success."`)

	time.Sleep(1 * time.Second)
	EndpointTest(t, router, "POST", "/api/v1/user/new", 200, postBody, `true`)

	body = strings.NewReader(`{"account":"tester16113"}`)
	EndpointTest(t, router, "GET", "/api/v1/user", 200, body, accountJson)

	// 0. delete all pending withdraws

	body = strings.NewReader(`{"account":"tester16113", "token":"0x"}`)
	EndpointTest(t, router, "DELETE", "/api/v1/withdraws/delete", 200, body, `true`)
	body = strings.NewReader(`{"account":"tester16113", "token":"1x"}`)
	EndpointTest(t, router, "DELETE", "/api/v1/withdraws/delete", 200, body, `true`)
	body = strings.NewReader(`{"account":"tester16113", "token":"2x"}`)
	EndpointTest(t, router, "DELETE", "/api/v1/withdraws/delete", 200, body, `true`)

	time.Sleep(1 * time.Second)
	// 1. create a new pending withdraw

	body = strings.NewReader(`{"account":"tester16113","amount":"5000000000","token":"0x"}`)
	EndpointTest(t, router, "POST", "/api/v1/withdraws/new", 200, body, `true`)

	// 2. get this pending withdraw

	time.Sleep(3 * time.Second)

	body = strings.NewReader(`{"account":"tester16113"}`)
	EndpointTest(t, router, "GET", "/api/v1/withdraws", 200, body, `[{"Account":"tester16113","Token":"0x","Amount":"5000000000","Pending":"1"}]`)

	// 3. delete this pending withdraw

	body = strings.NewReader(`{"account":"tester16113", "token":"0x"}`)
	EndpointTest(t, router, "DELETE", "/api/v1/withdraws/delete", 200, body, `true`)

	// 4. get this pending withdraw again (should be deleted)

	body = strings.NewReader(`{"account":"tester16113"}`)
	EndpointTest(t, router, "GET", "/api/v1/withdraws", 200, body, `null`)

	// 5. create a new pending withdraw

	body = strings.NewReader(`{"account":"tester16113","amount":"150000000","token":"0x"}`)
	EndpointTest(t, router, "POST", "/api/v1/withdraws/new", 200, body, `true`)

	body = strings.NewReader(`{"account":"tester16113","amount":"20000000","token":"1x"}`)
	EndpointTest(t, router, "POST", "/api/v1/withdraws/new", 200, body, `true`)

	body = strings.NewReader(`{"account":"tester16113","amount":"3000000","token":"2x"}`)
	EndpointTest(t, router, "POST", "/api/v1/withdraws/new", 200, body, `true`)

	// 6. get all pending withdraws

	body = strings.NewReader(`{"account":"tester16113"}`)
	EndpointTest(t, router, "GET", "/api/v1/withdraws", 200, body, "[{\"Account\":\"tester16113\",\"Token\":\"0x\",\"Amount\":\"150000000\",\"Pending\":\"1\"},{\"Account\":\"tester16113\",\"Token\":\"1x\",\"Amount\":\"20000000\",\"Pending\":\"1\"},{\"Account\":\"tester16113\",\"Token\":\"2x\",\"Amount\":\"3000000\",\"Pending\":\"1\"}]")

	// 7. process 1x pending withdraws

	body = strings.NewReader(`{"account":"tester16113","token":"0x"}`)
	EndpointTest(t, router, "PUT", "/api/v1/withdraws/process", 200, body, `true`)

	// 8. get all pending withdraws again

	body = strings.NewReader(`{"account":"tester16113", "token":"0x"}`)
	EndpointTest(t, router, "GET", "/api/v1/withdraws", 200, body, `[{"Account":"tester16113","Token":"0x","Amount":"150000000","Pending":"0"},{"Account":"tester16113","Token":"1x","Amount":"20000000","Pending":"1"},{"Account":"tester16113","Token":"2x","Amount":"3000000","Pending":"1"}]`)

	// 9. delete all processed withdraws

	body = strings.NewReader(`{"account":"tester16113", "token":"0x"}`)
	EndpointTest(t, router, "DELETE", "/api/v1/withdraws/clean", 200, body, `true`)

	// 10. get all pending withdraws again

	body = strings.NewReader(`{"account":"tester16113", "token":"0x"}`)
	EndpointTest(t, router, "GET", "/api/v1/withdraws", 200, body, "[{\"Account\":\"tester16113\",\"Token\":\"1x\",\"Amount\":\"20000000\",\"Pending\":\"1\"},{\"Account\":\"tester16113\",\"Token\":\"2x\",\"Amount\":\"3000000\",\"Pending\":\"1\"}]")

	// 11. delete all pending withdraws

	body = strings.NewReader(`{"account":"tester16113", "token":"0x"}`)
	EndpointTest(t, router, "DELETE", "/api/v1/withdraws/clean", 200, body, `true`)

	// 12. delete user

	body = strings.NewReader(`{"account":"tester16113"}`)
	EndpointTest(t, router, "DELETE", "/api/v1/user/delete", 200, body, `"DB deletion success."`)

}
