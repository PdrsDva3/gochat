package chat

import (
	"context"
	"github.com/jmoiron/sqlx"
	"gochat/internal/models"
	"gochat/internal/repository"
	"gochat/pkg/cerr"
)

type RepoChat struct {
	db *sqlx.DB
}

func InitChatRepository(db *sqlx.DB) repository.ChatRepo {
	return RepoChat{db: db}
}

func (r RepoChat) Create(ctx context.Context, chat models.ChatCreate) (int, error) {
	var id int
	transaction, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return 0, cerr.Err(cerr.Transaction, err).Error()
	}
	row := transaction.QueryRowContext(ctx, `INSERT INTO chat (name, description, id_creator) VALUES ($1, $2, $3) returning id;`,
		chat.Name, chat.Description, chat.IDCreator)
	err = row.Scan(&id)
	if err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {

			return 0, cerr.Err(cerr.Rollback, rbErr).Error()
		}
		return 0, cerr.Err(cerr.Scan, err).Error()
	}
	row = transaction.QueryRowContext(ctx, `INSERT INTO chat_user (id_chat, id_user)  VALUES ($1, $2) returning id_chat;`,
		id, chat.IDCreator)
	err = row.Scan(&id)
	if err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {

			return 0, cerr.Err(cerr.Rollback, rbErr).Error()
		}
		return 0, cerr.Err(cerr.Scan, err).Error()
	}
	for _, idUser := range chat.IDUsers {
		row1 := transaction.QueryRowContext(ctx, `INSERT INTO chat_user (id_chat, id_user) VALUES ($1, $2) returning id_chat`, id, idUser)
		err = row1.Scan(&id)
		if err != nil {
			if rbErr := transaction.Rollback(); rbErr != nil {
				return 0, cerr.Err(cerr.Rollback, rbErr).Error()
			}
			return 0, cerr.Err(cerr.Scan, err).Error()
		}
	}
	if err = transaction.Commit(); err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return 0, cerr.Err(cerr.Rollback, rbErr).Error()
		}
		return 0, cerr.Err(cerr.Commit, err).Error()
	}
	return id, nil
}

func (r RepoChat) List(ctx context.Context, id int) (*models.ChatList, error) {
	var list models.ChatList
	row, err := r.db.QueryContext(ctx, `SELECT id_chat FROM chat_user WHERE id_user=$1`, id)
	if err != nil {
		return nil, cerr.Err(cerr.Rows, err).Error()
	}
	for row.Next() {
		var chat models.ChatBase
		err = row.Scan(&chat.ID)
		if err != nil {
			return nil, cerr.Err(cerr.Scan, err).Error()
		}
		row1 := r.db.QueryRowContext(ctx, `SELECT name FROM chat WHERE id=$1`, chat.ID)
		err = row1.Scan(&chat.Name)
		if err != nil {
			return nil, cerr.Err(cerr.Scan, err).Error()
		}
		list.Chats = append(list.Chats, chat)
	}
	list.IDUser = id
	return &list, nil
}

func (r RepoChat) Get(ctx context.Context, id int) (*models.Chat, error) {
	var chat models.Chat
	row := r.db.QueryRowContext(ctx, `SELECT id, name, description, id_creator FROM chat where id=$1`, id)
	err := row.Scan(&chat.ID, &chat.Name, &chat.Description, &chat.IDCreator)
	if err != nil {
		return nil, cerr.Err(cerr.Scan, err).Error()
	}
	rows, err := r.db.QueryContext(ctx, `SELECT id_user FROM chat_user WHERE id_chat=$1`, chat.ID)
	if err != nil {
		return nil, cerr.Err(cerr.Rows, err).Error()
	}
	for rows.Next() {
		var idUser int
		err = rows.Scan(&idUser)
		if err != nil {
			return nil, cerr.Err(cerr.Scan, err).Error()
		}
		chat.IDUsers = append(chat.IDUsers, idUser)
	}
	return &chat, nil
}

func (r RepoChat) Delete(ctx context.Context, id int) error {
	transaction, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return cerr.Err(cerr.Transaction, err).Error()
	}
	result, err := transaction.ExecContext(ctx, `DELETE FROM chat where id_creator=$1`, id)

	if err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return cerr.Err(cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.Rows, err).Error()
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
	result, err = transaction.ExecContext(ctx, "DELETE FROM chat_user WHERE id_chat=$1", id)
	if err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return cerr.Err(cerr.Rollback, rbErr).Error()
		}
	}
	if result != nil {
		_, err = result.RowsAffected()
		if err != nil {
			if rbErr := transaction.Rollback(); rbErr != nil {
				return cerr.Err(cerr.Rollback, rbErr).Error()
			}
			return cerr.Err(cerr.Rows, err).Error()
		}
	}

	if err = transaction.Commit(); err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return cerr.Err(cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.Commit, err).Error()
	}

	return nil
}
