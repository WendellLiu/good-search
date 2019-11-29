package server

import (
	"fmt"
	"net"

	"github.com/wendellliu/good-search/pkg/config"
	"github.com/wendellliu/good-search/pkg/logger"
	pb "github.com/wendellliu/good-search/pkg/pb"

	"github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedGoodSearchServer
}

func Server() {
	port := fmt.Sprintf(":%s", config.Config.GRPCPort)

	lis, err := net.Listen("tcp", port)
	if err != nil {
		logger.Logger.Fatalf("failed to listen %s", port)
	}
	s := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_logrus.UnaryServerInterceptor(logger.Logger),
		)),
	)
	pb.RegisterGoodSearchServer(s, &server{})

	logger.Logger.Infof("grpc server successfully connect to port %s", port)

	err = s.Serve(lis)
	if err != nil {
		logger.Logger.Fatalf("failed to serve: %v", err)
	}
}
