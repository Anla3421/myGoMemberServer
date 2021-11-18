package controller

import (
	"crypto/md5"
	"fmt"
	"server/api/protocol"
	"server/env"
	"server/model/dao/registry"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type (
	RegistryReq struct {
		Account  string `form:"account" json:"account" binding:"required"`
		Password string `form:"password" json:"password" binding:"required"`
	}

	RegistryStorage struct {
		Err error
	}

	RegistryTask struct {
		Req     *RegistryReq
		Res     *protocol.Response
		Storage *RegistryStorage
	}
)

// NewRegistryTask:實體化Task
func NewRegistryTask() *RegistryTask {
	return &RegistryTask{
		Req:     &RegistryReq{},
		Res:     &protocol.Response{},
		Storage: &RegistryStorage{},
	}
}

// @Summary Registry a account
// @Description Give a ID & PWD to Registry
// @Tags Post
// @Accept json
// @Produce json
// @Param account/password body RegistryReq true "欲註冊之帳號及密碼"
// @Success 200 "success"
// @Router /registry [post]
// Registry
func Registry(c *gin.Context) {
	task := NewRegistryTask()
	c.Set(env.APIResKeyInGinContext, task.Res)

	if shouldBreak := task.ShouldBind(c); shouldBreak {
		c.Error(task.Storage.Err)
		return
	}
	if shouldBreak := task.AccountCheckAndResgistry(c); shouldBreak {
		c.Error(task.Storage.Err)
		return
	}
	return
}

// ShouldBind:參數解析
// 檢查input的資料key的型態與名稱正確性，Value不可為空
func (task *RegistryTask) ShouldBind(c *gin.Context) bool {
	if err := c.ShouldBindBodyWith(task.Req, binding.JSON); err != nil {
		fmt.Println("ShouldBindJSON fault", err)
		task.Storage.Err = err
		return true
	}
	return false
}

// AccountCheckAndResgistry:查詢重複註冊與否及新使用者註冊
func (task *RegistryTask) AccountCheckAndResgistry(c *gin.Context) bool {
	// 進DB查詢使用者是否存在
	accountIsExist := registry.IsExist(task.Req.Account)
	if accountIsExist {
		// task.Res.Code = 409
		err := fmt.Errorf("This account name is already used.")
		task.Storage.Err = err
		return true
	} else {
		//db
		hashedPassword := fmt.Sprintf("%x", md5.Sum([]byte(task.Req.Password))) //md5加密
		registry.RegMember(task.Req.Account, hashedPassword)
	}
	return false
}
