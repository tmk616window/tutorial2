package db

import (
	"database/sql"
	"fmt"
	"sidecar/config"
	"time"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func InitDB(dbUrl string, isLocal bool) (*sql.DB, error) {
	db, err := sql.Open("pgx", dbUrl)
	if err != nil {
		return nil, err
	}

	// db health check
	if err := db.Ping(); err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)
	db.SetConnMaxLifetime(300 * time.Second)

	boil.DebugMode = isLocal

	// DB をグローバルで呼べるように設定
	boil.SetDB(db) // nolint

	return db, nil
}

func URI(db config.Database) string {
	return fmt.Sprintf("user=%s password=%s database=%s host=%s port=%s", db.DBUSER, db.PASSWORD, db.DBNAME, db.DBHOST, db.DBPORT)
}
