package main

import "fmt";

// echo https://echo.labstack.com/guide/

import "github.com/labstack/echo/v4"

// import "github.com/Iljaaa/echotest/src/common";
import "github.com/Iljaaa/echotest/src/config";
import "github.com/Iljaaa/echotest/src/controllers";
import "github.com/Iljaaa/echotest/src/common/templates";

func main() {
    fmt.Println("startng test script.....")

    // get config
    config := config.GetConfig()
    fmt.Printf("%+v\n", config.DB)
    
    // starting server
    startEcho()
}

func startEcho () {

    e := echo.New()

    // template renderer
    t := templates.GetTemplates()
	e.Renderer = &templates.TemplateRenderer{
		Templates: t,
	}

	e.GET("/", func(c echo.Context) error {
        status, error := controllers.Template(c)
        if error != nil {
            e.Logger.Fatal(error)
        }
		return c.NoContent(status)
	})

	e.GET("/error", func(c echo.Context) error {
        status, body, error := controllers.Error()
        if error != nil { e.Logger.Fatal(error) }
		return c.String(status, body)
	})

	e.Logger.Fatal(e.Start(":8080"))
}