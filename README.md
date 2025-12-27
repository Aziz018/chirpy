# chirpy

## Migrations

### `goose up` migration
```bash
cd ./sql/schema/ # location of schemas
goose postgres "postgres://user:pass@localhost:5432/dbname?sslmode=disable" up
# or
goose -dir ./sql/schema/ postgres "postgres://user:pass@localhost:5432/dbname?sslmode=disable" up

```

### `goose down` migration
```bash
cd ./sql/schema/ # location of schemas
goose postgres "postgres://user:pass@localhost:5432/dbname?sslmode=disable" down
# or
goose -dir ./sql/schema/ postgres "postgres://user:pass@localhost:5432/dbname?sslmode=disable" down

```
## SQLC

```yml
# sqlc.yml file:
version: "2"
sql:
  - schema: "sql/schema"
    queries: "sql/queries"
    engine: "postgresql"
    gen:
      go:
        out: "internal/database"
```

```bash
sqlc generate
```
