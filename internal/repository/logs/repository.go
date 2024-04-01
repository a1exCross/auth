package logsrepository

import (
	"context"
	"fmt"

	"github.com/a1exCross/auth/internal/model"
	"github.com/a1exCross/auth/internal/repository"

	"github.com/a1exCross/common/pkg/client/db"

	sq "github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
)

const (
	tableName = "logs"

	idColumn        = "id"
	actionColumn    = "action"
	contentColumn   = "content"
	createdAtColumn = "created_at"
)

// NewRepository - возвращает методы для работы с репозиторием логов
func NewRepository(db db.Client) repository.LogsRepository {
	return repo{
		db: db,
	}
}

type repo struct {
	db db.Client
}

func (r repo) Create(ctx context.Context, params model.Log) (int64, error) {
	insertBuilder := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(actionColumn, contentColumn).
		Values(params.Action, params.Content).
		Suffix(fmt.Sprintf("RETURNING %s", idColumn))

	query, args, err := insertBuilder.ToSql()
	if err != nil {
		return 0, errors.Errorf("error at parse sql builder: %v", err)
	}

	var id int64

	q := db.Query{
		Name:     "logs_repository.Create",
		QueryRaw: query,
	}

	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&id)
	if err != nil {
		return 0, errors.Errorf("error at query to database: %v", err)
	}

	return id, nil
}

func (r repo) Get(ctx context.Context, id int64) (model.Log, error) {
	selectBuilder := sq.Select(actionColumn, contentColumn, createdAtColumn).
		PlaceholderFormat(sq.Dollar).
		From(tableName).
		Where(sq.Eq{idColumn: id})

	query, args, err := selectBuilder.ToSql()
	if err != nil {
		return model.Log{}, errors.Errorf("error at parse sql builder: %v", err)
	}

	q := db.Query{
		Name:     "logs_repository.Get",
		QueryRaw: query,
	}

	var log model.Log

	err = r.db.DB().ScanOneContext(ctx, &log, q, args...)
	if err != nil {
		return model.Log{}, errors.Errorf("error at query to database: %v", err)
	}

	return log, nil
}
