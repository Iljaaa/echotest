package controllers

import "fmt"
import "net/http"
import "errors"

import "github.com/labstack/echo/v4"

import "github.com/Iljaaa/echotest/src/models/users"; 

// hello world
func HelloWorld() (int, string, error) {
	fmt.Printf("hellow world controler\n")
    return http.StatusOK, "Hello, World", nil
}

// template render
func Template(c echo.Context) error {
	
	fmt.Printf("Template\n")
	u, _ := models.FindById(1)
	if u != nil {
		fmt.Printf("u %T %+v %v\n", u.Id, u.Id, u.Id)
	}

	// u := users.FindById(1)
	// fmt.Printf("u %T %+v\n", u)

	c.Render(http.StatusOK, "home.html", map[string]interface{}{
		"name": "HOME2",
		"msg": "Hello, Boatswain!",
	})

    return c.NoContent(http.StatusOK)
}

// exmaple with error
func Error() error {
	return errors.New("test error")
}


// exmaple with error
func Error404(c echo.Context) (int, error) {
	return http.StatusNotFound, errors.New("test error2222")
}