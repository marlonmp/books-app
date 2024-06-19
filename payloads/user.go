package payloads

import (
	"time"

	"github.com/google/uuid"
	"github.com/marlonmp/books-app/models"
	"github.com/marlonmp/books-app/valobjs"
)

type UserList struct {
	ID        uuid.UUID         `json:"id"`
	Username  string            `json:"username"`
	Nickname  string            `json:"nickname"`
	Bio       string            `json:"bio"`
	Status    models.UserStatus `json:"status,omitempty"`
	CreatedAt time.Time         `json:"created_at"`
	UpdatedAt time.Time         `json:"updated_at,omitempty"`
}

func UserListFromModel(u models.User) UserList {
	return UserList{
		ID:        u.ID,
		Username:  u.Username,
		Nickname:  u.Nickname,
		Bio:       u.Bio,
		Status:    u.Status,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func UserListFromModels(users []models.User) []UserList {
	payloads := make([]UserList, len(users))

	for i, u := range users {
		payloads[i] = UserListFromModel(u)
	}

	return payloads
}

type UserCreate struct {
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Bio      string `json:"bio"`
	Password string `json:"password"`
}

func (uc UserCreate) ToModel() (models.User, error) {
	password, err := valobjs.NewPassword(uc.Password)

	if err != nil {
		return models.User{}, err
	}

	user := models.NewUser(uc.Username, uc.Nickname, uc.Email, uc.Bio, password)

	return user, nil
}

type UserCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserProfile struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	Nickname  string    `json:"nickname"`
	Bio       string    `json:"bio"`
	CreatedAt time.Time `json:"created_at"`

	Books []BookList `json:"books"`
}

func UserProfileFromModel(u models.User) UserProfile {
	books := BookListFromModels(u.Books)

	return UserProfile{
		ID:        u.ID,
		Username:  u.Username,
		Nickname:  u.Nickname,
		Bio:       u.Bio,
		CreatedAt: u.CreatedAt,
		Books:     books,
	}
}
