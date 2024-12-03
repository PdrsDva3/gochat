package user

import (
	"context"
	"github.com/jmoiron/sqlx"
	"gochat/internal/models"
	"gochat/internal/repository"
	"gochat/pkg/cerr"
)

type RepoUser struct {
	db *sqlx.DB
}

func InitUserRepository(db *sqlx.DB) repository.UserRepo {
	return RepoUser{db: db}
}

func (r RepoUser) Create(ctx context.Context, user models.UserCreate) (int, error) {
	var id int
	transaction, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return 0, cerr.Err(cerr.Transaction, err).Error()
	}
	row := transaction.QueryRowContext(ctx, `INSERT INTO users (nickname, email, phone, name, surname, pwd, photo, description) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) returning id;`,
		user.Nickname, user.Email, user.Phone, user.Name, user.Surname, user.Password, user.Photo, user.Description)

	err = row.Scan(&id)
	if err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return 0, cerr.Err(cerr.Rollback, rbErr).Error()
		}
		return 0, cerr.Err(cerr.Scan, err).Error()
	}
	if err = transaction.Commit(); err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return 0, cerr.Err(cerr.Rollback, rbErr).Error()
		}
		return 0, cerr.Err(cerr.Commit, err).Error()
	}
	return id, nil
}

func (r RepoUser) Get(ctx context.Context, id int) (*models.User, error) {
	var user models.User
	row := r.db.QueryRowContext(ctx, `SELECT id, nickname, email, phone, name, surname, photo, description from users WHERE id = $1;`, id)
	err := row.Scan(&user.ID, &user.Nickname, &user.Email, &user.Phone, &user.Name, user.Surname, &user.Photo, &user.Description)
	if err != nil {
		return nil, cerr.Err(cerr.Scan, err).Error()
	}
	return &user, nil
}

func (r RepoUser) Delete(ctx context.Context, id int) error {
	transaction, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return cerr.Err(cerr.Transaction, err).Error()
	}
	result, err := transaction.ExecContext(ctx, `DELETE FROM users WHERE id=$1;`, id)
	if err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return cerr.Err(cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.ExecContext, err).Error()
	}
	count, err := result.RowsAffected()
	if err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return cerr.Err(cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.Rows, err).Error()
	}
	if count != 1 {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return cerr.Err(cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.NoOneRow, err).Error()
	}
	if err = transaction.Commit(); err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return cerr.Err(cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.Commit, err).Error()
	}
	return nil
}

func (r RepoUser) GetPWDbyEmail(ctx context.Context, user string) (int, string, error) {
	var pwd string
	var id int
	rows := r.db.QueryRowContext(ctx, `SELECT id, pwd from users WHERE email = $1;`, user)
	err := rows.Scan(&id, &pwd)
	if err != nil {
		return 0, "", cerr.Err(cerr.Scan, err).Error()
	}
	return id, pwd, nil
}
