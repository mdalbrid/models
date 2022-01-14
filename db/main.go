package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/mdalbrid/utils/config"
	"github.com/mdalbrid/utils/logger"
)

const (
	ErrNotFound = "object not found"
)

var (
	Ctx    context.Context
	cancel context.CancelFunc
	Pool   *pgxpool.Pool
)

func ReposInit(config *config.DatabasePostgresConfig) {
	Ctx, cancel = context.WithCancel(context.Background())

	logger.Info("Connect to database")

	connStr := fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%d sslmode=disable ",
		config.User, config.Password, config.Db, config.Host, config.Port,
	)
	var err error
	Pool, err = pgxpool.Connect(Ctx, connStr)
	if err != nil {
		panic("Unable to connect to database: " + err.Error())
	}
	logger.Info("Connected to database")

}

func CloseConn() {
	cancel()
	Pool.Close()
}
