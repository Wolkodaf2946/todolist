package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

const ( // названия таблиц из бд
	usersTable     = "users"
	todoListsTable = "todo_lists"
	userListsTable = "user_lists"
	todoItemsTable = "todo_items"
	listsItemTable = "lists_items"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
