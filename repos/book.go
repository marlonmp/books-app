package repos

import (
	"context"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"github.com/jackc/pgx"
	"github.com/marlonmp/books-app/models"
)

type BookFilters struct {
	BookID   uuid.UUID
	AuthorID uuid.UUID
	Status   models.BookStatus

	Search        string
	OrderBy       string
	Limit, Offset int
}

type BookRepo interface {
	// Return a list of books with the given filters, if find nothing, returns
	// an empty array
	FilterMany(ctx context.Context, bf *BookFilters) ([]models.Book, error)

	// Creates one book in the repo and return the created book
	CreateOne(ctx context.Context, b models.Book) (models.Book, error)

	// Returns one book with the given id, if find nothing, returns a
	// [NotFoundError]
	GetByID(ctx context.Context, id uuid.UUID) (models.Book, error)

	// Returns one book with the given id, returned the updated book, if find
	// nothing, returns a [NotFoundError]
	UpdateByID(ctx context.Context, id uuid.UUID, b models.Book) (models.Book, error)

	// Delete and returns one book with the given id, if find nothing, returns
	// a [NotFoundError]
	DeleteByID(ctx context.Context, id uuid.UUID) (models.Book, error)
}

type psqlBookRepo struct {
	conn *pgx.Conn
}

func PSQLBookRepo(conn *pgx.Conn) BookRepo {
	return psqlBookRepo{conn}
}

func (pbr psqlBookRepo) buildFilters(bf *BookFilters) (string, []any) {
	var filters strings.Builder

	values := make([]any, 0)
	counter := 1

	if bf == nil {
		bf = new(BookFilters)
	}

	if bf.BookID != uuid.Nil {
		filters.WriteString(` and "id" = $`)
		filters.WriteString(strconv.Itoa(counter))

		counter++
		values = append(values, bf.BookID)
	}

	if bf.AuthorID != uuid.Nil {
		filters.WriteString(` and "author_id" = $`)
		filters.WriteString(strconv.Itoa(counter))

		counter++
		values = append(values, bf.AuthorID)
	}

	if bf.Status != models.BookStatusUnknown {
		filters.WriteString(` and "status" = $`)
		filters.WriteString(strconv.Itoa(counter))

		counter++
		values = append(values, bf.Status)
	}

	if counter > 1 {
		// remove the first ` and`
		q := filters.String()[4:]
		filters.Reset()

		filters.WriteString(q)
	}

	if bf.OrderBy != "" {
		order := "asc"
		orderField := bf.OrderBy

		if bf.OrderBy[0] == '+' {
			orderField = orderField[1:]
		}

		if bf.OrderBy[0] == '-' {
			order = "desc"
			orderField = orderField[1:]
		}
		filters.WriteString(` order by "$`)
		filters.WriteString(strconv.Itoa(counter))
		filters.WriteString(`" `)
		filters.WriteString(order)

		values = append(values, orderField)
	}

	if bf.Limit > 0 {
		filters.WriteString(` limit $`)
		filters.WriteString(strconv.Itoa(counter))

		counter++
		values = append(values, bf.Limit)
	}

	if bf.Offset > 0 {
		filters.WriteString(` offset $`)
		filters.WriteString(strconv.Itoa(counter))

		counter++
		values = append(values, bf.Limit)
	}

	return filters.String(), values
}

func (pbr psqlBookRepo) FilterMany(ctx context.Context, bf *BookFilters) ([]models.Book, error) {
	var query strings.Builder

	query.WriteString(bookFilterMany)

	filters, values := pbr.buildFilters(bf)

	if len(filters) > 0 {
		query.WriteString(filters)
	}

	query.WriteRune(';')

	rows, err := pbr.conn.QueryEx(ctx, query.String(), nil, values...)

	if err != nil {
		return nil, err
	}

	books := make([]models.Book, 0)

	for rows.Next() {
		book := models.Book{}

		err = rows.Scan(
			&book.ID,
			&book.Title,
			&book.Description,
			&book.AuthorID,
			&book.BookPath,
			&book.CoverPath,
			&book.Status,
			&book.CreatedAt,
			&book.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		books = append(books, book)
	}

	return books, nil
}

func (pbr psqlBookRepo) CreateOne(ctx context.Context, b models.Book) (models.Book, error) {
	row := pbr.conn.QueryRowEx(ctx, bookCreateOne, nil, b.Title, b.Description, b.AuthorID, b.BookPath, b.CoverPath, b.Status)

	err := row.Scan(&b.ID, &b.CreatedAt, &b.UpdatedAt)

	if err != nil {
		return models.Book{}, nil
	}

	return b, nil
}

func (pbr psqlBookRepo) GetByID(ctx context.Context, id uuid.UUID) (b models.Book, err error) {
	row := pbr.conn.QueryRowEx(ctx, bookGetByID, nil, id)

	err = row.Scan(
		&b.ID,
		&b.Title,
		&b.Description,
		&b.AuthorID,
		&b.BookPath,
		&b.CoverPath,
		&b.Status,
		&b.CreatedAt,
		&b.UpdatedAt,
	)

	if AsNotFoundError(&err) {
		return
	}

	return
}

func (pbr psqlBookRepo) UpdateByID(ctx context.Context, id uuid.UUID, b models.Book) (models.Book, error) {
	row := pbr.conn.QueryRowEx(ctx, bookUpdateByID, nil, b.Title, b.Description, b.AuthorID, b.BookPath, b.CoverPath, b.Status, id)

	err := row.Scan(
		&b.ID,
		&b.Title,
		&b.Description,
		&b.AuthorID,
		&b.BookPath,
		&b.CoverPath,
		&b.Status,
		&b.CreatedAt,
		&b.UpdatedAt,
	)

	if AsNotFoundError(&err) {
		return b, err
	}

	return b, nil
}

func (pbr psqlBookRepo) DeleteByID(ctx context.Context, id uuid.UUID) (b models.Book, err error) {
	row := pbr.conn.QueryRowEx(ctx, bookDeleteByID, nil, id)

	err = row.Scan(
		&b.ID,
		&b.Title,
		&b.Description,
		&b.AuthorID,
		&b.BookPath,
		&b.CoverPath,
		&b.Status,
		&b.CreatedAt,
		&b.UpdatedAt,
	)

	if AsNotFoundError(&err) {
		return
	}

	return
}
