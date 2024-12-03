package friend

import (
	"context"
	"github.com/jmoiron/sqlx"
	"gochat/internal/models"
	"gochat/internal/repository"
	"gochat/pkg/cerr"
)

type RepoFriend struct {
	db *sqlx.DB
}

func InitFriendRepository(db *sqlx.DB) repository.FriendRepo {
	return RepoFriend{db: db}
}

func (r RepoFriend) Get(ctx context.Context, id int) (*models.FriendList, error) {
	var friend models.FriendList
	row, err := r.db.QueryContext(ctx, `SELECT id_friend FROM friends WHERE id_user=$1`, id)
	if err != nil {
		return nil, cerr.Err(cerr.Rows, err).Error()
	}
	for row.Next() {
		var idFriend int
		err = row.Scan(&idFriend)
		if err != nil {
			return nil, cerr.Err(cerr.Scan, err).Error()
		}
		friend.IDFriends = append(friend.IDFriends, idFriend)
	}
	friend.IDUser = id
	return &friend, nil
}

func (r RepoFriend) AddFriend(ctx context.Context, id1 int, id2 int) error {
	transaction, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return cerr.Err(cerr.Transaction, err).Error()
	}
	_ = transaction.QueryRowContext(ctx, `INSERT INTO friends (id_user, id_friend) VALUES ($1, $2), ($2, $1)`, id1, id2)

	if err = transaction.Commit(); err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return cerr.Err(cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.Commit, err).Error()
	}
	return nil
}

func (r RepoFriend) Delete(ctx context.Context, id1 int, id2 int) error {
	transaction, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return cerr.Err(cerr.Transaction, err).Error()
	}
	result, err := transaction.ExecContext(ctx, `DELETE FROM friends WHERE (id_user=$1 AND id_friend=$2) or (id_user=$2 AND id_friend=$1)`, id1, id2)
	if err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return cerr.Err(cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.Transaction, err).Error()
	}
	count, err := result.RowsAffected()
	if err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return cerr.Err(cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.Rows, err).Error()
	}
	if count != 2 {
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
