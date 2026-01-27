package postgre_sql

import "context"

func (c *client) Exec(ctx context.Context, q string, args ...any) error {
	_, err := c.pool.Exec(ctx, q, args...)
	return err
}

func (c *client) Query(ctx context.Context, q string, args ...any) (Rows, error) {
	return c.pool.Query(ctx, q, args...)
}

func (c *client) QueryRow(ctx context.Context, q string, args ...any) Row {
	return c.pool.QueryRow(ctx, q, args...)
}

func (c *client) Begin(ctx context.Context) (Tx, error) {
	tx, err := c.pool.Begin(ctx)
	if err != nil {
		return nil, err
	}
	return &transaction{tx: tx}, nil
}

func (c *client) Ping(ctx context.Context) error {
	return c.pool.Ping(ctx)
}

func (c *client) Close() {
	c.pool.Close()
}

func (t *transaction) Raw() any {
	return t.tx // pgx.Tx
}
