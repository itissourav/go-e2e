DB_URL=postgres://postgres:109798@localhost:5432/postgres?sslmode=disable

migrate-up:
	migrate -path migrations -database "$(DB_URL)" up

migrate-down:
	migrate -path migrations -database "$(DB_URL)" down 1

migrate-down-all:
	migrate -path migrations -database "$(DB_URL)" down -all

migrate-force:
	migrate -path migrations -database "$(DB_URL)" force $(version)

create-migration:
	migrate create -ext sql -dir migrations -seq $(name)