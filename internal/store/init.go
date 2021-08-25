// Package store
/*
	Реализует общение с базой данных, принимает и возвращает объекты сгенерированного grpc кода

	Инициализируется методом New()
*/
package store

import (
	"fmt"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

type Store struct {
	db *sqlx.DB
}

// New - создаёт объект для общения с БД
// Инициализируется только этим методом
func New(login, password, endpoint, dbName string) (*Store, error) {

	host := fmt.Sprintf("postgres://%s:%s@%s/%s", login, password, endpoint, dbName)
	host = "postgres://postgres:1401@localhost:5432/db_test"

	db, err := sqlx.Connect("pgx", host)
	if err != nil {
		return nil, fmt.Errorf("error connecting to the database! Err(%v)", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("database ping error! Err(%v)", err)
	}

	return &Store{db: db}, nil
}
