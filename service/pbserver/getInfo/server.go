package getInfo

import (
	proto "github.com/Anla3421/myGoProtobuf/myGoMemberServer/go"
)

// Server :
type Server struct {
	proto.UnimplementedGetInfoServer
}

// NewService : 產生服務
func NewService() *Server {
	return &Server{}
}
