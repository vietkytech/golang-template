package services

import (
	"context"
	"net"

	"git.chotot.org/fse/multi-rejected-reasons/multi-rejected-reasons/config"
	"git.chotot.org/fse/multi-rejected-reasons/multi-rejected-reasons/proto/multirr"
	"git.chotot.org/go-common/kit/logger"
	"google.golang.org/grpc"
)

var log = logger.GetLogger("multirr-service")

type MultiRRConfig struct {
	GrpcConfig *config.GrpcServerConfig
}

// MultiRRSvcServer must be embedded to have forward compatible implementations.
type MultiRRSvcServer struct {
	multirr.UnimplementedMultiRRSvcServer
}

func NewRRServer(cfg *MultiRRConfig) *MultiRRSvcServer {
	lis, err := net.Listen(cfg.GrpcConfig.Network, cfg.GrpcConfig.Address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	rrserver := MultiRRSvcServer{}
	s := grpc.NewServer()
	multirr.RegisterMultiRRSvcServer(s, &rrserver)

	log.Infof("grpc server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	return &rrserver
}

func (server *MultiRRSvcServer) HealthCheck(ctx context.Context, req *multirr.HealthCheckRequest) (*multirr.HealthCheckResponse, error) {
	return &multirr.HealthCheckResponse{
		Msg: "OK",
	}, nil
}

// func (server *MultiRRSvcServer) GetRRTemplate(context.Context, *GetRRTemplateRequest) (*GetRRTemplateResponse, error) {
// 	return nil, status.Errorf(codes.Unimplemented, "method GetRRTemplate not implemented")
// }
// func (server *MultiRRSvcServer) GetAllRRTemplates(context.Context, *GetAllRRTemplatesRequest) (*GetAllRRTemplatesResponse, error) {
// 	return nil, status.Errorf(codes.Unimplemented, "method GetAllRRTemplates not implemented")
// }
// func (server *MultiRRSvcServer) CreateRRTemplate(context.Context, *CreateRRTemplateRequest) (*CreateRRTemplateResponse, error) {
// 	return nil, status.Errorf(codes.Unimplemented, "method CreateRRTemplate not implemented")
// }
// func (server *MultiRRSvcServer) UpdateRRTemplate(context.Context, *UpdateRRTemplateRequest) (*UpdateRRTemplateResponse, error) {
// 	return nil, status.Errorf(codes.Unimplemented, "method UpdateRRTemplate not implemented")
// }
// func (server *MultiRRSvcServer) DeleteRRTemplate(context.Context, *DeleteRRTemplateRequest) (*DeleteRRTemplateResponse, error) {
// 	return nil, status.Errorf(codes.Unimplemented, "method DeleteRRTemplate not implemented")
// }
// func (server *MultiRRSvcServer) GetAdsRRs(context.Context, *GetAdsRRsRequest) (*GetAdsRRsResponse, error) {
// 	return nil, status.Errorf(codes.Unimplemented, "method GetAdsRRs not implemented")
// }
// func (server *MultiRRSvcServer) CheckAdsRRs(context.Context, *CheckAdsRRsRequest) (*CheckAdsRRsResponse, error) {
// 	return nil, status.Errorf(codes.Unimplemented, "method CheckAdsRRs not implemented")
// }
// func (server *MultiRRSvcServer) mustEmbedUnimplementedMultiRRSvcServer() {}
