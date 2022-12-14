package controllers

// import "fmt"
import "net/http"
// import "html/template"
// 

// echo
import "github.com/labstack/echo/v4"

// templates
import "github.com/Iljaaa/echotest/src/common"

type ProfileController struct {

}

func (pc *ProfileController) Profile(c echo.Context) error {

	// buf := &bytes.Buffer{}
    
	// fmt.Printf("err: %+v", err)
	// todo: lazy load

	data := map[string]interface{}{
		"testParam" : "testValue",
	}

	// fmt.Println("123")

	content, err := common.LazyProfileRender(c, "profile.index", data)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.HTML(http.StatusOK, content)
}
