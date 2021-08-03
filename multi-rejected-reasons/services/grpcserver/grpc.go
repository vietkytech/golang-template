package grpcserver

import (
	"net"

	"git.chotot.org/fse/multi-rejected-reasons/multi-rejected-reasons/config"
	"git.chotot.org/fse/multi-rejected-reasons/multi-rejected-reasons/handlers"
	"git.chotot.org/fse/multi-rejected-reasons/multi-rejected-reasons/proto/multirr"
	"git.chotot.org/go-common/kit/logger"
	"google.golang.org/grpc"
)

var log = logger.GetLogger("multirr-server")

// MultiRRServer must be embedded to have forward compatible implementations.
type MultiRRServer struct {
	handlers.MultiRRHandler
}

func NewRRServer(cfg *config.GrpcServerConfig) *MultiRRServer {
	lis, err := net.Listen(cfg.Network, cfg.Address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	rrserver := &MultiRRServer{}
	s := grpc.NewServer()
	multirr.RegisterMultiRRSvcServer(s, rrserver)

	log.Infof("grpc server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	return rrserver
}
