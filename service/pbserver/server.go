package pbserver

import (
	"context"
	"fmt"
	"log"
	"net"
	"server/model/dao/logdel"
	"server/model/dao/memberlog"
	"server/model/dao/registry"
	"server/service/pbserver/getInfo"

	proto "github.com/Anla3421/myGoProtobuf/myGoMemberServer/go"
	protobuf "github.com/Anla3421/myGoProtobuf/myGoMemberServer/go"

	"google.golang.org/grpc"
)

//gPRC server 連線建立＆監聽
type Server struct {
	protobuf.UnimplementedMygrpcServiceServer
}

func setServer(s *grpc.Server) {
	proto.RegisterGetInfoServer(s, getInfo.NewService())
}

func GrpcServer() {
	fmt.Println("starting gRPC server...")

	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("gRPC server:failed to listen: %v \n", err)
	}

	grpcServer := grpc.NewServer()
	protobuf.RegisterMygrpcServiceServer(grpcServer, &Server{})

	setServer(grpcServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("gRPC server:failed to serve: %v \n", err)
	}
}

//gPRC server function
func (*Server) QueryLogIsExist(ctx context.Context, req *protobuf.QueryLogIsExistRequest) (*protobuf.QueryLogIsExistResponse, error) {
	fmt.Printf("gRPC server:func QueryLogIsExist is invoked with %v \n", req)

	res := &protobuf.QueryLogIsExistResponse{
		Dataexist: registry.IsExist(req.Account),
	}
	return res, nil
}

func (*Server) QueryLog(ctx context.Context, req *protobuf.QueryLogRequest) (*protobuf.QueryLogResponse, error) {
	fmt.Printf("gRPC server:func QueryLog is invoked with %v \n", req)

	res := &protobuf.QueryLogResponse{
		Account:   req.Account,
		Logintime: memberlog.QueryLog(req.Account),
		// Logintime: time.Now().Unix(),
	}
	return res, nil
}

func (*Server) DeleteLogIsExist(ctx context.Context, req *protobuf.DeleteLogIsExistRequest) (*protobuf.DeleteLogIsExistResponse, error) {
	fmt.Printf("gRPC server:func DeleteLogIsExist is invoked with %v \n", req)
	res := &protobuf.DeleteLogIsExistResponse{
		Dataexist: logdel.IFDeleteLog(req.Account),
	}
	return res, nil
}

func (*Server) DeleteLog(ctx context.Context, req *protobuf.DeleteLogRequest) (*protobuf.DeleteLogResponse, error) {
	fmt.Printf("gRPC server:func DeleteLog is invoked with %v \n", req)
	logdel.DeleteLog(req.Account)
	res := &protobuf.DeleteLogResponse{
		Account:   req.Account,
		Logintime: memberlog.QueryLog(req.Account),
	}
	return res, nil
}
