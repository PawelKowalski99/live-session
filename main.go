package main

import (
	"embed"
	"live-session-task/cmd"
)

//go:embed migrations/user/01_user_schema.sql
var embedMigrations embed.FS

// @title Cache-Database
// @version 1.0
// @description This is a cache and database example integration.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email pawelkowalski99@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
// @schemes http
func main() {
	cmd.Start(embedMigrations)
}
