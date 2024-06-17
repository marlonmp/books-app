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

func (us userService) FilterMany(ctx context.Context, uf *repos.UserFilters) ([]payloads.UserListPayload, error) {
	users, err := us.users.FilterMany(ctx, uf)

	if err != nil {
		return nil, err
	}

	usersPayload := payloads.UserListPLFromModels(users)

	return usersPayload, nil
}
