include local.env
export

DB_URL := $(MIGRATE_DB_URL)

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