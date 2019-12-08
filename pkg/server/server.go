package server

import (
	"fmt"
	"net"

	"github.com/sirupsen/logrus"
	"github.com/wendellliu/good-search/pkg/common/dbAdapter"
	"github.com/wendellliu/good-search/pkg/config"
	"github.com/wendellliu/good-search/pkg/logger"
	pb "github.com/wendellliu/good-search/pkg/pb"
	"github.com/wendellliu/good-search/pkg/server/handlers"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"google.golang.org/grpc"
)

func Load(db dbAdapter.Database) {
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
	pb.RegisterGoodSearchServer(s, &handlers.Server{DB: db})

	logger.Logger.Infof("grpc server successfully connect to port %s", port)
	err = s.Serve(lis)
	if err != nil {
		logger.Logger.Fatalf("failed to serve: %v", err)
	}
}
