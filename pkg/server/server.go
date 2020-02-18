package server

import (
	"fmt"
	"net"

	"github.com/sirupsen/logrus"
	"github.com/wendellliu/good-search/pkg/config"
	"github.com/wendellliu/good-search/pkg/dto"
	"github.com/wendellliu/good-search/pkg/es"
	"github.com/wendellliu/good-search/pkg/logger"
	pb "github.com/wendellliu/good-search/pkg/pb"
	"github.com/wendellliu/good-search/pkg/server/handlers"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"google.golang.org/grpc"
)

type Dependencies struct {
	Repo dto.DTO
	Es   es.Elasticsearch
}

func Load(dependencies Dependencies) {
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
	pb.RegisterGoodSearchServer(s, &handlers.Server{
		Repository: dependencies.Repo,
		Es:         dependencies.Es,
	})

	logger.Logger.Infof("grpc server successfully connect to port %s", port)
	err = s.Serve(lis)
	if err != nil {
		logger.Logger.Fatalf("failed to serve: %v", err)
	}
}
