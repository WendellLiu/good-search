package server

import (
	"context"
	"fmt"
	"net"

	"github.com/sirupsen/logrus"
	"github.com/wendellliu/good-search/pkg/common/dbAdapter"
	"github.com/wendellliu/good-search/pkg/config"
	"github.com/wendellliu/good-search/pkg/dto"
	"github.com/wendellliu/good-search/pkg/logger"
	pb "github.com/wendellliu/good-search/pkg/pb"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedGoodSearchServer
	db dbAdapter.Database
}

func (s *server) GetExperience(ctx context.Context, req *pb.GetExperienceReq) (*pb.GetExperienceResp, error) {
	experience, err := dto.GetExperience(context.Background(), s.db, req.Id)

	if err != nil {
		logger.Logger.Error(err)
	}
	//logger.Logger.WithFields(logrus.Fields{"experience": fmt.Sprintf("%+v", experience)}).Info("get result")
	return &pb.GetExperienceResp{Type: experience.Type}, nil
}

func Server(db dbAdapter.Database) {
	port := fmt.Sprintf(":%s", config.Config.GRPCPort)

	lis, err := net.Listen("tcp", port)
	if err != nil {
		logger.Logger.Fatalf("failed to listen %s", port)
	}
	loggerEntry := logrus.NewEntry(logger.Logger)
	s := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_logrus.UnaryServerInterceptor(loggerEntry),
		)),
	)
	pb.RegisterGoodSearchServer(s, &server{db: db})

	logger.Logger.Infof("grpc server successfully connect to port %s", port)
	err = s.Serve(lis)
	if err != nil {
		logger.Logger.Fatalf("failed to serve: %v", err)
	}
}
