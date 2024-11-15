package controller

import (
	"fmt"
	"server/api/protocol"
	"server/domain/repository/model/dao/query"
	"server/domain/repository/model/dto"
	"server/domain/repository/redis"
	"server/env"

	"github.com/gin-gonic/gin"
)

type (
	QueryStorage struct {
		UserInfo dto.Response
		Err      error
	}
	QueryTask struct {
		Res     *protocol.Response
		Storage *QueryStorage
	}
)

// NewQueryTask:實體化task
func NewQueryTask() *QueryTask {
	return &QueryTask{
		Res:     &protocol.Response{},
		Storage: &QueryStorage{},
	}
}

// @Summary Query User
// @Description Query User Data
// @Tags Get
// @Accept json
// @Produce json
// @Param account path string true "欲查詢資料之使用者"
// @Success 200 "success"
// @Router /query/{account} [get]
// Query User
// Query:查詢使用者
func Query(c *gin.Context) {
	task := NewQueryTask()
	c.Set(env.APIResKeyInGinContext, task.Res)
	if shouldBreak := task.CheckUserInfo(c); shouldBreak {
		return
	}

	return
}

// CheckUserInfo:檢查使用者資料是否存在
func (task *QueryTask) CheckUserInfo(c *gin.Context) bool {
	account := c.Param("account")
	// redis 先用redis做查詢的動作，若無資料再往DB做查詢
	redisInfoIsExist, dbdata := redis.QueryInfo(account) //redis:true=有此人，false=沒此人
	if !redisInfoIsExist {
		// db
		dbInfoIsExist := query.QueryInfoIsExist(account) //db:true=有此人，false=沒此人
		if !dbInfoIsExist {
			err := fmt.Errorf("User not exist")
			task.Storage.Err = err
			// task.Res.Code = 401
			return true
		}
		dbdata = query.QueryInfo(account)
		dbdata.Account = account
	}
	// task.Res.Code = 200
	// task.Res.Message = "Success"
	task.Res.Result = dbdata
	return false
}
