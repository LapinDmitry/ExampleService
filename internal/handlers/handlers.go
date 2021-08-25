// Package handlers
/*
	Содержит в себе объект обработчика Handlers вызова удалённых функций по grpc
*/
package handlers

import (
	"context"
	"fmt"
	storage "github.com/LapinDmitry/ExampleService/internal/store"
	"github.com/LapinDmitry/ExampleService/internal/utils"
	gen "github.com/LapinDmitry/ExampleService/third_party/grpcGenerated"
)

// Handlers Обработчик grpc запросов
// Cоздаётся и инициализируется через функцию Start
type Handlers struct {
	gen.UnimplementedServiceExampleServer
	store *storage.Store
}

// Start Инициализирует и запускает обработчик соединений
func Start() (*Handlers, error) {
	login := utils.GetEnv("EXAMPLE_SERVER_DB_LOGIN", "postgres").(string)
	pass := utils.GetEnv("EXAMPLE_SERVER_DB_PASSWORD", "").(string)
	endpoint := utils.GetEnv("EXAMPLE_SERVER_DB_ENDPOINT", "localhost:5432").(string)
	dbname := utils.GetEnv("EXAMPLE_SERVER_DB_NAME", "db_test").(string)

	store, err := storage.New(login, pass, endpoint, dbname)
	if err != nil {
		return nil, fmt.Errorf("storage connection error! Err(%v)", err)
	}

	fmt.Print("GRPS handler is started\n")
	return &Handlers{store: store}, nil
}

func (Handlers) MustEmbedUnimplementedServiceExampleServiceServer() {}

func (h *Handlers) CreateUser(ctx context.Context, req *gen.CreateUserRequest) (*gen.User, error) {

	fmt.Print("CreateUser\n")
	return h.store.CreateUser(req)
}
func (h *Handlers) UpdateUser(ctx context.Context, req *gen.UpdateUserRequest) (*gen.User, error) {

	fmt.Print("UpdateUser\n")
	return h.store.UpdateUser(req)
}
func (h *Handlers) DeleteUser(ctx context.Context, req *gen.DeleteUserRequest) (*gen.DeleteUserResponse, error) {

	fmt.Print("DeleteUser\n")
	return &gen.DeleteUserResponse{}, h.store.DeleteUser(req)
}
func (h *Handlers) ListUser(ctx context.Context, req *gen.ListUserRequest) (*gen.ListUserResponse, error) {

	fmt.Print("ListUser\n")
	list, err := h.store.GetUsersList(int(req.Page), int(req.Limit))
	return &gen.ListUserResponse{Users: list}, err
}
func (h *Handlers) GetUser(ctx context.Context, req *gen.GetUserRequest) (*gen.User, error) {

	fmt.Print("CGetUser\n")
	return h.store.GetUser(req.Id)
}
func (h *Handlers) CreateItem(ctx context.Context, req *gen.CreateItemRequest) (*gen.Item, error) {

	fmt.Print("CreateItem\n")
	return h.store.CreateItem(req)
}
func (h *Handlers) UpdateItem(ctx context.Context, req *gen.UpdateItemRequest) (*gen.Item, error) {

	fmt.Print("UpdateItem\n")
	return h.store.UpdateItem(req)
}
