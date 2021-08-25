package store

import (
	req "crud-grpc-server/internal/store/sqlRequests"
	gen "crud-grpc-server/third_party/grpcGenerated"
	"google.golang.org/protobuf/types/known/timestamppb"
	"strconv"
	"time"
)

// Конвертирует user из типа для запросов к БД в тип, сгенерированный grpc
func userReqToGen(user *req.User) *gen.User {
	tCreate, _ := time.Parse(CommonLayout, *user.CreateAt)
	tUpdate, _ := time.Parse(CommonLayout, *user.UpdateAt)

	return &gen.User{
		Id:        strconv.Itoa(*user.Id),
		Name:      *user.Name,
		Age:       *user.Age,
		CreatedAt: timestamppb.New(tCreate),
		UpdatedAt: timestamppb.New(tUpdate),
	}
}

// Конвертирует item из типа для запросов к БД в тип, сгенерированный grpc
func itemReqToGen(item *req.Item) *gen.Item {
	tCreate, _ := time.Parse(CommonLayout, *item.CreateAt)
	tUpdate, _ := time.Parse(CommonLayout, *item.UpdateAt)

	return &gen.Item{
		Id:        strconv.Itoa(*item.Id),
		Name:      *item.Name,
		UserId:    strconv.Itoa(*item.UserId),
		CreatedAt: timestamppb.New(tCreate),
		UpdatedAt: timestamppb.New(tUpdate),
	}
}
