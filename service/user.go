package service

import (
	"context"

	"github.com/marlonmp/books-app/payload"
	"github.com/marlonmp/books-app/repo"
)

type userService struct {
	users repo.UserRepo
	books repo.BookRepo
}

func (us userService) FilterMany(ctx context.Context, uf *repo.UserFilters) ([]payload.UserListPayload, error) {
	users, err := us.users.FilterMany(ctx, uf)

	if err != nil {
		return nil, err
	}

	usersPayload := payload.UserListPLFromModels(users)

	return usersPayload, nil
}
