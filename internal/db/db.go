package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"fibertest/internal/config"

	"database/sql"
	_ "github.com/lib/pq"
	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/dialects/postgresql"
)

var DB *reform.DB
var SqlDb *sql.DB

func ConnectDb() {
	config := config.GetConfig()

	connStr := fmt.Sprintf("user=%v password=%v dbname=%v sslmode=%v host=%v port=%v",
		config.Postgres.User,
		config.Postgres.Password,
		config.Postgres.Name,
		config.Postgres.SSLMode,
		config.Postgres.Host,
		config.Postgres.Port,
	)

	SqlDb, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatalf("Ошибка при создании пула соединений: %v", err)
	}

	SqlDb.SetMaxOpenConns(config.Postgres.MaxOpenConns)
	SqlDb.SetMaxIdleConns(config.Postgres.MaxIdleConns)
	SqlDb.SetConnMaxLifetime(time.Second * config.Postgres.ConnMaxLifetime)

	err = SqlDb.Ping()
	if err != nil {
		log.Fatalf("Ошибка при проверке соединения: %v", err)
	}

	logger := log.New(os.Stderr, "SQL: ", log.Flags())

	DB = reform.NewDB(SqlDb, postgresql.Dialect, reform.NewPrintfLogger(logger.Printf))
}
