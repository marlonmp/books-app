package services

import (
	"context"

	"github.com/marlonmp/books-app/payloads"
	"github.com/marlonmp/books-app/repos"
)

type userService struct {
	users repos.UserRepo
	books repos.BookRepo
}

func (us userService) FilterMany(ctx context.Context, uf *repos.UserFilters) ([]payloads.UserList, error) {
	users, err := us.users.FilterMany(ctx, uf)

	if err != nil {
		return nil, err
	}

	usersPayload := payloads.UserListFromModels(users)

	return usersPayload, nil
}
