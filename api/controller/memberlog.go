package controller

import (
	"context"
	"fmt"
	"log"
	"server/api/protocol"
	"server/env"
	"server/service/mygrpc/protobuf"
	"server/service/pbclient"

	"github.com/gin-gonic/gin"
)

type (
	QueryLogReq struct {
		Account string `form:"account" json:"account" binding:"required"`
	}
	QueryLogStorage struct {
		Err error
	}

	QueryLogTask struct {
		Req     *QueryLogReq
		Res     *protocol.Response
		Storage *QueryLogStorage
	}
)

func NewQueryLogTask() *QueryLogTask {
	return &QueryLogTask{
		Req:     &QueryLogReq{},
		Res:     &protocol.Response{},
		Storage: &QueryLogStorage{},
	}
}

// @Summary Query User log
// @Description Query User log
// @Tags Get
// @Accept json
// @Produce json
// @Param account query string true "欲查詢歷程之使用者"
// @Success 200 "success"
// @Router /memberlog [get]
// Query User log
func Memberlog(c *gin.Context) {
	task := NewQueryLogTask()
	c.Set(env.APIResKeyInGinContext, task.Res)

	if shouldBreak := task.ShouldBind(c); shouldBreak {
		c.Error(task.Storage.Err)
		return
	}
	if shouldBreak := task.CheckUserExistByGRPC(c); shouldBreak {
		c.Error(task.Storage.Err)
		return
	}
	if shouldBreak := task.SortUserInfoByGRPC(c); shouldBreak {
		c.Error(task.Storage.Err)
		return
	}
}

func (task *QueryLogTask) ShouldBind(c *gin.Context) bool {
	//檢查input的資料key的型態與名稱正確性，Value不可為空
	if err := c.ShouldBindQuery(task.Req); err != nil {
		fmt.Println("ShouldBindQuery fault", err)
		task.Storage.Err = err
		return true
	}
	return false
}

func (task *QueryLogTask) CheckUserExistByGRPC(c *gin.Context) bool {
	grpcQueryLogIsExistRequest := &protobuf.QueryLogIsExistRequest{
		Account: task.Req.Account,
	}
	grpcLogIsExistResponse, err := pbclient.Client.QueryLogIsExist(context.Background(), grpcQueryLogIsExistRequest)
	if err != nil {
		log.Fatalf("gRPC client:error while calling QueryLogIsExist Service: %v \n", err)
		task.Storage.Err = err
		// task.Res.Code = 500
		// task.Res.Message = "Internal error"
		return true
	}
	log.Printf("gRPC client:Response from QueryLogIsExist Service: %v", grpcLogIsExistResponse)

	if !grpcLogIsExistResponse.Dataexist {
		err := fmt.Errorf("User not exist")
		task.Storage.Err = err
		// task.Res.Code = 401
		return true
	}
	return false
}

func (task *QueryLogTask) SortUserInfoByGRPC(c *gin.Context) bool {
	// grpc
	// login log check and add-on
	grpcQueryLogRequest := &protobuf.QueryLogRequest{
		Account: task.Req.Account,
	}
	grpcQueryResponse, err := pbclient.Client.QueryLog(context.Background(), grpcQueryLogRequest)
	if err != nil {
		log.Fatalf("gRPC client:error while calling QueryLog Service: %v \n", err)
		task.Storage.Err = err
		// task.Res.Code = 500
		// task.Res.Message = "Internal error"
		return true
	}
	log.Printf("gRPC client:Response from QueryLog Service: %v", grpcQueryResponse)

	task.Res.Result = grpcQueryResponse
	return false
}
