package server

import (
	"context"
	echoSwagger "github.com/swaggo/echo-swagger"
	app_user "live-session-task/core/application"
	_ "live-session-task/docs"
)

// cacheAndDbRoutes ...
func (es *EchoServer) cacheAndDbRoutes() {

	user := app_user.NewCrawlerHttpService(context.Background(), es.db, es.cache)

	es.GET("/users", user.Get)

	es.GET("/swagger/*", echoSwagger.WrapHandler)

}

// All routes
func (es *EchoServer) routes() {
	es.cacheAndDbRoutes()
}
