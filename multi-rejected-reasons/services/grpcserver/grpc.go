package grpcserver

import (
	"net"

	"git.chotot.org/fse/multi-rejected-reasons/multi-rejected-reasons/config"
	"git.chotot.org/fse/multi-rejected-reasons/multi-rejected-reasons/handlers"
	"git.chotot.org/fse/multi-rejected-reasons/multi-rejected-reasons/proto/multirr"
	"git.chotot.org/go-common/kit/logger"
	"google.golang.org/grpc"
)

var log = logger.GetLogger("multirr-grpc-server")

type MultiRRServerConfig struct {
	Config     *handlers.MultiRRHandlerConfig
	GrpcConfig *config.GrpcServerConfig
}

func NewRRServer(cfg *MultiRRServerConfig) multirr.MultiRRSvcServer {
	lis, err := net.Listen(cfg.GrpcConfig.Network, cfg.GrpcConfig.Address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	rrserver := handlers.NewRRHandler(cfg.Config)
	s := grpc.NewServer()
	multirr.RegisterMultiRRSvcServer(s, rrserver)

	log.Infof("grpc server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	return rrserver
}
