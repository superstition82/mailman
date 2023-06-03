migrate-up:
	migrate -path db/migration -database "sqlite3://db.sqlite" -verbose up

migrate-down:
	migrate -path db/migration -database "sqlite3://db.sqlite" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: migrate-db migrate-down sqlc test