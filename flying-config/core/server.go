package core

import (
	"flying-config/api"
	"flying-config/global"
	"flying-config/middleware"
	grpc_recovery "flying-config/middleware"
	pb "flying-config/proto/app"
	pb0 "flying-config/proto/client"
	pb1 "flying-config/proto/namespace"
	pb2 "flying-config/proto/release"
	"fmt"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
	"time"
)

type server interface {
	ListenAndServe() error
}

func RunServer() {
	listen, err := net.Listen("tcp", global.GVA_CONFIG.System.Address)
	if err != nil {
		global.LOG.Error("Failed to listen: %v", zap.Any("err", err))
	}
	address := global.GVA_CONFIG.System.Address

	grpcServer := grpc.NewServer(middleware.TLSInterceptor(),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_validator.UnaryServerInterceptor(),

			grpc_recovery.UnaryServerInterceptor(grpc_recovery.RecoveryInterceptor()),
		)),
	)
	pb.RegisterAppServiceServer(grpcServer, &api.App{})
	pb0.RegisterClientServiceServer(grpcServer, &api.Client{})
	pb1.RegisterNamespaceServiceServer(grpcServer, &api.Namespace{})
	pb2.RegisterReleaseServiceServer(grpcServer, &api.Release{})
	RunStopNotify(grpcServer) //发送信号给admin
	time.Sleep(1000)
	global.LOG.Info("server run success on ", zap.String("address", address))
	fmt.Printf(`
	启动完成✈️
	地址:%s
`, address)
	if err = grpcServer.Serve(listen); err != nil {
		global.LOG.Error("ListenAndServe: ", zap.Any("err", err))
	}

}
