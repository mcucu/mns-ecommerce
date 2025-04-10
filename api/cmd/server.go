package cmd

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mcucu/mns-ecommerce/internal/handlers"
	"github.com/mcucu/mns-ecommerce/internal/services"
	"github.com/mcucu/mns-ecommerce/internal/utils"
)

type IServer interface {
	StartApp()
}

type Server struct {
	utils.Options
	*services.Services
}

// NewServer create object server
func NewServer(opt utils.Options, svc *services.Services) IServer {
	return &Server{
		Options:  opt,
		Services: svc,
	}
}

func (s *Server) StartApp() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// e.Use(middleware.CORS())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderContentType, echo.HeaderAuthorization, echo.HeaderOrigin, echo.HeaderAccept, echo.HeaderAccessControlAllowOrigin, "X-API-Key"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPatch, http.MethodPut, http.MethodDelete, http.MethodOptions},
	}))

	e.Use(middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{
		Skipper: func(c echo.Context) bool {
			return strings.EqualFold(c.Request().RequestURI, "/") || strings.Contains(c.Request().RequestURI, "/swagger") || strings.Contains(c.Request().RequestURI, "/health") || strings.Contains(c.Request().RequestURI, "/metrics")
		},
		KeyLookup: "header:Authorization,header:X-API-Key",
		Validator: func(key string, c echo.Context) (bool, error) {
			return key == os.Getenv("APP_SECRET"), nil
		},
	}))

	handler := handlers.Options{
		Options:  s.Options,
		Services: s.Services,
	}

	// Register Routes
	Router(e, handler)

	// Start Server
	appPort := os.Getenv("PORT")
	if appPort == "" {
		appPort = "3000"
	}
	address := fmt.Sprintf(":%s", appPort)
	e.Logger.Fatal(e.Start(address))
}
