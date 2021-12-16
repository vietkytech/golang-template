package grpcserver

import (
	"net"

	"github.com/vietkytech/golang-template/golang-template/config"
	"github.com/vietkytech/golang-template/golang-template/handlers"
	"github.com/vietkytech/golang-template/golang-template/proto/multirr"
	"git.chotot.org/go-common/kit/logger"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"google.golang.org/grpc"
)

var log = logger.GetLogger("multirr-grpc-server")

type MultiRRServerConfig struct {
	GrpcConfig *config.GrpcServerConfig
}

func NewRRServer(cfg *MultiRRServerConfig, rrserver *handlers.MultiRRHandler) error {
	lis, err := net.Listen(cfg.GrpcConfig.Network, cfg.GrpcConfig.Address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer(
		grpc.StreamInterceptor(grpc_prometheus.StreamServerInterceptor),
		grpc.UnaryInterceptor(grpc_prometheus.UnaryServerInterceptor),
	)
	multirr.RegisterMultiRRSvcServer(s, rrserver)

	log.Infof("grpc server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		return err
	}
	log.Infof("grpc server started at %v", lis.Addr())

	grpc_prometheus.Register(s)

	return nil
}
