package controller

import (
	"crypto/md5"
	"fmt"
	"server/api/protocol"
	"server/domain/repository/model/dao/chpwd"
	"server/domain/repository/model/dao/logout"
	"server/env"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type (
	ChPWDReq struct {
		Jwt         string `form:"jwt" json:"jwt" binding:"required"`
		NewPassword string `form:"newpassword" json:"newpassword" binding:"required"`
	}
	ChPWDStorage struct {
		Err error
	}
	ChPWDTask struct {
		Req     *ChPWDReq
		Res     *protocol.Response
		Storage *ChPWDStorage
	}
)

func NewChPWDTask() *ChPWDTask {
	return &ChPWDTask{
		Req:     &ChPWDReq{},
		Res:     &protocol.Response{},
		Storage: &ChPWDStorage{},
	}
}

// @Summary Change Password
// @Description Change Password
// @Tags Post
// @Accept json
// @Produce json
// @Param Jwt/NewPassword body ChPWDReq true "JWT及新密碼"
// @Success 200 "success"
// @Router /chpwd [post]
// Change Password
func ChangePwd(c *gin.Context) {
	task := NewChPWDTask()
	c.Set(env.APIResKeyInGinContext, task.Res)

	if shouldBreak := task.ShouldBind(c); shouldBreak {
		c.Error(task.Storage.Err)
		return
	}
	if shouldBreak := task.ChPWD(c); shouldBreak {
		c.Error(task.Storage.Err)
		return
	}
}

func (task *ChPWDTask) ShouldBind(c *gin.Context) bool {
	//檢查input的資料key的型態與名稱正確性，Value不可為空
	if err := c.ShouldBindBodyWith(task.Req, binding.JSON); err != nil {
		fmt.Println("ShouldBindJSON fault", err)
		task.Storage.Err = err
		return true
	}
	return false
}

func (task *ChPWDTask) ChPWD(c *gin.Context) bool {
	account := chpwd.GetAccount(task.Req.Jwt)
	fmt.Println("輸入的JWT : " + task.Req.Jwt + "\t新密碼 : " + task.Req.NewPassword + "\t查詢到的帳號 : " + account)
	jwtIsExist := logout.JwtIsExistAndDeleteIfExist(task.Req.Jwt)
	if !jwtIsExist {
		err := fmt.Errorf("check your input info")
		task.Storage.Err = err
		// task.Res.Code = 401
		return true
	}

	//db
	password := fmt.Sprintf("%x", md5.Sum([]byte(task.Req.NewPassword)))
	fmt.Println("輸入的帳號 : " + account + "\t密碼 : " + password)
	chpwd.Chpwd(password, account)
	task.Res.Code = 200
	task.Res.Message = "Password changed, use new password to login again"

	return false
}
