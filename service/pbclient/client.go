package pbclient

import (
	"fmt"
	"log"

	protobuf "github.com/Anla3421/myGoProtobuf/myGoMemberServer/go"

	"google.golang.org/grpc"
)

//gPRC client 連線建立
func init() {
	fmt.Println("gRPC client initial")
	CreateConn()
}

var Client protobuf.MygrpcServiceClient

func CreateConn() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	Client = protobuf.NewMygrpcServiceClient(conn)
}

// //gPRC client function
// func QueryLogIsExist(account string) *protobuf.QueryLogResponse {
// 	req := &protobuf.QueryLogRequest{
// 		Account: account,
// 	}
// 	res, err := Client.QueryLogIsExist(context.Background(), req)
// 	if err != nil {
// 		log.Fatalf("gRPC client:error while calling QueryLogIsExist Service: %v \n", err)
// 	}
// 	log.Printf("gRPC client:Response from QueryLogIsExist Service: %v", res)
// 	return res
// }

// func QueryLog(account string) *protobuf.QueryLogResponse {
// 	req := &protobuf.Request{
// 		Account: account,
// 	}
// 	res, err := Client.QueryLog(context.Background(), req)
// 	if err != nil {
// 		log.Fatalf("gRPC client:error while calling QueryLog Service: %v \n", err)
// 	}
// 	log.Printf("gRPC client:Response from QueryLog Service: %v", res)
// 	return res
// }

// func DeleteLogIsExist(account string) *protobuf.DeleteLogResponse {
// 	req := &protobuf.Request{
// 		Account: account,
// 	}
// 	res, err := Client.DeleteLogIsExist(context.Background(), req)
// 	if err != nil {
// 		log.Fatalf("gRPC client:error while calling DeleteLogIsExist Service: %v \n", err)
// 	}
// 	log.Printf("gRPC client:Response from DeleteLogIsExist Service: %v", res)
// 	return res
// }

// func DeleteLog(account string) {
// 	req := &protobuf.Request{
// 		Account: account,
// 	}
// 	res, err := Client.DeleteLog(context.Background(), req)
// 	if err != nil {
// 		log.Fatalf("gRPC client:error while calling DeleteLog Service: %v \n", err)
// 	}
// 	log.Printf("gRPC client:Response from DeleteLog Service: %v", res)
// }
