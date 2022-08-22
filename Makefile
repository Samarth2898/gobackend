postgres:
	docker run --name postgres14 -p 5433:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14-alpine

createdb:
	docker exec -it postgres14 createdb --username=root --owner=root simple_bank

dropdb: 
	docker exec -it postgres14 dropdb simple_bank

migrateup: 
	migrate -path db/migration -database "postgresql://root:secret@localhost:5433/simple_bank?sslmode=disable" -verbose up

migrateup1: 
	migrate -path db/migration -database "postgresql://root:secret@localhost:5433/simple_bank?sslmode=disable" -verbose up 1

migratedown: 
	migrate -path db/migration -database "postgresql://root:secret@localhost:5433/simple_bank?sslmode=disable" -verbose down

migratedown1: 
	migrate -path db/migration -database "postgresql://root:secret@localhost:5433/simple_bank?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/Samarth2898/gobackend/db/sqlc Store

.PHONY:	postgres createdb dropdb migrateup migratedown sqlc server mock migrateup1 migratedown1