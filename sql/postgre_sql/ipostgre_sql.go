package postgre_sql

import (
	"context"
)

type IDB interface {
	Exec(ctx context.Context, query string, args ...any) error
	Query(ctx context.Context, query string, args ...any) (Rows, error)
	QueryRow(ctx context.Context, query string, args ...any) Row

	Begin(ctx context.Context) (Tx, error)
	Ping(ctx context.Context) error
	Close()
	Raw() any
}

type Tx interface {
	Exec(ctx context.Context, query string, args ...any) error
	Query(ctx context.Context, query string, args ...any) (Rows, error)
	QueryRow(ctx context.Context, query string, args ...any) Row

	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
	Raw() any
}

type Rows interface {
	Next() bool
	Scan(dest ...any) error
	Close()
	Err() error
}

type Row interface {
	Scan(dest ...any) error
}

func (c *client) Raw() any {
	return c.pool // *pgxpool.Pool
}
