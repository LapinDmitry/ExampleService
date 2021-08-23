package handlers

import (
	"context"
	gen "crud-grpc-server/third_party/grpcGenerate"
)

type Handlers struct {
	gen.UnimplementedServiceExampleServiceServer
}

func (h *Handlers) mustEmbedUnimplementedServiceExampleServiceServer() {
	panic("implement me")
}

func (h *Handlers) CreateUser(context.Context, *gen.CreateUserRequest) (*gen.User, error) {
	println("CreateUser")
	return &gen.User{}, nil
}
func (h *Handlers) UpdateUser(context.Context, *gen.UpdateUserRequest) (*gen.User, error) {
	println("UpdateUser")
	return &gen.User{}, nil
}
func (h *Handlers) DeleteUser(context.Context, *gen.DeleteUserRequest) (*gen.DeleteUserResponse, error) {
	println("DeleteUser")
	return &gen.DeleteUserResponse{}, nil
}
func (h *Handlers) ListUser(context.Context, *gen.ListUserRequest) (*gen.ListUserResponse, error) {
	println("ListUser")
	return &gen.ListUserResponse{}, nil
}
func (h *Handlers) GetUser(context.Context, *gen.GetUserRequest) (*gen.User, error) {
	println("GetUser")
	return &gen.User{}, nil
}
func (h *Handlers) CreateItem(context.Context, *gen.CreateItemRequest) (*gen.Item, error) {
	println("CreateItem")
	return &gen.Item{}, nil
}
func (h *Handlers) UpdateItem(context.Context, *gen.UpdateItemRequest) (*gen.Item, error) {
	println("UpdateItem")
	return &gen.Item{}, nil
}
