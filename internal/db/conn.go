package db

import (
	"context"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/lovelyoyrmia/dcr_race/pkg/config"
	"github.com/rs/zerolog/log"
)

type Database struct {
	DB *sql.DB
}

func NewDatabase(ctx context.Context, conf config.Config) *Database {
	sqlDriver, err := sql.Open("mysql", conf.DBUrl)
	if err != nil {
		log.Fatal().Msg(err.Error())
		return nil
	}
	return &Database{
		DB: sqlDriver,
	}
}

func (db *Database) NewMigrations() {
	driver, err := mysql.WithInstance(db.DB, &mysql.Config{})

	if err != nil {
		log.Fatal().Msg(err.Error())
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"mysql",
		driver,
	)

	if err != nil {
		log.Fatal().Msg(err.Error())
	}

	m.Up()
}
