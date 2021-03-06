package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

const (
	usersTable        string = "users"
	poemsTable        string = "poems"
	authorTable       string = "authors"
	authorsListsTable string = "authors_list"
)

type Config struct {
	Host     string
	Port     string
	UserName string
	Password string
	DBname   string
	SSLmode  string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.UserName, cfg.DBname, cfg.Password, cfg.SSLmode))
	if err != nil {
		return nil, err
	}
	err = db.Ping()

	if err != nil {
		return nil, err
	}

	return db, nil
}
