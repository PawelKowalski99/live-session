package server

import (
	"context"
	app_user "live-session-task/core/application"
)

// a TODO Routes
func (es *EchoServer) estatesRoutes() {

	user := app_user.NewCrawlerHttpService(context.Background(), es.db, es.cache)

	es.GET("/users", user.Get)

}

// All routes
func (es *EchoServer) routes() {
	es.estatesRoutes()
}
