@echo off
setlocal

:: Load environment variables from local.env
for /f "usebackq tokens=1,* delims==" %%A in ("local.env") do (
    if not "%%A"=="" if not "%%A:~0,1%%"=="#" (
        set "%%A=%%B"
    )
)

set DB_URL=%MIGRATE_DB_URL%

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