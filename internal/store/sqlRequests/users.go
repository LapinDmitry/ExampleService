package sqlRequests

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

// GetAllUsersRows - Загрузить записи со всеми пользователями и их предметами
//
func GetAllUsersRows(db *sqlx.DB, page, limit int) ([]*UsersItemsRows, error) {
	var ans []*UsersItemsRows
	err := db.Select(&ans, getUsers, page*limit, limit)
	if err != nil {
		return nil, fmt.Errorf("request execution error! Err(%v)", err)
	}

	for i := range ans {
		ans[i].Item.UserId = ans[i].User.Id
	}
	return ans, err
}

const getUsers = `
SELECT 
    U.id AS user_id,
    U.name AS user_name,
    U.age AS user_age,
    U.create_at AS user_create_at,
    U.update_at AS user_update_at,
       
	"Items".id AS item_id,
	"Items".name AS item_name,
    "Items".create_at AS item_create_at,
    "Items".update_at AS item_update_at
FROM (
	SELECT *
	FROM "Users"
	ORDER BY "Users".id
	OFFSET $1
	LIMIT $2
	) U
LEFT JOIN "Items" ON U.id = "Items".user_id
ORDER BY U.id
`

// GetUserRows - записи пользователя с предметами
//
func GetUserRows(db *sqlx.DB, id int) ([]*UsersItemsRows, error) {
	var ans []*UsersItemsRows
	err := db.Select(&ans, getUser, id)
	if err != nil {
		return nil, fmt.Errorf("request execution error! Err(%v)", err)
	}

	for i := range ans {
		ans[i].Item.UserId = ans[i].User.Id
	}
	return ans, err
}

const getUser = `
SELECT 
    "Users".id AS user_id,
    "Users".name AS user_name,
    "Users".age AS user_age,
    "Users".create_at AS user_create_at,
    "Users".update_at AS user_update_at,
       
	"Items".id AS item_id,
	"Items".name AS item_name,
    "Items".create_at AS item_create_at,
    "Items".update_at AS item_update_at
FROM "Users"
LEFT JOIN "Items" ON "Users".id = "Items".user_id
WHERE "Users".id = $1
`

// CreateUser - создать пользователя
//
func CreateUser(db *sqlx.DB, user *User) (*User, error) {

	// Создадим юзера
	rows, err := db.NamedQuery(SQLRecCreateUser, user)
	if err != nil {
		return nil, fmt.Errorf("request execution error! Err(%v)", err)
	}

	// Узнаем его id
	rows.Next()
	err = rows.StructScan(user)
	if err != nil {
		return nil, fmt.Errorf("error of response scanning! User(%v) didn't get its id. Err(%v)", *user, err)
	}

	return user, nil
}

const SQLRecCreateUser = `
INSERT INTO "Users"
(name, age, user_type, create_at, update_at)
VALUES (:user_name, :user_age, :user_type, :user_create_at, :user_update_at)
RETURNING id AS user_id
`

// UpdateUser - Обновить пользователя
//
func UpdateUser(db *sqlx.DB, user *User) (*User, error) {
	/*
		Тут должна была быть транзакция, но я её не осилил
		Метод .LastInsertId() отказывался работать
		Возвращалась ошибка "LastInsertId is not supported by this driver"
		Поэтому тут без транзакции
	*/

	// Создадим юзера
	rows, err := db.NamedQuery(SQLRecUpdateUser, user)
	if err != nil {
		return nil, fmt.Errorf("request execution error! Err(%v)", err)
	}

	// Возьмем метку create
	rows.Next()
	err = rows.StructScan(user)
	if err != nil {
		return nil, fmt.Errorf("error of response scanning! User(%v) \"created_at\" is missing. Err(%v)", *user, err)
	}

	return user, nil
}

const SQLRecUpdateUser = `
UPDATE "Users"
SET name=:user_name, age=:user_age, user_type=:user_type, update_at=:user_update_at
WHERE id = :user_id
RETURNING create_at AS user_create_at
`

// DeleteUser Удалить пользователя
//
func DeleteUser(db *sqlx.DB, user *User) error {

	// Создадим юзера
	res, err := db.NamedExec(SQLRecDeleteUser, user)
	if err != nil {
		return fmt.Errorf("request execution error! Err(%v)", err)
	}

	// Возьмем метку create
	qt, _ := res.RowsAffected()
	if qt != 1 {
		return fmt.Errorf("уккщк! The user was not deleted. Number of modified records:%v", qt)
	}

	return nil
}

const SQLRecDeleteUser = `
DELETE  FROM "Users"
WHERE id = :user_id
`
