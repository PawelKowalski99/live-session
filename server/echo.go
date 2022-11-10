package server

import (
	"context"
	"database/sql"
	"github.com/labstack/echo/v4"
	"live-session-task/internal/cache"
	"net/http"

	"github.com/labstack/echo/v4/middleware"
)

// Micro framework class
type EchoServer struct {
	*echo.Echo
	ctx   context.Context
	db    *sql.DB
	cache cache.Cache
	port  string
}

// Configure echo
func (es *EchoServer) configure() {
	// console output
	es.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format:           "${time_custom} : ${method} -> ${uri}, status=${status} ::${error}\n",
		CustomTimeFormat: "15:04:05.00000",
	}))

	// recover from panic
	es.Use(middleware.Recover())

	// cors
	es.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost}, // you can add more for RESTFUL
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))

	// any auth middleware/microservice here
}

// Run server with Echo
func (es *EchoServer) Run() {
	es.Logger.Fatal(es.Start(":" + es.port))
}

// New Server instance
func NewEchoServer(ctx context.Context,
	db *sql.DB,
	appPort string,
	rdb cache.Cache,
) Server {
	if appPort == "" {
		appPort = "8080"
	}

	server := &EchoServer{
		echo.New(),
		ctx,
		db,
		rdb,
		appPort,
	}
	server.configure()
	server.routes()

	return server
}
