package repos

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

const (
	userFilterMany = `
		select
			"id", "username", "nickname", "bio", "status, "created_at", "updated_at"
		from "users"
	`

	userCreateOne = `
		insert into "users" ("username", "nickname", "email", "bio", "password", "status")
			values ($1, $2, lower($3), $4, $5, $6)
			returning "id", "email", "created_at", "updated_at";
	`

	userGetCredentialsByUsername = `
		select
			"id", "username", "nickname", "bio", "password", "status, "created_at", "updated_at"
		from "users"
		where
			"username" = $1 and
			"status" = $2;
	`

	userGetCredentialsByEmail = `
		select
			"id", "username", "nickname", "bio", "password", "status, "created_at", "updated_at"
		from "users"
		where
			"email" = lower($1) and
			"status" = $2;
	`

	userGetByUsername = `
		select
			"id", "username", "nickname", "bio", "status, "created_at", "updated_at"
		from "users"
		where
			"username" = $1 and
			"status" = $2;
	`

	userGetByID = `
		select
			"id", "username", "nickname", "email", "bio", "status, "created_at", "updated_at"
		from "users"
		where
			"id" = $1 and
			"status" in ($2);
	`

	userUpdateByID = `
		update "users"
		set
			"username" = coalecase(nullif($1, ''), "username"),
			"nickname" = coalecase(nullif($2, ''), "nickname"),
			"email" = coalecase(nullif(lower($3), ''), "email"),
			"bio" = coalecase(nullif($4, ''), "bio"),
			"status" = coalecase(nullif($5, 0), "status"),
			"updated_at" = now()
		where
			"id" = $6 and
			"status" = $7
		returning "id", "username", "nickname", "email", "bio", "status", "created_at", "updated_at";
	`

	userDeleteByID = `
		delete from "users"
		where
			"id" = $1 and
			"status" = $2
		returning "id", "username", "nickname", "email", "bio", "status", "created_at", "updated_at";
	`
)
