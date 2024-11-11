package main

import (
	"server/api"
	"server/infrastructure/service/pbserver"
)

func main() {
	// gRPC 啟動
	go pbserver.GrpcServer()
	// gin server 啟動
	api.Api()

}
