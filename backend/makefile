create_migrations:
	migrate create -ext sql -dir db/migrations -seq $(name)

postgres_up:
	docker-compose up -d

postgres_down:
	docker-compose down 

db_up:
	docker exec -it fintech_postgres createdb --username=root --owner=root fintech_db

db_down:
	docker exec -it fintech_postgres dropdb --username=root fintech_db

migrate_up:
	migrate -path db/migrations -database "postgres://root:secret@localhost:5432/fintech_db?sslmode=disable" up

migrate_down:
	migrate -path db/migrations -database "postgres://root:secret@localhost:5432/fintech_db?sslmode=disable" down

sqlc:
	sqlc generate