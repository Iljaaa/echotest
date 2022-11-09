# echotest

echo 
https://echo.labstack.com/

templates

postgres
https://github.com/jackc/pgx/v5
https://github.com/jackc/pgx/wiki/Getting-started-with-pgx
go get github.com/jackc/pgx/v5


migrations 
https://github.com/pressly/goose

go install github.com/pressly/goose/v3/cmd/goose@latest

не понял, возможно зависит от этой библиотеки
go install github.com/golang-migrate/migrate/v4


go build -tags='no_mysql no_sqlite3' -o goose 

status
./goose postgres "host=db user=postgres password=example dbname=sportscools sslmode=disable" status

cretae migration
./goose -dir ./../migrations create init sql

./goose -dir ./../migrations postgres "host=db user=postgres password=example dbname=sportscools sslmode=disable" up

validation
https://github.com/go-playground/validator