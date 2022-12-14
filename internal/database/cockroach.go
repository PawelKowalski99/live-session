package database

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/pressly/goose/v3"
	"live-session-task/env"
	"os"

	//	_ "github.com/lib/pq"
	"embed"
)

// DB Class
type CockRoach struct {
	ctx        context.Context
	config     env.EnvApp
	migrations embed.FS
}

// Ping
func (c *CockRoach) ping(err error, db *sql.DB) error {
	if err != nil {
		return err
	}

	// try ping
	err = db.Ping()
	if err != nil {
		return err
	}

	return nil
}

// Connect
func (c *CockRoach) connect() (*sql.DB, error) {
	// try open connection

	dbAddr := fmt.Sprintf(
		"%s://%s:%s@%s:%s/%s?sslmode=disable&options=%s",
		c.config.DB_ENGINE,
		c.config.DB_USERNAME,
		c.config.DB_PASSWORD,
		c.config.DB_HOST,
		c.config.DB_PORT,
		c.config.DB_DATABASE,
		c.config.DB_OPTIONS,
	)

	conn, err := pgx.Connect(context.Background(), dbAddr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		return nil, err
	}
	defer conn.Close(context.Background())

	db, err := sql.Open(c.config.DB_ENGINE, dbAddr)

	// try ping
	if c.ping(err, db) != nil {
		return nil, err
	}

	goose.SetBaseFS(c.migrations)

	if err := goose.SetDialect("postgres"); err != nil {
		return nil, err
	}

	if err := goose.Up(db, "migrations/user"); err != nil {
		return nil, err
	}

	return db, nil
}

// Get connection
func (c *CockRoach) ConnectDB() *sql.DB {
	counts := 0

	for {
		db, err := c.connect()

		if try(err, db, &counts) == nil {
			return db
		}
		continue
	}
}

// Constructor
func NewCockRoachDatabase(ctx context.Context, ec env.EnvApp, fs embed.FS) Database {
	return &CockRoach{ctx: ctx, config: ec, migrations: fs}
}
