package controller

import (
	"context"
	"fmt"
	"log"
	"server/api/protocol"
	"server/env"
	"server/model/redis"
	"server/service/mygrpc/protobuf"
	"server/service/pbclient"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type (
	LogDelReq struct {
		Account string `form:"account" Json:"account" binding:"required"`
	}
	LogDelStorage struct {
		Err error
	}
	LogDelTask struct {
		Req     *LogDelReq
		Res     *protocol.Response
		Storage *LogDelStorage
	}
)

func NewLogDelTask() *LogDelTask {
	return &LogDelTask{
		Req:     &LogDelReq{},
		Res:     &protocol.Response{},
		Storage: &LogDelStorage{},
	}
}

func Logdel(c *gin.Context) {
	task := NewLogDelTask()
	c.Set(env.APIResKeyInGinContext, task.Res)

	if shouldBreak := task.Logdel(c); shouldBreak {
		c.Error(task.Storage.Err)
		return
	}
	if shouldBreak := task.CheckUserExistByGRPC(c); shouldBreak {
		c.Error(task.Storage.Err)
		return
	}

}

func (task *LogDelTask) Logdel(c *gin.Context) bool {
	//檢查input的資料key的型態與名稱正確性，Value不可為空
	if err := c.ShouldBindBodyWith(task.Req, binding.JSON); err != nil {
		fmt.Println("ShouldBindJSON fault", err)
		task.Storage.Err = err
		return true
	}
	return false
}

func (task *LogDelTask) CheckUserExistByGRPC(c *gin.Context) bool {
	//grpc
	grpcRequest := &protobuf.DeleteLogIsExistRequest{
		Account: task.Req.Account,
	}
	grpcResponse, err := pbclient.Client.DeleteLogIsExist(context.Background(), grpcRequest)
	if err != nil {
		log.Fatalf("gRPC client:error while calling DeleteLogIsExist Service: %v \n", err)
		task.Storage.Err = err
		task.Res.Code = 500
		task.Res.Message = "Internal error"
		return true
	}
	log.Printf("gRPC client:Response from DeleteLogIsExist Service: %v", grpcResponse)

	if !grpcResponse.Dataexist {
		err := fmt.Errorf("Log or User not exist")
		task.Storage.Err = err
		// task.Res.Code = 401
		return true
	}
	redis.DeleteInfo(task.Req.Account)
	// task.Res.Code = 200
	// task.Res.Message = "success"
	return false
}
