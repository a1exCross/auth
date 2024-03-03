package pg

import (
	"github.com/a1exCross/auth/internal/client/db"
	"github.com/jackc/pgx/v4/pgxpool"

	"context"
	"fmt"
)

type pgClient struct {
	masterDBC db.DB
}

// New - создает новый клиент БД
func New(ctx context.Context, dsn string) (db.Client, error) {
	dbc, err := pgxpool.Connect(ctx, dsn)
	if err != nil {
		return nil, fmt.Errorf("falied to connect to database: %v", err)
	}
	return &pgClient{
		masterDBC: NewDB(dbc),
	}, nil
}

func (c *pgClient) DB() db.DB {
	return c.masterDBC
}

func (c *pgClient) Close() error {
	if c.masterDBC != nil {
		c.masterDBC.Close()
	}

	return nil
}
