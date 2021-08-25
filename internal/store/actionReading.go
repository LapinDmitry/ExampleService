package store

import (
	req "crud-grpc-server/internal/store/sqlRequests"
	gen "crud-grpc-server/third_party/grpcGenerated"
	"fmt"
	"strconv"
)

// GetUsersList - читать постранично набор пользователей с наборами предметов
//
func (s *Store) GetUsersList(page, limit int) ([]*gen.User, error) {
	rows, err := req.GetAllUsersRows(s.db, page, limit)
	if err != nil {
		return nil, fmt.Errorf("error of users list loading! Err(%v)", err)
	}
	return usersRowsToUsers(rows), nil
}

// GetUser - читать пользователя с набором предметов
//
func (s *Store) GetUser(id string) (*gen.User, error) {

	idNum, err := strconv.Atoi(id)
	if err != nil {
		return nil, fmt.Errorf("error! id:\"%s\" is not number", id)
	}

	rows, err := req.GetUserRows(s.db, idNum)
	if err != nil {
		return nil, fmt.Errorf("error of user loading! Err(%v)", err)
	}
	if len(rows) == 0 {
		return nil, fmt.Errorf("error! user was not found")
	}

	return usersRowsToUsers(rows)[0], nil
}

// Преобразует массив записей результата запроса к БД в массив пользователей
func usersRowsToUsers(rows []*req.UsersItemsRows) []*gen.User {
	userMap := make(map[int]*req.User)
	itemsMap := make(map[int][]*req.Item)
	for _, row := range rows {
		userMap[*row.User.Id] = &row.User

		if row.Item.Id != nil {
			itemsMap[*row.User.Id] = append(itemsMap[*row.User.Id], &row.Item)
		} else {
			itemsMap[*row.User.Id] = []*req.Item{}
		}
	}

	res := make([]*gen.User, 0, len(userMap))
	for id := range userMap {
		user := userMap[id]
		items := itemsMap[id]

		genUser := userReqToGen(user)
		genUser.Items = make([]*gen.Item, 0, len(items))

		for _, item := range items {
			genUser.Items = append(genUser.Items, itemReqToGen(item))
		}

		res = append(res, genUser)
	}

	return res
}
