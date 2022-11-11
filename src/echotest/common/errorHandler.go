package common

import "fmt"
import "net/http"

// echo
import "github.com/labstack/echo/v4"

type ErrorData struct {
    Code int;
    Message interface{};
}

//
// Error handler
//
func CustomHTTPErrorHandler(err error, c echo.Context) {

    fmt.Printf("CustomHTTPErrorHandler: %v\n", err)

    data := ErrorData{
        Code: http.StatusInternalServerError,
    }

    // var msg interface{}

    if he, ok := err.(*echo.HTTPError); ok {
        data.Code = he.Code
        if isDebug() {
            data.Message = he.Message
        }
        
    }

    // work with codes
    switch data.Code {
        case http.StatusNotFound: // 404
            c.Logger().Info(fmt.Sprintf("%d not found; request: %s", http.StatusNotFound, c.Request().URL))
            c.Render(data.Code, "error404.html", data)
            return
        case http.StatusUnauthorized: // 401
        c.Logger().Info(fmt.Sprintf("%d not found; request: %s", http.StatusNotFound, c.Request().URL))
            c.Render(data.Code, "error401.html", data)
            return
        default : 
            c.Logger().Error(err)
            c.Render(data.Code, "error.html", data)
    }

   
}