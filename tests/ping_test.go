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

	var res health.Pong
	err := json.Unmarshal(response.Body.Bytes(), &res)
	assert.Nil(t, err)
	assert.Equal(t, res.Message, "pong", "メッセージ不一致")
}
