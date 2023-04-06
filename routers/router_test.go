package routers

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

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

	EndpointTest(t, router, "GET", "/api/v1/user", 200, nil, "pang")

	EndpointTest(t, router, "POST", "/api/v1/new_token", 200, nil, `pang`)
	EndpointTest(t, router, "POST", "/api/v1/user/new", 200, nil, `{"message":"pong"}`)

	EndpointTest(t, router, "DELETE", "/api/v1/delete_token", 200, nil, `pang`)
	EndpointTest(t, router, "DELETE", "/api/v1/user/delete", 200, nil, `{"message":"pong"}`)

	EndpointTest(t, router, "PUT", "/api/v1/edit_token", 200, nil, `pang`)
	EndpointTest(t, router, "PUT", "/api/v1/user/transfer", 200, nil, `{"message":"pong"}`)
	EndpointTest(t, router, "PUT", "/api/v1/user/deposit", 200, nil, `{"message":"pong"}`)
	EndpointTest(t, router, "PUT", "/api/v1/user/withdraw", 200, nil, `{"message":"pong"}`)

}
