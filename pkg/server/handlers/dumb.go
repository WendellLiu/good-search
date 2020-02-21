package handlers

import (
	"context"

	pb "github.com/wendellliu/good-search/pkg/pb"
)

func (s *Server) Dumb(ctx context.Context, req *pb.DumbReq) (*pb.DumbResp, error) {
	//localLogger := logger.Logger.WithFields(
	//logrus.Fields{"endpoint": "Dumb"},
	//)

	return &pb.DumbResp{}, nil
}
