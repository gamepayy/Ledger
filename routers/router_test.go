package routers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	v1 "gamepayy_ledger/routers/api/v1"
)

func EndpointTest(t *testing.T, router *gin.Engine, requestType, endpoint string, responseCode int, body io.Reader, expectedResponse any) {
	t.Logf("Running test for endpoint: %s with request type: %s", endpoint, requestType) // This line is new
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(requestType, endpoint, body)

	router.ServeHTTP(w, req)

	assert.Equal(t, responseCode, w.Code)
	assert.Equal(t, expectedResponse, w.Body.String())

}

func setup() (*gin.Engine, *sync.Mutex) {
	// Setup
	router := InitRouter() // initialize your router here
	mu := &sync.Mutex{}    // initialize your mutex here

	return router, mu
}

func TestApiRoutes(t *testing.T) {
	router, mu := setup()

	mutexCycle(mu)

	// Clean up any existing data for account and token
	accountDeleteBody := strings.NewReader(`{"account":"tester16113"}`)
	EndpointTest(t, router, "DELETE", "/api/v1/user/delete", 200, accountDeleteBody, `{"message":"Account successfully deleted."}`)
	accountDeleteBody = strings.NewReader(`{"account":"tester1111155555"}`)
	EndpointTest(t, router, "DELETE", "/api/v1/user/delete", 200, accountDeleteBody, `{"message":"Account successfully deleted."}`)

	tokenDeleteBody := strings.NewReader(`{"address":"0x"}`)
	EndpointTest(t, router, "DELETE", "/api/v1/token/delete", 200, tokenDeleteBody, `"DB deletion success"`)

	// Create new account
	account := v1.Ledger{Account: "tester16113", Balance: "500000000000"}
	accountBytes, _ := json.Marshal(account)
	accountBody := bytes.NewReader(accountBytes)

	EndpointTest(t, router, "POST", "/api/v1/user/new", 200, accountBody, `{"message":"Account created successfully."}`)

	account = v1.Ledger{Account: "tester1111155555", Balance: "0"}
	accountBytes, _ = json.Marshal(account)
	accountBody = bytes.NewReader(accountBytes)
	EndpointTest(t, router, "POST", "/api/v1/user/new", 200, accountBody, `{"message":"Account created successfully."}`)
	EndpointTest(t, router, "GET", "/api/v1/user?account=tester16113", 200, nil, `{"account":"tester16113","balance":"500000000000"}`)

	// Create new token
	token := v1.Token{Address: "0x", Name: "Ether", Symbol: "ETH", Decimals: "18"}
	tokenBytes, _ := json.Marshal(token)
	tokenBody := bytes.NewReader(tokenBytes)

	EndpointTest(t, router, "POST", "/api/v1/token/new", 200, tokenBody, `true`)
	EndpointTest(t, router, "GET", "/api/v1/token?address=0x", 200, nil, `{"address":"0x","name":"Ether","symbol":"ETH","decimals":"18"}`)

	// Update token
	token.Symbol = "MATIC"
	tokenBytes, _ = json.Marshal(token)
	tokenBody = bytes.NewReader(tokenBytes)

	EndpointTest(t, router, "PUT", "/api/v1/token/update", 200, tokenBody, `true`)
	EndpointTest(t, router, "GET", "/api/v1/token?address=0x", 200, nil, `{"address":"0x","name":"Ether","symbol":"MATIC","decimals":"18"}`)

	mutexCycle(mu)

	// Transfer balance
	transfer := v1.TransferRequest{From: "tester16113", To: "tester1111155555", Amount: "50000000000", Currency: "0x"}
	transferBytes, _ := json.Marshal(transfer)
	transferBody := bytes.NewReader(transferBytes)

	EndpointTest(t, router, "PUT", "/api/v1/user/transfer", 200, transferBody, `true`)

	mutexCycle(mu)

	// Check balance
	EndpointTest(t, router, "GET", "/api/v1/user?account=tester16113", 200, nil, `{"account":"tester16113","balance":"450000000000"}`)
	EndpointTest(t, router, "GET", "/api/v1/user?account=tester1111155555", 200, nil, `{"account":"tester1111155555","balance":"50000000000"}`)

	// Deposit balance
	deposit := v1.LedgerChangeRequest{Account: "tester1111155555", Amount: "50000000000", Currency: "0x"}
	depositBytes, _ := json.Marshal(deposit)
	depositBody := bytes.NewReader(depositBytes)

	EndpointTest(t, router, "PUT", "/api/v1/user/deposit", 200, depositBody, `true`)

	// Check balance after deposit
	EndpointTest(t, router, "GET", "/api/v1/user?account=tester1111155555", 200, nil, `{"account":"tester1111155555","balance":"100000000000"}`)

	withdrawToDelete := v1.Withdraw{Account: "tester1111155555", Token: "0x"}
	withdrawToDeleteBytes, _ := json.Marshal(withdrawToDelete)
	withdrawToDeleteBody := bytes.NewReader(withdrawToDeleteBytes)

	// Delete withdraw
	EndpointTest(t, router, "DELETE", "/api/v1/withdraws/delete", 200, withdrawToDeleteBody, `{"message":"Withdrawal deleted successfully."}`)

	mutexCycle(mu)

	// Withdraw balance
	withdraw := v1.Withdraw{Account: "tester1111155555", Amount: "50000000000", Token: "0x"}
	withdrawBytes, _ := json.Marshal(withdraw)
	withdrawBody := bytes.NewReader(withdrawBytes)

	EndpointTest(t, router, "POST", "/api/v1/user/withdraw", 200, withdrawBody, `{"message":"Withdrawal successful."}`)

	// Check balance after withdrawal
	EndpointTest(t, router, "GET", "/api/v1/user?account=tester1111155555", 200, nil, `{"account":"tester1111155555","balance":"50000000000"}`)

}

// missing tests for pending withdraws api

func TestPendingWithdrawsRoutes(t *testing.T) {

	router, mu := setup()

	account := v1.Ledger{Account: "tester16113"}
	accountBytes, _ := json.Marshal(account)
	accountBody := bytes.NewReader(accountBytes)

	EndpointTest(t, router, "DELETE", "/api/v1/user/delete", 200, accountBody, `{"message":"Account successfully deleted."}`)

	mutexCycle(mu)

	account = v1.Ledger{Account: "tester16113", Balance: "500000000000"}
	accountBytes, _ = json.Marshal(account)
	accountBody = bytes.NewReader(accountBytes)
	EndpointTest(t, router, "POST", "/api/v1/user/new", 200, accountBody, `{"message":"Account created successfully."}`)

	mutexCycle(mu)

	accountResponse := `{"account":"tester16113","balance":"500000000000"}`
	EndpointTest(t, router, "GET", "/api/v1/user?account=tester16113", 200, nil, accountResponse)

	// Withdraw balance
	withdraw := v1.Withdraw{Account: "tester16113", Amount: "5000000000", Token: "0x"}
	withdrawBytes, _ := json.Marshal(withdraw)
	withdrawBody := bytes.NewReader(withdrawBytes)

	EndpointTest(t, router, "POST", "/api/v1/user/withdraw", 200, withdrawBody, `{"message":"Withdrawal successful."}`)

	// Delete withdraw
	withdraw = v1.Withdraw{Account: "tester16113", Token: "0x"}
	withdrawBytes, _ = json.Marshal(withdraw)
	withdrawBody = bytes.NewReader(withdrawBytes)
	EndpointTest(t, router, "DELETE", "/api/v1/withdraws/delete", 200, withdrawBody, `{"message":"Withdrawal deleted successfully."}`)

	// Create another withdraw
	withdraw = v1.Withdraw{Account: "tester16113", Amount: "150000000", Token: "1x"}
	withdrawBytes, _ = json.Marshal(withdraw)
	withdrawBody = bytes.NewReader(withdrawBytes)

	EndpointTest(t, router, "POST", "/api/v1/user/withdraw", 200, withdrawBody, `{"message":"Withdrawal successful."}`)

	mutexCycle(mu)

	// Process the withdraw
	withdraw = v1.Withdraw{Account: "tester16113", Token: "0x"}
	withdrawBytes, _ = json.Marshal(withdraw)
	withdrawBody = bytes.NewReader(withdrawBytes)
	EndpointTest(t, router, "PUT", "/api/v1/withdraws/process", 200, withdrawBody, `{"message":"Withdrawal processed successfully."}`)

	// Delete the processed withdraw
	withdraw = v1.Withdraw{Account: "tester16113", Token: "0x"}
	withdrawBytes, _ = json.Marshal(withdraw)
	withdrawBody = bytes.NewReader(withdrawBytes)
	EndpointTest(t, router, "DELETE", "/api/v1/withdraws/delete", 200, withdrawBody, `{"message":"Withdrawal deleted successfully."}`)

	// Delete account
	account = v1.Ledger{Account: "tester16113"}
	accountBytes, _ = json.Marshal(account)
	accountBody = bytes.NewReader(accountBytes)
	EndpointTest(t, router, "DELETE", "/api/v1/user/delete", 200, accountBody, `{"message":"Account successfully deleted."}`)
}

func TestRateLimitingMiddlewareOnGet(t *testing.T) {

	router, _ := setup()

	var wg sync.WaitGroup
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if i == 10 {
				EndpointTest(t, router, "GET", "/api/v1/user?account=tester1111155555", 429, nil, "Too many requests. Try again in 1s")
			}
			EndpointTest(t, router, "GET", "/api/v1/user?account=tester1111155555", 200, nil, `{"account":"tester1111155555","balance":"50000000000"}`)
		}()
	}
	wg.Wait() // Wait for all requests to finish

}

func mutexCycle(mutex *sync.Mutex) {
	mutex.Lock()
	time.Sleep(10 * time.Millisecond)
	mutex.Unlock()
}
