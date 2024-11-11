package getInfo

import (
	"context"
	"fmt"
	"server/domain/repository/model/dao/query"

	proto "github.com/Anla3421/myGoProtobuf/myGoMemberServer/go"
)

// rpc GetMemberIDByJWT(GetMemberIDByJWTReq) returns (GetMemberIDByJWTRes)
func (*Server) GetMemberIDByJWT(ctx context.Context, req *proto.GetMemberIDByJWTReq) (*proto.GetMemberIDByJWTRes, error) {
	fmt.Printf("gRPC server:func QueryLogIsExist is invoked with %v \n", req)

	data, err := query.GetInfoByJWT(req.JWT)
	if err != nil {
		return nil, err
	}
	res := &proto.GetMemberIDByJWTRes{
		ID:      data.ID,
		Account: data.Account,
	}
	return res, nil
}
