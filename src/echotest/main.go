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
        return controllers.Template(c)
	})

	e.GET("/profile", func(c echo.Context) error {
        con := controllers.ProfileController{}
        return con.Profile(c)
	}, common.AuthMiddleware)


	e.GET("/login", func(c echo.Context) error {
        return controllers.Login(c)
	})

	e.POST("/login", func(c echo.Context) error {
		return controllers.LoginPost(c)
	})

    e.GET("/logout", func(c echo.Context) error {
        return controllers.LoginOut(c)
	})

    // test then controller return error
	e.GET("/error", func(c echo.Context) error {
        return controllers.Error()
	})

    // thest then controller return code and error
	e.GET("/error404", func(c echo.Context) error {
        status, err := controllers.Error404(c)
        // if err != nil {e.Logger.Fatal(err)}
        return echo.NewHTTPError(status, err)
	})

	e.Logger.Fatal(e.Start(":8080"))
}