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
	Errors struct {
		Login string
		Password string
	}
}


func (cv *CustomValidator) Validate(i interface{}) error {
    if err := cv.Validator.Struct(i); err != nil {
      // Optionally, you could return the error to give each route more control over the status code
      return err
    }
    return nil
}

//
// login page
// 
func Login(c echo.Context) (int, error) {
	login := c.FormValue("login")
	password := c.FormValue("password")
	fmt.Printf("login: %s   pass: %s\n", login, password);

	var user User
	err := c.Bind(&user)
	if err != nil {
		return http.StatusBadRequest, err
	}


	data := LoginPageData{
		Login: login,
		Errors: Errors{},
	}

	
	isValid, errors := validateData(c, user)
	if !isValid {
		data.Errors = errors
	}

	// here login
	
	
	c.Render(http.StatusOK, "login.html", data)
    return http.StatusOK, nil
}


func validateData (c echo.Context, user User) (bool, Errors) {
	ret := Errors {}

	if err := c.Validate(user); err != nil {
		// return http.StatusBadRequest, err

		for _, err1 := range err.(validator.ValidationErrors) {

			if (err1.Field() == "Login"){
				switch err1.Tag() {
				case "required":
					ret.Login = "Logon requried"
				default:
					// c.Error(err1)
					ret.Login = "Error field"
				}
			}

			if (err1.Field() == "Password"){
				switch err1.Tag() {
				case "required":
					ret.Password = "Password requried"
				default:
					//c.Error(err1)
					ret.Password = "Error field"
				}
			}


		}

		return false, ret
	}

	return true, ret
}