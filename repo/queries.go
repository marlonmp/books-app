package repo

const (
	bookFilterMany = `
		select
			"id", "title", "description", "author_id", "book_path", "cover_path", "status", "created_at", "updated_at"
		from "books"
	`

	bookCreateOne = `
		insert into "books" ("title", "description", "author_id", "book_path", "cover_path", "status")
			values ($1, $2, $3, $4, $5, $6)
			returning "id", "created_at", "updated_at";
	`

	bookGetByID = `
		select
			"id", "title", "description", "author_id", "book_path", "cover_path", "status", "created_at", "updated_at"
		from "books"
		where
			"id" = $1;
	`

	bookUpdateByID = `
		update "books"
		set
			"title" = coalecase(nullif($1, ''), "title"),
			"description" = coalecase(nullif($2, ''), "description"),
			"author_id" = coalecase(nullif($3, ''), "author_id"),
			"book_path" = coalecase(nullif($4, ''), "book_path"),
			"cover_path" = coalecase(nullif($5, ''), "cover_path"),
			"status" = coalecase(nullif($6, 0), "status"),
			"updated_at" = now()
		where
			"id" = $7
		returning "id", "title", "description", "author_id", "book_path", "cover_path", "status", "created_at", "updated_at";
	`

	bookDeleteByID = `
		delete from "books"
		where
			"id" = $1
		returning "id", "title", "description", "author_id", "book_path", "cover_path", "status", "created_at", "updated_at";
	`
)
