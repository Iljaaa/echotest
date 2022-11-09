package main

import "fmt"
// import "errors"
// import "net/http"

// echo https://echo.labstack.com/guide/

// echo 
import "github.com/labstack/echo/v4"

// validator
import "github.com/go-playground/validator"

// postgres
// import "github.com/jackc/pgx/v5"
// import "context"
// import "os"

import "github.com/Iljaaa/echotest/src/common";
import "github.com/Iljaaa/echotest/src/common/db"; 
import "github.com/Iljaaa/echotest/src/config";
import "github.com/Iljaaa/echotest/src/controllers";

func main() {
    fmt.Println("startng test script.....")

    // get config
    config := config.GetConfig()
    fmt.Printf("%+v\n", config.DB)

    // connect to db
    db.InitConnection()
    
    // starting server
    startEcho()
}

func startEcho () {

    e := echo.New()
    e.Debug = true

    // error handler
    e.HTTPErrorHandler = common.CustomHTTPErrorHandler

    // template renderer
    t := common.GetTemplates()
	e.Renderer = &common.TemplateRenderer{
		Templates: t,
	}

    // validator
    e.Validator = &controllers.CustomValidator{Validator: validator.New()}
    

    // auth middle ware
    // e.Use(authMiddleware)


	e.GET("/", func(c echo.Context) error {
        status, error := controllers.Template(c)
        if error != nil {
            e.Logger.Fatal(error)
        }
		return c.NoContent(status)
	})

	e.GET("/profile", func(c echo.Context) error {
        status, body, error := controllers.Error()
        if error != nil { e.Logger.Fatal(error) }
		return c.String(status, body)
	}, common.AuthMiddleware)


	e.GET("/login", func(c echo.Context) error {
        status, error := controllers.Login(c)
        if error != nil {e.Logger.Fatal(error)}
		return c.NoContent(status)
	})

	e.POST("/login", func(c echo.Context) error {
        status, error := controllers.Login(c)
        if error != nil {e.Logger.Fatal(error)}
		return c.NoContent(status)
	})

	e.GET("/error", func(c echo.Context) error {
        status, body, error := controllers.Error()
        if error != nil { e.Logger.Fatal(error) }
		return c.String(status, body)
	})

	e.Logger.Fatal(e.Start(":8080"))
}