postgres:
	docker run --name go-lang-postgres -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root --rm -p 10101:5432 -d postgres:15

createdb:
	docker exec -it go-lang-postgres createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it go-lang-postgres dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://root:root@localhost:10101/simple_bank?sslmode=disable" --verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:root@localhost:10101/simple_bank?sslmode=disable" --verbose down

sqlc:
	docker run --rm -v "D:/Projects/Go Projects/Simple Bank:/src" -w /src sqlc/sqlc generate

.PHONY: postgres createdb dropdb migrateup migratedown sqlc