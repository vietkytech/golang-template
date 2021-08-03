package echoserver

import (
	"context"
	"fmt"
	"net/http"

	"git.chotot.org/fse/multi-rejected-reasons/multi-rejected-reasons/config"
	"git.chotot.org/fse/multi-rejected-reasons/multi-rejected-reasons/handlers"
	"git.chotot.org/fse/multi-rejected-reasons/multi-rejected-reasons/proto/multirr"
	"git.chotot.org/go-common/common-lib/echo/errorhandler"
	"git.chotot.org/go-common/echopprof"
	"git.chotot.org/go-common/echoprometheus"
	"git.chotot.org/go-common/kit/logger"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var log = logger.GetLogger("multirr-http-server")

type MultiRREchoConfig struct {
	EchoConfig *config.EchoServerConfig
	JWTConfig  *config.JWTConfig
}

// AuthPayload jwt payload
type AuthPayload struct {
	jwt.StandardClaims
	Typ string `json:"omitempty,typ"`
	// Vrf []string `json:"omitempty,vrf"`
}

type MultiRREchoServer struct {
	config  *MultiRREchoConfig
	handler *handlers.MultiRRHandler
}

func NewMultiRREchoServer(config *MultiRREchoConfig) *MultiRREchoServer {
	handler := handlers.NewRRHandler(&handlers.MultiRRHandlerConfig{})
	return &MultiRREchoServer{
		config:  config,
		handler: handler,
	}
}

func (h *MultiRREchoServer) health(ctx echo.Context) error {
	result, _ := h.handler.HealthCheck(context.Background(), &multirr.HealthCheckRequest{})
	return ctx.String(http.StatusOK, result.Msg)
}

func (h *MultiRREchoServer) StartServer() error {
	e := echo.New()
	e.Debug = h.config.EchoConfig.Debug
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize: 1 << 10, // 1 KB
	}))

	echopprof.Wrapper(e)
	echoprometheus.NewPrometheus("").Use(e)

	e.HTTPErrorHandler = errorhandler.NewErrorHandler(e.Debug)

	jwtConfig := middleware.JWTConfig{
		Claims:     &AuthPayload{},
		SigningKey: []byte(h.config.JWTConfig.Secret),
	}

	private := e.Group("/api/v1/private/")
	private.Use(middleware.JWTWithConfig(jwtConfig))

	e.GET("/health", h.health)

	err := e.Start(fmt.Sprintf(":%s", h.config.EchoConfig.Port))
	if err != nil {
		log.Fatalf("Error while starting Echo server: %s", err)
	}
	return nil
}
