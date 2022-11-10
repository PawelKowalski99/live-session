package main

import (
	"embed"
	"live-session-task/cmd"
)

//go:embed migrations/user/01_user_schema.sql
var embedMigrations embed.FS

// Call the entry point
func main() {
	cmd.Start(embedMigrations)
}
