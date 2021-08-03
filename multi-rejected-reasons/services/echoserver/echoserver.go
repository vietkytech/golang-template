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
	echo    *echo.Echo
}

func NewMultiRREchoServer(config *MultiRREchoConfig) *MultiRREchoServer {
	handler := handlers.NewRRHandler(&handlers.MultiRRHandlerConfig{})
	e := echo.New()
	return &MultiRREchoServer{
		config:  config,
		handler: handler,
		echo:    e,
	}
}

func (h *MultiRREchoServer) health(ctx echo.Context) error {
	result, _ := h.handler.HealthCheck(context.Background(), &multirr.HealthCheckRequest{})
	return ctx.String(http.StatusOK, result.Msg)
}

func (h *MultiRREchoServer) StartServer() error {
	h.echo.Debug = h.config.EchoConfig.Debug
	h.echo.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize: 1 << 10, // 1 KB
	}))

	echopprof.Wrapper(h.echo)
	echoprometheus.NewPrometheus("").Use(h.echo)

	h.echo.HTTPErrorHandler = errorhandler.NewErrorHandler(h.echo.Debug)

	jwtConfig := middleware.JWTConfig{
		Claims:     &AuthPayload{},
		SigningKey: []byte(h.config.JWTConfig.Secret),
	}

	private := h.echo.Group("/api/v1/private/")
	private.Use(middleware.JWTWithConfig(jwtConfig))

	h.echo.GET("/health", h.health)

	err := h.echo.Start(fmt.Sprintf(":%s", h.config.EchoConfig.Port))
	if err != nil {
		log.Fatalf("Error while starting Echo server: %s", err)
	}
	return nil
}

func (h *MultiRREchoServer) Shutdown(ctx context.Context) error {
	return h.echo.Shutdown(ctx)
}
