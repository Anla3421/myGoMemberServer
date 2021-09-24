package controller

import (
	"fmt"
	"server/api/protocol"
	"server/env"
	"server/model/dao/login"
	"server/model/dao/registry"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

var mySigningKey = []byte("iamsalt")

type (
	LoginReq struct {
		Account  string `form:"account" json:"account" binding:"required"`
		Password string `form:"password" json:"password" binding:"required"`
	}
	LoginStorage struct {
		Err error
		JWT string
	}
	LoginTask struct {
		Req            *LoginReq
		Res            *protocol.Response
		Storage        *LoginStorage
		MyCustomClaims *MyCustomClaims
	}

	MyCustomClaims struct {
		Account string
		jwt.StandardClaims
	}
)

func NewLoginTask() *LoginTask {
	return &LoginTask{
		Req:            &LoginReq{},
		Res:            &protocol.Response{},
		Storage:        &LoginStorage{},
		MyCustomClaims: &MyCustomClaims{},
	}
}

func Login(c *gin.Context) {
	task := NewLoginTask()
	c.Set(env.APIResKeyInGinContext, task.Res)

	if shouldBreak := task.ShouldBind(c); shouldBreak {
		c.Error(task.Storage.Err)
		return
	}
	if shouldBreak := task.CheckUserExist(c); shouldBreak {
		c.Error(task.Storage.Err)
		return
	}
	if shouldBreak := task.CheckPWD(c); shouldBreak {
		c.Error(task.Storage.Err)
		return
	}

	if shouldBreak := task.CreateJWT(c); shouldBreak {
		c.Error(task.Storage.Err)
		return
	}

	if shouldBreak := task.ExecuteLogin(c); shouldBreak {
		c.Error(task.Storage.Err)
		return
	}

}

func (task *LoginTask) ShouldBind(c *gin.Context) bool {
	//檢查input的資料key的型態與名稱正確性，Value不可為空
	if err := c.ShouldBindBodyWith(task.Req, binding.JSON); err != nil {
		fmt.Println("ShouldBindJSON fault", err)
		task.Storage.Err = err
		return true
	}
	return false
}

func (task *LoginTask) CheckUserExist(c *gin.Context) bool {
	//先判斷使用者是否存在，存在再判斷密碼是否正確
	Bool := registry.IsExist(task.Req.Account)
	if !Bool {
		task.Storage.Err = fmt.Errorf("Login failed, registery first, please")
		task.Res.Code = 403
		task.Res.Message = "Login failed, registery first, please"
		return true
	}
	return false
}

func (task *LoginTask) CheckPWD(c *gin.Context) bool {
	loginBool_Pwd := login.IsRight(task.Req.Account, task.Req.Password)
	if !loginBool_Pwd {
		task.Storage.Err = fmt.Errorf("Wrong password")
		task.Res.Code = 401
		task.Res.Message = "Wrong password"
		return true
	}
	return false
}

func (task *LoginTask) CreateJWT(c *gin.Context) bool {
	// Create the Claims
	now := time.Now()
	claims := MyCustomClaims{
		Account: task.Req.Account,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: now.Add(60 * 60 * 24 * time.Second).Unix(),
			Issuer:    "Me",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	if err != nil {
		task.Storage.Err = err
		// task.Res.Code = 400
		// task.Res.Message = "Unexpected failed" //JWT產生過程報錯
		return true
	}
	task.Storage.JWT = ss
	return false
}

func (task *LoginTask) ExecuteLogin(c *gin.Context) bool {
	//db
	login.LoginJwt(task.Storage.JWT, task.Req.Account)
	LoginTime := time.Now().Unix()
	login.Loginlog(task.Req.Account, LoginTime)

	// All Done, give response
	task.Res.Result = task.Storage.JWT
	return false
}
