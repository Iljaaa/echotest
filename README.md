# echotest

### framework server 
echo: [https://echo.labstack.com/](https://echo.labstack.com/)  


### templates

### postgres   
pgx: [https://github.com/jackc/pgx/v5](https://github.com/jackc/pgx/v5)  

``
go get github.com/jackc/pgx/v5
``

[https://github.com/jackc/pgx/wiki/Getting-started-with-pgx](https://github.com/jackc/pgx/wiki/Getting-started-with-pgx)

### migrations 
[https://github.com/pressly/goose](https://github.com/pressly/goose)

``go install github.com/pressly/goose/v3/cmd/goose@latest``

*not sure. maybe depends on this lib*   
``go install github.com/golang-migrate/migrate/v4``

install to bin directory   
``go build -tags='no_mysql no_sqlite3' -o goose`` 

#### status
```./goose postgres "host=db user=postgres password=example dbname=sportscools sslmode=disable" status```

#### create migration
```./goose -dir ./../migrations create init sql```

#### up
```./goose -dir ./../migrations postgres "host=db user=postgres password=example dbname=sportscools sslmode=disable" up```

### validation
[https://github.com/go-playground/validator](https://github.com/go-playground/validator)

- Configs
- html|template
  - layouts
- error handler
- db
- migration
- models
- **React**
- Validation
- Auth
  - Logout
  - **Safe cookie code**
