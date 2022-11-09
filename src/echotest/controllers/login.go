package controllers

import "fmt"
import "net/http"
// import "errors"

import "github.com/labstack/echo/v4"

import "github.com/go-playground/validator"
// import "github.com/labstack/echo/v4/middleware"


type (
	User struct {
		Login  string `form:"login" validate:"required"`
		Password string `form:"password" validate:"required"`
	}

	CustomValidator struct {
		Validator *validator.Validate
	}
)


type Errors struct {
	Login string
	Password string
}

type LoginPageData struct {
	Login string
	Errors map[string]string
	/*Errors struct {
		Login string
		Password string
	}*/
}


func (cv *CustomValidator) Validate(i interface{}) error {
    if err := cv.Validator.Struct(i); err != nil {
      // Optionally, you could return the error to give each route more control over the status code
      return err
    }
    return nil
}


var errorDescription map[string]string

func init() {
	errorDescription = map[string]string {
		"Login.required": "Login requried",
		"Password.required": "Passoword requried",
	}
}

//
// login page
// 
func Login(c echo.Context) (int, error) {
	// login := c.FormValue("login")
	// password := c.FormValue("password")

	var user User
	err := c.Bind(&user)
	if err != nil {
		return http.StatusBadRequest, err
	}

	data := LoginPageData{
		Login: user.Login,
	}
	
	isValid, errors := validateData(c, user)
	if !isValid {
		data.Errors = errors
	}

	// here login
	
	c.Render(http.StatusOK, "login.html", data)
    return http.StatusOK, nil
}


func validateData (c echo.Context, user User) (bool, map[string]string) {
	errorssss := make(map[string]string)
	err := c.Validate(user);

	if err != nil {
		// return http.StatusBadRequest, err

		for _, err1 := range err.(validator.ValidationErrors) {

			field := err1.Field()

			errorDescriptionCode := fmt.Sprint(field, ".", err1.Tag())

			description, exists := errorDescription[errorDescriptionCode]

			if !exists {
				description = fmt.Sprint("Error '", err1.Tag(), "' on field '", field, "'")
			}


			_, errorExist := errorssss[field]
			if !errorExist { errorssss[field] = description }
		}

		return false, errorssss
	}

	// validate

	return true, errorssss
}