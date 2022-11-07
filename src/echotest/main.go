package main

import "fmt";

// echo https://echo.labstack.com/guide/

// import "net/http"
import "github.com/labstack/echo/v4"

import "github.com/Iljaaa/echotest/src/common";
import "github.com/Iljaaa/echotest/src/config";
import "github.com/Iljaaa/echotest/src/controllers";
// import "echo_test/common";

func main() {
    fmt.Println("startng test script.....")
    common.Test()
    // fmt.Sprintf("asdasdas %T", config.Configuration)
    config := config.GetConfig()
    // a := "DB"
    fmt.Printf("%+v\n", config.DB)
    // fmt.Printf("%+v\n", config.db)

    // start echo
    startEcho()
}

func startEcho () {
    e := echo.New()
	e.GET("/", func(c echo.Context) error {
        status, body, error := controllers.HelloWorld()
        if error != nil {
            e.Logger.Fatal(error)
        }
		return c.String(status, body)
	})
	e.GET("/error", func(c echo.Context) error {
        status, body, error := controllers.Error()
        if error != nil { e.Logger.Fatal(error) }
		return c.String(status, body)
	})
	e.Logger.Fatal(e.Start(":8080"))
}