package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/pkg/database/user"
	"github.com/caarlos0/env"
	_ "github.com/go-sql-driver/mysql"
)

//DbConn stores the connexion to the database
var (
	DbConn *sql.DB
)

const (
	attemptsDBConnexion = 3
	waitForConnexion    = 3
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
		return err
	}

	for index := 1; index <= attemptsDBConnexion; index++ {
		err = db.Ping()
		if err != nil {
			if index < attemptsDBConnexion {
				log.Printf("db connection failed, %d retry : %v", index, err)
				time.Sleep(waitForConnexion * time.Second)
			}
			continue
		} else {
			DbConn = db
			user.NewUserRepositoryImpl(DbConn)
		}

		break
	}

	return nil
}
