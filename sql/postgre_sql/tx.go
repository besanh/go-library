package postgre_sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type transaction struct {
	tx pgx.Tx
}

func (t *transaction) Exec(ctx context.Context, q string, args ...any) error {
	_, err := t.tx.Exec(ctx, q, args...)
	return err
}

func (t *transaction) Query(ctx context.Context, q string, args ...any) (Rows, error) {
	return t.tx.Query(ctx, q, args...)
}

func (t *transaction) QueryRow(ctx context.Context, q string, args ...any) Row {
	return t.tx.QueryRow(ctx, q, args...)
}

func (t *transaction) Commit(ctx context.Context) error {
	return t.tx.Commit(ctx)
}

func (t *transaction) Rollback(ctx context.Context) error {
	return t.tx.Rollback(ctx)
}
