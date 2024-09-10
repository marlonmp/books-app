package payloads

import (
	"time"

	"github.com/google/uuid"
	"github.com/marlonmp/books-app/models"
)

type BookList struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	BookPath    string    `json:"book_path"`
	CoverPath   string    `json:"cover_path"`
	CreatedAt   time.Time `json:"created_at"`
}

func BookListFromModel(b models.Book) BookList {
	return BookList{
		ID:          b.ID,
		Title:       b.Title,
		Description: b.Description,
		BookPath:    b.BookPath,
		CoverPath:   b.CoverPath,
		CreatedAt:   b.CreatedAt,
	}
}

func BookListFromModels(books []models.Book) []BookList {
	payloads := make([]BookList, len(books))

	for i, book := range books {
		payloads[i] = BookListFromModel(book)
	}

	return payloads
}
