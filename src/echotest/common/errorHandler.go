package common

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

    // fmt.Printf("data: %v\n", msg)
    // data.Message = msg

    // code 401 requree auth
    if data.Code == 401 {
        c.Render(data.Code, "error401.html", data)
        return
    }

    c.Render(data.Code, "error.htl", data)
}