package controller

import (
	"fmt"
	"server/api/protocol"
	"server/env"
	"server/model/dao/logout"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type (
	LogoutReq struct {
		Jwt string `form:"jwt" json:"jwt" binding:"required"`
	}
	LogoutStorage struct {
		Err error
	}
	LogoutTask struct {
		Req     *LogoutReq
		Res     *protocol.Response
		Storage *LogoutStorage
	}
)

func NewLogoutTask() *LogoutTask {
	return &LogoutTask{
		Req:     &LogoutReq{},
		Res:     &protocol.Response{},
		Storage: &LogoutStorage{},
	}
}

func Logout(c *gin.Context) {
	task := NewLogoutTask()
	c.Set(env.APIResKeyInGinContext, task.Res)
	if shouldBreak := task.ShouldBind(c); shouldBreak {
		c.Error(task.Storage.Err)
		return
	}
	if shouldBreak := task.JwtIsExistAndDeleteIfExist(c); shouldBreak {
		c.Error(task.Storage.Err)
		return
	}
}

func (task *LogoutTask) ShouldBind(c *gin.Context) bool {
	//檢查input的資料key的型態與名稱正確性，Value不可為空
	if err := c.ShouldBindBodyWith(task.Req, binding.JSON); err != nil {
		fmt.Println("ShouldBindJSON fault", err)
		task.Storage.Err = err
		return true
	}
	return false
}

func (task *LogoutTask) JwtIsExistAndDeleteIfExist(c *gin.Context) bool {
	JwtIsExist := logout.JwtIsExistAndDeleteIfExist(task.Req.Jwt)
	if !JwtIsExist {
		task.Res.Code = 200
	} else {
		task.Res.Code = 201
	}
	task.Res.Message = "Logout Success"
	return false
}
