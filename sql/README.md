# SQL Package

The `sql` package handles relational database connectivity. Currently, it supports PostgreSQL explicitly by wrapping the high-performance `github.com/jackc/pgx/v5` driver.

## Architecture Flow (PostgreSQL)

1. **Configuration**: Initializes using a robust `Config` encompassing the Data Source Name (`DSN`), and critical pooling bounds (`MaxConns`, `MinConns`, `MaxConnLifetime`).
2. **Pool Bootstrapping**: Uses `pgxpool` to establish a persistent connection logic over standard `database/sql`, gaining access to PostgreSQL-specific native features.
3. **Health Validation**: Overrides the standard health check period (defaults to 30s) and forces an immediate `Ping` against the database to guarantee availability on boot.

## Usage

```go
import "github.com/besanh/go-library/sql/postgre_sql"
```

### Connecting to PostgreSQL

```go
cfg := postgre_sql.Config{
    DSN:             "postgres://user:pass@localhost:5432/mydb",
    MaxConns:        50,
    MinConns:        5,
    MaxConnLifetime: time.Hour,
}

db, err := postgre_sql.NewPostgreSql(context.Background(), cfg)
if err != nil {
    log.Fatal("Could not connect to database:", err)
}

// Access the underlying *pgxpool.Pool 
// pool := db.(*postgre_sql.client).pool
```
