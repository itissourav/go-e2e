@echo off

set DB_URL=postgres://postgres:109798@localhost:5432/postgres?sslmode=disable

if "%1"=="up" (
    migrate -path migrations -database "%DB_URL%" up
    goto end
)

if "%1"=="down" (
    migrate -path migrations -database "%DB_URL%" down 1
    goto end
)

if "%1"=="down-all" (
    migrate -path migrations -database "%DB_URL%" down -all
    goto end
)

if "%1"=="create" (
    if "%2"=="" (
        echo Please provide migration name
        echo Example: migrate.bat create create_users_table
        goto end
    )
    migrate create -ext sql -dir migrations -seq %2
    goto end
)

echo Usage:
echo migrate.bat up
echo migrate.bat down
echo migrate.bat down-all
echo migrate.bat create migration_name

:end
pause