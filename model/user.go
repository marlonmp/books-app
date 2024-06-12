package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/marlonmp/books-app/valobj"
)

type UserStatus uint

const (
	UserStatusUnknown UserStatus = iota
	UserStatusUnverified
	UserStatusActive
	UserStatusBanned
	UserStatusDeleted
)

func (us UserStatus) String() string {
	switch us {
	case UserStatusUnverified:
		return "Unverified"
	case UserStatusActive:
		return "Active"
	case UserStatusBanned:
		return "Banned"
	case UserStatusDeleted:
		return "Deleted"
	default:
		return "Unknown"
	}
}

type User struct {
	ID uuid.UUID

	Username,
	Nickname,
	Bio string

	CreatedBook []Book

	Password valobj.Password

	Status UserStatus

	CreatedAt,
	UpdatedAt time.Time
}

func NewUser(username, nickname, bio, pwd string) (User, error) {
	pwd_, err := valobj.NewPassword(pwd)

	if err != nil {
		return User{}, err
	}

	user := User{
		ID:        uuid.New(),
		Username:  username,
		Nickname:  nickname,
		Bio:       bio,
		Password:  pwd_,
		Status:    UserStatusUnverified,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return user, nil
}
