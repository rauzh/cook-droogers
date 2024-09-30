package postgres

import (
	"context"
	"cookdroogers/internal/repo"
	"cookdroogers/models"
	"database/sql"
	"fmt"
	trmsqlx "github.com/avito-tech/go-transaction-manager/drivers/sqlx/v2"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type UserPgRepo struct {
	db         *sqlx.DB
	txResolver *trmsqlx.CtxGetter
}

func NewUserPgRepo(db *sql.DB) repo.UserRepo {
	dbx := sqlx.NewDb(db, "pgx")

	return &UserPgRepo{db: dbx, txResolver: trmsqlx.DefaultCtxGetter}
}

func (usr *UserPgRepo) Create(ctx context.Context, user *models.User) error {

	q := "INSERT INTO users (name, email, type, password) VALUES ($1, $2, $3, $4) RETURNING user_id"

	var userID uint64
	err := usr.txResolver.DefaultTrOrDB(ctx, usr.db).QueryRowxContext(ctx, q,
		user.Name, user.Email, user.Type, user.Password).Scan(&userID)

	if err != nil {
		return errors.Wrap(PgDbErr, err.Error())
	}

	user.UserID = userID
	return nil
}

func (usr *UserPgRepo) GetByEmail(ctx context.Context, email string) (*models.User, error) {

	q := "SELECT user_id, name, email, password, type FROM users WHERE email=$1"

	user := models.User{}

	err := usr.txResolver.DefaultTrOrDB(ctx, usr.db).QueryRowxContext(ctx, q, email).Scan(
		&user.UserID, &user.Name, &user.Email, &user.Password, &user.Type)

	if err != nil {
		return nil, errors.Wrap(PgDbErr, err.Error())
	}

	return &user, nil
}

func (usr *UserPgRepo) Get(ctx context.Context, id uint64) (*models.User, error) {

	q := "SELECT user_id, name, email, password, type FROM users WHERE user_id=$1"

	user := models.User{}

	err := usr.txResolver.DefaultTrOrDB(ctx, usr.db).QueryRowxContext(ctx, q, id).Scan(
		&user.UserID, &user.Name, &user.Email, &user.Password, &user.Type)

	if err != nil {
		return nil, errors.Wrap(PgDbErr, err.Error())
	}

	return &user, nil
}

func (usr *UserPgRepo) Update(ctx context.Context, user *models.User) error {

	q := "UPDATE users SET name=$1, email=$2, type=$3, password=$4 WHERE user_id=$5"

	_, err := usr.txResolver.DefaultTrOrDB(ctx, usr.db).ExecContext(ctx, q,
		user.Name, user.Email, user.Type, user.Password, user.UserID)

	if err != nil {
		return errors.Wrap(PgDbErr, err.Error())
	}

	return nil
}

func (usr *UserPgRepo) UpdateType(ctx context.Context, userID uint64, typ models.UserType) error {

	q := "UPDATE users SET type=$1 WHERE user_id=$2"

	_, err := usr.txResolver.DefaultTrOrDB(ctx, usr.db).ExecContext(ctx, q, typ, userID)

	if err != nil {
		return errors.Wrap(PgDbErr, err.Error())
	}

	return nil
}

func (usr *UserPgRepo) SetRole(ctx context.Context, role models.UserType) error {

	var roleStr string

	switch role {
	case models.NonMemberUser:
		roleStr = "NonMemberUser"
	case models.ManagerUser:
		roleStr = "ManagerUser"
	case models.ArtistUser:
		roleStr = "ArtistUser"
	case models.AdminUser:
		roleStr = "AdminUser"
	}

	q := fmt.Sprintf("SET ROLE %s;", roleStr)

	_, err := usr.txResolver.DefaultTrOrDB(ctx, usr.db).ExecContext(ctx, q)

	if err != nil {
		return errors.Wrap(PgDbErr, err.Error())
	}

	return nil
}

func (usr *UserPgRepo) GetForAdmin(ctx context.Context) ([]models.User, error) {
	q := "SELECT user_id, name, email, password, type FROM users"

	users := make([]models.User, 0)

	rows, err := usr.txResolver.DefaultTrOrDB(ctx, usr.db).QueryxContext(ctx, q)

	if errors.Is(err, sql.ErrNoRows) {
		return users, nil
	}
	if err != nil {
		return nil, errors.Wrap(PgDbErr, err.Error())
	}
	defer rows.Close()

	for rows.Next() {

		user := models.User{}

		err := rows.Scan(&user.UserID, &user.Name, &user.Email, &user.Password, &user.Type)
		if err != nil {
			return nil, errors.Wrap(PgDbErr, err.Error())
		}

		users = append(users, user)
	}

	return users, nil
}
