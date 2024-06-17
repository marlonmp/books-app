package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/marlonmp/books-app/valobj"
)

type UserStatus uint8

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
	Email,
	Bio string

	Books []Book

	Password valobj.Password

	Status UserStatus

	CreatedAt,
	UpdatedAt time.Time
}

func NewUser(username, nickname, email, bio string, pwd valobj.Password) User {
	user := User{
		Username: username,
		Nickname: nickname,
		Email:    email,
		Bio:      bio,
		Password: pwd,
		Status:   UserStatusUnverified,
	}

	return user
}
