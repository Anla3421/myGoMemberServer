package main

import (
	"server/api"
	"server/service/pbserver"
)

func main() {
	// gRPC 啟動
	go pbserver.GrpcServer()
	// gin server 啟動
	api.Api()

}
