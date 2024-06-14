package repo

const (
	bookFilterMany = `
	select
		"id", "title", "description", "author_id", "book_path", "cover_path", "status", "created_at", "updated_at"
	from "books"
	`
)
