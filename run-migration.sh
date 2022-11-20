
#!/bin/sh
export DATABASE_URL="postgres://postgres:hwhwhwlol@localhost:5432/user_db?sslmode=disable" \
export MIGRATION_PATH="/Users/abim/Documents/go/src/learn/go-grpc-example/services/user/migrations/pgsql" \
# go run services/migration/main.go migration:status
# go run services/migration/main.go migration:down
go run services/migration/main.go migration:up
# go run services/migration/main.go migration:create create_table_auth --table=auth