package controller

import (
	"server/api/protocol"
	"server/env"

	"github.com/gin-gonic/gin"
)

func NewResponseTask(c gin.Context) *protocol.Response {
	return &protocol.Response{}
}

//輸入不存在的Url會回傳的訊息
func NotFound(c *gin.Context) {
	task := NewResponseTask
	c.Set(env.APIResKeyInGinContext, task)

	return
}
