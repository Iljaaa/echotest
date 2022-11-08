package main

import "fmt";

// echo https://echo.labstack.com/guide/

// echo 
import "github.com/labstack/echo/v4"

// postgres
import "github.com/jackc/pgx/v5"
import "context"
import "os"

// import "github.com/Iljaaa/echotest/src/common";
import "github.com/Iljaaa/echotest/src/config";
import "github.com/Iljaaa/echotest/src/controllers";
import "github.com/Iljaaa/echotest/src/common/templates";

func main() {
    fmt.Println("startng test script.....")

    // get config
    config := config.GetConfig()
    fmt.Printf("%+v\n", config.DB)

    // connect to db
    connectToDb()
    
    // starting server
    startEcho()
}

func connectToDb() {

    config := config.GetConfig()
    fmt.Printf("%+v\n", config.DB)

    dbConnectString := fmt.Sprintf ("postgres://postgres:example@db:5432/sportscools")
    fmt.Printf("connect to db: %s", dbConnectString)

    conn, err := pgx.Connect(context.Background(), "postgres://postgres:example@db:5432/sportscools")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())
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