// Package sqlRequests
/*
	Пакет для запросов к БД
	В запросах участвуют структуры (Item, User, UsersItemsRows), через поля которых передаётся информация
*/
package sqlRequests

type Item struct {
	Id       *int    `db:"item_id"`
	Name     *string `db:"item_name"`
	UserId   *int    `db:"item_user_id"`
	CreateAt *string `db:"item_create_at"`
	UpdateAt *string `db:"item_update_at"`
}

type User struct {
	Id       *int    `db:"user_id"`
	Name     *string `db:"user_name"`
	Age      *int32  `db:"user_age"`
	Type     *int    `db:"user_type"`
	CreateAt *string `db:"user_create_at"`
	UpdateAt *string `db:"user_update_at"`
}

type UsersItemsRows struct {
	User
	Item
}
