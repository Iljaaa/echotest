package controllers

import "fmt"
import "net/http"
import "errors"
import "io/ioutil"
import "regexp"

// echo
import "github.com/labstack/echo/v4"

import "github.com/Iljaaa/echotest/src/common"; 
import "github.com/Iljaaa/echotest/src/models/users"; 

type IndexPageData struct {
	JsFiles []string
}

//
// index page
//
func IndexPage(c echo.Context) error {

	// get js fils
	files, err := ioutil.ReadDir("../frontend/build/static/js")
    if err != nil { return echo.NewHTTPError(http.StatusInternalServerError, err) }

	jsFiles := []string{}
    for _, file := range files {

		res1, e := regexp.MatchString(`^main\.[a-z0-9]+\.js$`, file.Name())
		if e == nil && res1 { 
			jsFiles = append(jsFiles, "/record/static/js/" + file.Name())
		}
    }

	// get css files
	files, err = ioutil.ReadDir("../frontend/build/static/css")
    if err != nil { return echo.NewHTTPError(http.StatusInternalServerError, err) }

	cssFile := []string{}
	for _, file := range files {
		if file.IsDir() { continue }

		res1, e := regexp.MatchString(`[a-z0-9]+\.css$`, file.Name())
		if e == nil && res1 {
			cssFile = append(cssFile, "/record/static/css/" + file.Name())
		}
    }

	// var data map[string]interface{}
	// data = make(map[string]interface{})
	// data["JsFiles"] = jsFiles

	/*data := map[string]interface{}{
		"JsFiles": jsFiles,
	}*/

	content, err := common.LazyRender(c, "index", map[string]interface{}{
		"JsFiles": jsFiles,
		"CssFiles": cssFile,
	})

	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.HTML(http.StatusOK, content)
}

// hello world
func HelloWorld() (int, string, error) {
    return http.StatusOK, "Hello, World", nil
}

// template render
func Template(c echo.Context) error {
	
	u, _ := users.FindById(1)
	if u != nil {
		fmt.Printf("u %T %+v %v\n", u.Id, u.Id, u.Id)
	}

	// u := users.FindById(1)

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