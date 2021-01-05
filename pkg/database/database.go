package database

import (
	"database/sql"
	"fmt"
	"github.com/caarlos0/env"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

//DbConn stores the connexion to the database
var (
	DbConn *sql.DB
)

// Config for DB connection
type Config struct {
	DbHost     string `env:"DB_HOST"`
	DbName     string `env:"MYSQL_DATABASE"`
	DbUser     string `env:"MYSQL_USER"`
	DbPassword string `env:"MYSQL_PASSWORD"`
	DbConn     *sql.DB
}

// Connect connection to database
func Connect() error {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		return fmt.Errorf("%+v", err)
	}
	dsn := cfg.DbUser + ":" + cfg.DbPassword + "@" + cfg.DbHost + "/" + cfg.
		DbName + "?parseTime=true&charset=utf8"

	db, err := sql.Open("mysql", dsn)

	if err != nil {
		log.Print(err)
		return err
	}

	err = db.Ping()

	if err == nil {
		DbConn = db
	}

	return nil
}
