package store

import (
	req "crud-grpc-server/internal/store/sqlRequests"
	gen "crud-grpc-server/third_party/grpcGenerated"
	"fmt"
	"strconv"
	"time"
)

const CommonLayout = "2006-1-2T15:04:05Z"

// CreateUser - создать пользователя с набором предметов
//
func (s *Store) CreateUser(createUser *gen.CreateUserRequest) (*gen.User, error) {
	tm := time.Now()

	t := openTransaction(s.db)
	user, err := t.CreateUser(createUser, tm)
	if err != nil {
		return nil, fmt.Errorf("error creating user record! Err(%v)", err)
	}

	createItems := createUser.Items
	for _, createItem := range createItems {
		createItem.UserId = strconv.Itoa(*user.Id)
	}

	items, err := t.CreateItems(createItems, tm)
	if err != nil {
		return nil, fmt.Errorf("error creating items records! Err(%v)", err)
	}

	t.Commit()
	return reqUserReqItemsToGenUser(user, items), nil
}

// UpdateUser - обновить пользователя и набор предметов
//
func (s *Store) UpdateUser(updateUser *gen.UpdateUserRequest) (*gen.User, error) {
	tm := time.Now()

	t := openTransaction(s.db)
	user, err := t.UpdateUser(updateUser, tm)
	if err != nil {
		return nil, fmt.Errorf("error updating user record! Err(%v)", err)
	}

	updateItems := updateUser.Items
	items, err := t.UpdateItems(updateItems, tm)
	if err != nil {
		return nil, fmt.Errorf("error updating items records! Err(%v)", err)
	}

	t.Commit()
	return reqUserReqItemsToGenUser(user, items), nil
}

// преобразует *req.User + []*req.Item = *gen.User
func reqUserReqItemsToGenUser(user *req.User, items []*req.Item) *gen.User {
	genUser := userReqToGen(user)
	genUser.Items = make([]*gen.Item, len(items))
	for i, item := range items {
		genUser.Items[i] = itemReqToGen(item)
	}
	return genUser
}

// DeleteUser - удалить пользователя и связанные предметы
//
func (s *Store) DeleteUser(deleteUser *gen.DeleteUserRequest) error {

	t := openTransaction(s.db)
	err := t.DeleteUser(deleteUser)
	if err != nil {
		return fmt.Errorf("error deleting records! Err(%v)", err)
	}

	t.Commit()
	return nil
}

// CreateItem - создать предмет с привязкой к пользователю
//
func (s *Store) CreateItem(createItem *gen.CreateItemRequest) (*gen.Item, error) {
	tm := time.Now()

	createItems := []*gen.CreateItemRequest{createItem}

	t := openTransaction(s.db)
	items, err := t.CreateItems(createItems, tm)
	if err != nil {
		return nil, fmt.Errorf("error creating item records! Err(%v)", err)
	}

	t.Commit()
	return itemReqToGen(items[0]), nil
}

// UpdateItem - обновить предмет
//
func (s *Store) UpdateItem(updateItem *gen.UpdateItemRequest) (*gen.Item, error) {
	tm := time.Now()

	updateItems := []*gen.UpdateItemRequest{updateItem}

	t := openTransaction(s.db)
	items, err := t.UpdateItems(updateItems, tm)
	if err != nil {
		return nil, fmt.Errorf("error updating item records! Err(%v)", err)
	}

	t.Commit()
	return itemReqToGen(items[0]), nil
}
