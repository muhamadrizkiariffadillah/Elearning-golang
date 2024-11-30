postgres:
	docker run --name postgres-learning -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=123qweasdzxc -d postgres:latest

createdb:
	docker exec -it postgres-learning createdb --username=root --owner=root e-learning

dropdb:
	docker exec -it postgres-learning dropdb e-learning

create_migration:
	migrate create -ext sql -dir ./migration -seq init_schema

migrationup:
	migrate -path ./migration -database "postgresql://root:123qweasdzxc@localhost:5432/e-learning?sslmode=disable" -verbose up

migrationdown:
	migrate -path ./migration -database "postgresql://root:123qweasdzxc@localhost:5432/e-learning?sslmode=disable" -verbose down


.PHONY: postgres createdb dropdb create_migration migrationup migrationdown