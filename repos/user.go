package repos

import (
	"context"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"github.com/jackc/pgx"
	"github.com/marlonmp/books-app/models"
)

type UserFilters struct {
	UserID uuid.UUID
	Status models.UserStatus

	Search        string
	OrderBy       string
	Limit, Offset int
}

type UserRepo interface {
	FilterMany(ctx context.Context, uf *UserFilters) ([]models.User, error)

	CreateOne(ctx context.Context, u models.User) (models.User, error)

	GetCredentialsByUsername(ctx context.Context, username string, status models.UserStatus) (models.User, error)

	GetCredentialsByEmail(ctx context.Context, email string, status models.UserStatus) (models.User, error)

	GetByUsername(ctx context.Context, username string, status models.UserStatus) (models.User, error)

	GetByID(ctx context.Context, id uuid.UUID, status models.UserStatus) (models.User, error)

	UpdateByID(ctx context.Context, id uuid.UUID, status models.UserStatus, u models.User) (models.User, error)

	DeleteByID(ctx context.Context, id uuid.UUID, status models.UserStatus) (models.User, error)
}

type psqlUserRepo struct {
	conn *pgx.Conn
}

func PSQLUserRepo(conn *pgx.Conn) UserRepo {
	return psqlUserRepo{conn}
}

func (pur psqlUserRepo) buildFilters(uf *UserFilters) (string, []any) {
	var filters strings.Builder

	values := make([]any, 0)
	counter := 1

	if uf == nil {
		uf = new(UserFilters)
	}

	if uf.UserID != uuid.Nil {
		filters.WriteString(` and "id" = $`)
		filters.WriteString(strconv.Itoa(counter))

		counter++
		values = append(values, uf.UserID)
	}

	if uf.Status != models.UserStatusUnknown {
		filters.WriteString(` and "status" = $`)
		filters.WriteString(strconv.Itoa(counter))

		counter++
		values = append(values, uf.Status)
	}

	if counter > 1 {
		// remove the first ` and`
		q := filters.String()[4:]
		filters.Reset()

		filters.WriteString(q)
	}

	if uf.OrderBy != "" {
		order := "asc"
		orderField := uf.OrderBy

		if uf.OrderBy[0] == '+' {
			orderField = orderField[1:]
		}

		if uf.OrderBy[0] == '-' {
			order = "desc"
			orderField = orderField[1:]
		}
		filters.WriteString(` order by "$`)
		filters.WriteString(strconv.Itoa(counter))
		filters.WriteString(`" `)
		filters.WriteString(order)

		values = append(values, orderField)
	}

	if uf.Limit > 0 {
		filters.WriteString(` limit $`)
		filters.WriteString(strconv.Itoa(counter))

		counter++
		values = append(values, uf.Limit)
	}

	if uf.Offset > 0 {
		filters.WriteString(` offset $`)
		filters.WriteString(strconv.Itoa(counter))

		counter++
		values = append(values, uf.Limit)
	}

	return filters.String(), values
}

func (pur psqlUserRepo) buildFilterQuery(queryStr string, uf *UserFilters) (string, []any) {
	var query strings.Builder

	query.WriteString(queryStr)

	filters, values := pur.buildFilters(uf)

	if len(filters) > 0 {
		query.WriteString(filters)
	}

	query.WriteRune(';')

	return query.String(), values
}

func (pur psqlUserRepo) FilterMany(ctx context.Context, uf *UserFilters) ([]models.User, error) {
	query, values := pur.buildFilterQuery(userFilterMany, uf)

	rows, err := pur.conn.QueryEx(ctx, query, nil, values...)

	if err != nil {
		return nil, err
	}

	users := make([]models.User, 0)

	for rows.Next() {
		u := models.User{}

		err = rows.Scan(
			&u.ID,
			&u.Username,
			&u.Nickname,
			&u.Bio,
			&u.Status,
			&u.CreatedAt,
			&u.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		users = append(users, u)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (pur psqlUserRepo) CreateOne(ctx context.Context, u models.User) (models.User, error) {
	err := pur.
		conn.
		QueryRowEx(ctx, userCreateOne, nil, u.Username, u.Nickname, u.Email, u.Bio, u.Password, u.Status).
		Scan(&u.ID, &u.Email, &u.CreatedAt, &u.UpdatedAt)

	if err != nil {
		return models.User{}, err
	}

	return u, nil
}

func (pur psqlUserRepo) GetCredentialsByUsername(ctx context.Context, username string, status models.UserStatus) (u models.User, err error) {
	err = pur.
		conn.
		QueryRowEx(ctx, userGetCredentialsByUsername, nil, u.Username, status).
		Scan(&u.ID, &u.Username, &u.Password)

	if AsNotFoundError(&err) {
		return
	}

	return
}

func (pur psqlUserRepo) GetCredentialsByEmail(ctx context.Context, email string, status models.UserStatus) (u models.User, err error) {
	err = pur.
		conn.
		QueryRowEx(ctx, userGetCredentialsByEmail, nil, u.Email, status).
		Scan(&u.ID, &u.Email, &u.Password)

	if AsNotFoundError(&err) {
		return
	}

	return
}

func (pur psqlUserRepo) GetByUsername(ctx context.Context, username string, status models.UserStatus) (u models.User, err error) {
	err = pur.
		conn.
		QueryRowEx(ctx, userGetByUsername, nil, u.Username, status).
		Scan(&u.ID, &u.Username, &u.Nickname, &u.Bio, &u.Status, &u.CreatedAt, &u.UpdatedAt)

	if AsNotFoundError(&err) {
		return
	}

	return
}

func (pur psqlUserRepo) GetByID(ctx context.Context, id uuid.UUID, status models.UserStatus) (u models.User, err error) {
	err = pur.
		conn.
		QueryRowEx(ctx, userGetByID, nil, u.ID, status).
		Scan(&u.ID, &u.Username, &u.Nickname, &u.Email, &u.Bio, &u.Status, &u.CreatedAt, &u.UpdatedAt)

	if AsNotFoundError(&err) {
		return
	}

	return
}

func (pur psqlUserRepo) UpdateByID(ctx context.Context, id uuid.UUID, status models.UserStatus, u models.User) (models.User, error) {
	err := pur.
		conn.
		QueryRowEx(ctx, userUpdateByID, nil, u.Username, u.Nickname, u.Email, u.Bio, u.Status, u.ID, status).
		Scan(&u.ID, &u.Username, &u.Nickname, &u.Email, &u.Bio, &u.Status, &u.CreatedAt, &u.UpdatedAt)

	if AsNotFoundError(&err) {
		return models.User{}, err
	}

	return u, nil
}

func (pur psqlUserRepo) DeleteByID(ctx context.Context, id uuid.UUID, status models.UserStatus) (u models.User, err error) {
	err = pur.
		conn.
		QueryRowEx(ctx, userDeleteByID, nil, u.ID, status).
		Scan(&u.ID, &u.Username, &u.Nickname, &u.Email, &u.Bio, &u.Status, &u.CreatedAt, &u.UpdatedAt)

	if AsNotFoundError(&err) {
		return
	}

	return
}
