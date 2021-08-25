package store

import (
	req "crud-grpc-server/internal/store/sqlRequests"
	gen "crud-grpc-server/third_party/grpcGenerated"
	"github.com/jmoiron/sqlx"
	"strconv"
	"time"
)

// Реализует набор взаимодействий с БД, которые выполняются в рамках одной транзакции
// Инициализируется через функцию openTransaction()
type transaction struct {
	//db *sqlx.Tx
	db *sqlx.DB
}

// openTransaction() открывает транзакцию и возвращает объект транзакции
func openTransaction(db *sqlx.DB) *transaction {
	//tr := db.MustBegin()
	tr := &transaction{db: db}
	return tr
}

// Commit - Коммитит все внесённые изменения
func (t *transaction) Commit() {
	//t.db.Commit()
}

// CreateUser - создать запись пользователя
func (t *transaction) CreateUser(createUser *gen.CreateUserRequest, timeCreate time.Time) (*req.User, error) {
	tm := timeCreate.UTC().Format(CommonLayout)
	userType := int(createUser.UserType)

	user := &req.User{
		Name:     &createUser.Name,
		Age:      &createUser.Age,
		Type:     &userType,
		CreateAt: &tm,
		UpdateAt: &tm,
	}

	var err error
	user, err = req.CreateUser(t.db, user)

	return user, err
}

// UpdateUser - обновить запись пользователя
func (t *transaction) UpdateUser(updateUser *gen.UpdateUserRequest, timeUpdate time.Time) (*req.User, error) {
	tm := timeUpdate.UTC().Format(CommonLayout)
	userType := int(updateUser.UserType)

	userId, _ := strconv.Atoi(updateUser.Id)

	user := &req.User{
		Id:       &userId,
		Name:     &updateUser.Name,
		Age:      &updateUser.Age,
		Type:     &userType,
		UpdateAt: &tm,
	}

	var err error
	user, err = req.UpdateUser(t.db, user)

	return user, err
}

// DeleteUser - Удалить запись пользователя и все связанные с ней записи айтемов
func (t *transaction) DeleteUser(deleteUser *gen.DeleteUserRequest) error {
	userId, _ := strconv.Atoi(deleteUser.Id)

	user := &req.User{
		Id: &userId,
	}

	var err error
	err = req.DeleteUser(t.db, user)

	return err
}

// CreateItems - создать несколько объектов
func (t *transaction) CreateItems(createItems []*gen.CreateItemRequest, timeCreate time.Time) ([]*req.Item, error) {
	tm := timeCreate.UTC().Format(CommonLayout)

	items := make([]*req.Item, len(createItems))
	for i, genItem := range createItems {
		userId, _ := strconv.Atoi(genItem.UserId)

		items[i] = &req.Item{
			Name:     &genItem.Name,
			CreateAt: &tm,
			UpdateAt: &tm,
			UserId:   &userId,
		}
	}

	var err error
	items, err = req.CreateItems(t.db, items)

	return items, err
}

// UpdateItems - Обновить несколько объектов
func (t *transaction) UpdateItems(updateItems []*gen.UpdateItemRequest, timeCreate time.Time) ([]*req.Item, error) {
	tm := timeCreate.UTC().Format(CommonLayout)

	items := make([]*req.Item, len(updateItems))
	for i, genItem := range updateItems {
		id, _ := strconv.Atoi(genItem.Id)

		items[i] = &req.Item{
			Id:       &id,
			Name:     &genItem.Name,
			UpdateAt: &tm,
		}
	}

	var err error
	items, err = req.UpdateItems(t.db, items)

	return items, err
}
