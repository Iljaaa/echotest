package common

import "fmt"
import "net/http"
import "strconv"
import "time"

// echo 
import "github.com/labstack/echo/v4"

// crypt
import "golang.org/x/crypto/bcrypt"


import "github.com/Iljaaa/echotest/src/models/users"; 

//
// Auth cookie name
//
const AuthCookieName = "4499_auth"

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {

        cookie, err := c.Cookie(AuthCookieName)

	    fmt.Printf("error: %+v, %+v\n", err, cookie)

	    if err != nil {
    		return echo.NewHTTPError(http.StatusUnauthorized, err)
    	}

        // parse cookie string
        if cookie.Value == "" {
            return echo.NewHTTPError(http.StatusUnauthorized, "Empty cookie value")
        }
        
        // convert user id to int
        userId, convErr := strconv.Atoi(cookie.Value)
        if convErr != nil || !(userId > 0)  {
            return echo.NewHTTPError(http.StatusUnauthorized, "Error user id convert")
        }

        u, findUserError := models.FindById(userId)
        if findUserError != nil || u == nil  {
            return echo.NewHTTPError(http.StatusUnauthorized, findUserError)
        }

        // todo: here check hash
        fmt.Println("is authed")


        
        // return echo.NewHTTPError(http.StatusUnauthorized, "Please provide valid credentials")
        
        // c.Response().Header().Set("Custom-Header", "blah!!!")
        // fmt.Println("auth middle ware")
        // return next(c)

        return  next(c)
    }
}

// hash of password
func HashPassword (password string) (string, error) {
    hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    return string(hash), err
}

// Compate password
func ComparePassword (hash string, password string) error {
    return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

// Write auth cookie
func WriteAuthCookie(c echo.Context, userId int) error {
	cookie := new(http.Cookie)
	cookie.Name = AuthCookieName
    // todo: generate valuse
	cookie.Value = strconv.Itoa(userId)
	// cookie.Expires = time.Now().Add(24 * time.Hour)
	c.SetCookie(cookie)
    return nil
}

// clear auth cookie
func ClearAuthCookie(c echo.Context) error {
    cookie := new(http.Cookie)
	cookie.Name = AuthCookieName
    // todo: generate valuse
	cookie.Value = ""
	cookie.Expires = time.Now() // 1.Add(24 * time.Hour)
	c.SetCookie(cookie)
    return nil
}