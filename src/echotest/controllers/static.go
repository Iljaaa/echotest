package controllers

import "fmt"
import "net/http"
import "errors"

import "github.com/labstack/echo/v4"

// hello world
func HelloWorld() (int, string, error) {
	fmt.Printf("hellow world controler")
    return http.StatusOK, "Hello, World", nil
}

// template render
func Template(c echo.Context) (int, error) {
	c.Render(http.StatusOK, "home.html", map[string]interface{}{
		"name": "HOME",
		"msg": "Hello, Boatswain!",
	})
    return http.StatusOK, nil
}

// exmaple with error
func Error() (int, string, error) {
	return 400, "", errors.New("test error")
}