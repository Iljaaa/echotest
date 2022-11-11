package controllers

import "fmt"
import "net/http"
// import "errors"

import "github.com/labstack/echo/v4"
import "github.com/go-playground/validator"
// import "github.com/labstack/echo/v4/middleware"

import "github.com/Iljaaa/echotest/src/models/users"

// auth
import "github.com/Iljaaa/echotest/src/common"


type (
	UserData struct {
		Login  string `form:"login" validate:"required"`
		Password string `form:"password" validate:"required"`
	}

	CustomValidator struct {
		Validator *validator.Validate
	}
)


/*type Errors struct {
	Login string
	Password string
}*/

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
// int, error
func Login(c echo.Context) error {	
	c.Render(http.StatusOK, "login.html", LoginPageData{})
	return c.NoContent(http.StatusOK)
}


//
// login page with post data
// 
func LoginPost(c echo.Context) error {
	// login := c.FormValue("login")
	// password := c.FormValue("password")

	var u UserData
	err := c.Bind(&u)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	data := LoginPageData{
		Login: u.Login,
	}
	
	isValid, dbUser, errors := validateData(c, u)
	if !isValid {
		data.Errors = errors
		return c.Render(http.StatusOK, "login.html", data)
	}


	fmt.Printf("isValid %T %+v\n", isValid, isValid)
	fmt.Printf("errors %T %+v\n", errors, errors)
	fmt.Printf("dbUser %T %+v\n", dbUser, dbUser)
	
	// save auth
	err = common.WriteAuthCookie(c, dbUser.Id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	
	return c.Redirect(http.StatusFound, "/profile")
}

func validateData (c echo.Context, user UserData) (bool, *users.User, map[string]string) {

	errors := make(map[string]string)

	validate := c.Validate(user);

	if validate != nil {
		// return http.StatusBadRequest, err

		for _, err1 := range validate.(validator.ValidationErrors) {

			field := err1.Field()

			errorDescriptionCode := fmt.Sprint(field, ".", err1.Tag())

			description, exists := errorDescription[errorDescriptionCode]

			if !exists {
				description = fmt.Sprint("Error '", err1.Tag(), "' on field '", field, "'")
			}

			_, errorExist := errors[field]
			if !errorExist { errors[field] = description }
		}

		return false, nil, errors
	}

	// user find
	dbUser, err := users.FindByLogin(user.Login)
	if err != nil || dbUser == nil {
		errors["Login"] = "User not found"
		return false, nil, errors
	}

	// password hash
	hash, err := common.HashPassword(user.Password)
	if err != nil {
		errors["Password"] = "Password hash not created"
		return false, nil, errors
	}

	// err2 := bcrypt.CompareHashAndPassword([]byte(hash), []byte("123456"))
	
	// exal password
	comparePasswordError := common.ComparePassword(hash, "123456")
	if comparePasswordError != nil {
		errors["Password"] = "Password not correct"
		return false, nil, errors
	}


	// if  {
	// 	errors["Password"] = "Wrong password"
	// }

	return true, dbUser, errors
}


//
// login page with post data
// 
func LoginOut(c echo.Context) error {
	fmt.Println("Profile logout")

	// clear cookie
	err := common.ClearAuthData(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	
	return c.Redirect(http.StatusFound, "/login")
}
