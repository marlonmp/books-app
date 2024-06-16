package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/marlonmp/books-app/valobj"
)

type BookStatus uint8

const (
	BookStatusUnknown BookStatus = iota
	BookStatusDraft
	BookStatusPublic
	BookStatusPrivate
	BookStatusDeleted
)

type Book struct {
	ID uuid.UUID

	Title,
	Description string

	AuthorID uuid.UUID
	Author   *User

	BookPath,
	CoverPath string

	BookFile,
	CoverFile *valobj.File

	Status BookStatus

	CreatedAt,
	UpdatedAt time.Time
}

func NewBook(title, description string, authorID uuid.UUID) Book {
	return Book{
		Title:       title,
		Description: description,
		AuthorID:    authorID,
		Status:      BookStatusDraft,
	}
}
