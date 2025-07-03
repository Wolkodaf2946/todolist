package repository

import (
	"fmt"

	"github.com/Wolkodaf2946/todolist"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user todolist.User) (int, error) {

	var id int

	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) values ($1, $2, $3) RETURNING id") // напишем запрос
	// при создании пользователя используем INSERT указывая имя username и пароль в таблицу
	// числа с $ - это плэйсхолдеры, в которые буду подставлены значения, которые мы передадим в качестве аргументов
	// к функции для выполнения запроса к бд.
	// в конце запроса RETURNING id, который будет возвращать айди новой записи после операции INSERT

	row := r.db.QueryRow(query, user.Name, user.Username, user.Password) // выполняем sql-запрос
	// метод принимает запрос и аргументы, которые будут вставлены в плэйсхолдеры из запроса.
	// метод возвращает объект row (хранит в себе информацию о возвращаемой строки из базы)
	// В нашем случае запрос возвразает id

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
