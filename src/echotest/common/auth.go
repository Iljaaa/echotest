package common

import "fmt"
import "net/http"

// echo 
import "github.com/labstack/echo/v4"

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
        // c.Error("bbbb");

        cookie, err := c.Cookie("username")

	    fmt.Printf("error: %+v\n", err)
	    if err != nil {
    		return echo.NewHTTPError(http.StatusUnauthorized, err)
    	}

	    fmt.Printf("cookie: %s %v", cookie.Name, cookie.Value)
	    
        
        return echo.NewHTTPError(http.StatusUnauthorized, "Please provide valid credentials")
        
        // c.Response().Header().Set("Custom-Header", "blah!!!")
        // fmt.Println("auth middle ware")
        // return next(c)
    }
}
