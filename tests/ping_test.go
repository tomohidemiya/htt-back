package main

import (
	health "app/backend/controllers"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestPingRouter(t *testing.T) {
	response := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(response)

	context.Request, _ = http.NewRequest(http.MethodGet, "/ping", nil)
	health.CheckHertBeat(context)

	assert.Equal(t, response.Code, 200, "API エラー発生")

	var resBody health.Pong
	err := json.Unmarshal(response.Body.Bytes(), &resBody)
	assert.Nil(t, err)
	assert.Equal(t, resBody.Message, "pong", "レスポンスメッセージ不一致")
}
