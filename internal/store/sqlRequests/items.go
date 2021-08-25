package sqlRequests

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

// CreateItems - Создать несколько Items
//
func CreateItems(db *sqlx.DB, items []*Item) ([]*Item, error) {
	// Если список атемов нулевой
	if items == nil || len(items) == 0 {
		return []*Item{}, nil
	}

	// Создадим айтемы
	rows, err := db.NamedQuery(SQLRecCreateItems, items)
	if err != nil {
		return nil, fmt.Errorf("request execution error! Err(%v)", err)
	}

	// Возьмем их id
	for _, item := range items {
		ok := rows.Next()
		err = rows.StructScan(item)
		if !ok || err != nil {
			return nil, fmt.Errorf("error of response scanning! Item(%v) didn't get its id. Err(%v)", *item, err)
		}
	}

	return items, nil
}

const SQLRecCreateItems = `
INSERT INTO "Items"
(name, user_id, create_at, update_at)
VALUES (:item_name, :item_user_id, :item_create_at, :item_update_at)
RETURNING id AS item_id
`

// UpdateItems - Обновить несколько Items
//
func UpdateItems(db *sqlx.DB, items []*Item) ([]*Item, error) {

	for _, item := range items {
		// Обновим item
		rows, err := db.NamedQuery(SQLRecUpdateItems, item)
		if err != nil {
			return nil, fmt.Errorf("request execution error! Err(%v)", err)
		}

		// Возьмем метку create и user_id
		ok := rows.Next()
		err = rows.StructScan(item)
		if !ok || err != nil {
			return nil, fmt.Errorf("error of response scanning! Item(%v) \"created_at\" is missing. Err(%v)", *item, err)
		}
	}

	return items, nil
}

const SQLRecUpdateItems = `
UPDATE "Items"
SET name=:item_name,  update_at=:item_update_at
WHERE id = :item_id
RETURNING create_at AS item_create_at, user_id AS item_user_id
`
