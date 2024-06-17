package payload

import (
	"time"

	"github.com/google/uuid"
	"github.com/marlonmp/books-app/model"
)

type UserListPayload struct {
	ID        uuid.UUID        `json:"id"`
	Username  string           `json:"username"`
	Nickname  string           `json:"nickname"`
	Bio       string           `json:"bio"`
	Status    model.UserStatus `json:"status,omitempty"`
	CreatedAt time.Time        `json:"created_at"`
	UpdatedAt time.Time        `json:"updated_at,omitempty"`
}

func UserListPLFromModel(u model.User) UserListPayload {
	return UserListPayload{
		ID:        u.ID,
		Username:  u.Username,
		Nickname:  u.Nickname,
		Bio:       u.Bio,
		Status:    u.Status,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func UserListPLFromModels(users []model.User) []UserListPayload {
	payloads := make([]UserListPayload, len(users))

	for i, u := range users {
		payloads[i] = UserListPLFromModel(u)
	}

	return payloads
}
