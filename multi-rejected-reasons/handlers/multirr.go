package handlers

import (
	"context"

	"git.chotot.org/fse/multi-rejected-reasons/multi-rejected-reasons/proto/multirr"
	"git.chotot.org/go-common/kit/logger"
)

var log = logger.GetLogger("multirr-handler")

type MultiRRHandlerConfig struct {
}

// MultiRRHandler must be embedded to have forward compatible implementations.
type MultiRRHandler struct {
	multirr.UnimplementedMultiRRSvcServer
}

func NewRRHandler(cfg *MultiRRHandlerConfig) *MultiRRHandler {
	return &MultiRRHandler{}
}

func (server *MultiRRHandler) HealthCheck(ctx context.Context, req *multirr.HealthCheckRequest) (*multirr.HealthCheckResponse, error) {
	return &multirr.HealthCheckResponse{
		Msg: "OK",
	}, nil
}

// func (server *MultiRRHandler) GetRRTemplate(context.Context, *GetRRTemplateRequest) (*GetRRTemplateResponse, error) {
// 	return nil, status.Errorf(codes.Unimplemented, "method GetRRTemplate not implemented")
// }
// func (server *MultiRRHandler) GetAllRRTemplates(context.Context, *GetAllRRTemplatesRequest) (*GetAllRRTemplatesResponse, error) {
// 	return nil, status.Errorf(codes.Unimplemented, "method GetAllRRTemplates not implemented")
// }
// func (server *MultiRRHandler) CreateRRTemplate(context.Context, *CreateRRTemplateRequest) (*CreateRRTemplateResponse, error) {
// 	return nil, status.Errorf(codes.Unimplemented, "method CreateRRTemplate not implemented")
// }
// func (server *MultiRRHandler) UpdateRRTemplate(context.Context, *UpdateRRTemplateRequest) (*UpdateRRTemplateResponse, error) {
// 	return nil, status.Errorf(codes.Unimplemented, "method UpdateRRTemplate not implemented")
// }
// func (server *MultiRRHandler) DeleteRRTemplate(context.Context, *DeleteRRTemplateRequest) (*DeleteRRTemplateResponse, error) {
// 	return nil, status.Errorf(codes.Unimplemented, "method DeleteRRTemplate not implemented")
// }
// func (server *MultiRRHandler) GetAdsRRs(context.Context, *GetAdsRRsRequest) (*GetAdsRRsResponse, error) {
// 	return nil, status.Errorf(codes.Unimplemented, "method GetAdsRRs not implemented")
// }
// func (server *MultiRRHandler) CheckAdsRRs(context.Context, *CheckAdsRRsRequest) (*CheckAdsRRsResponse, error) {
// 	return nil, status.Errorf(codes.Unimplemented, "method CheckAdsRRs not implemented")
// }
// func (server *MultiRRHandler) mustEmbedUnimplementedMultiRRHandler() {}
