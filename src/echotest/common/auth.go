package common

import "fmt"
import "net/http"
import "strconv"
import "time"
import "errors"

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

        // get user 
        u, err := getUserFromCoockie(c)
        if err != nil || u == nil  {
            return echo.NewHTTPError(http.StatusUnauthorized, err)
        }

        // renew coockie

        // todo: here check hash
        fmt.Println("is authed")
        
        // return echo.NewHTTPError(http.StatusUnauthorized, "Please provide valid credentials")
        
        // c.Response().Header().Set("Custom-Header", "blah!!!")
        // fmt.Println("auth middle ware")
        // return next(c)

        return next(c)
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

// clear auth data
func ClearAuthData(c echo.Context) error {
    // clear cookie
    cookie := new(http.Cookie)
	cookie.Name = AuthCookieName
	cookie.Value = ""
	cookie.Expires = time.Now() // 1.Add(24 * time.Hour)
	c.SetCookie(cookie)
    return nil
}

type UserData struct {
	Name string
}

// get auth user data
func GetAuthUser (c echo.Context) (*UserData, error) {

    u, err := getUserFromCoockie(c)
    if err != nil || u == nil  {
        return nil, err
    }

    userData := UserData{
        Name : u.Name,
    }

    return &userData, nil
}

//
// Get cookie, check cookie, find user, valid cooke hash
//
func getUserFromCoockie (c echo.Context) (*users.User, error) {
    
    // 
    cookie, err := c.Cookie(AuthCookieName)
    if err != nil {
        return nil, err
    }

    // parse cookie string
    if cookie.Value == "" {
        return nil, errors.New("Empty cookie value")
    }

    // convert user id to int
    userId, convErr := strconv.Atoi(cookie.Value)
    if convErr != nil || !(userId > 0)  {
        return nil,  errors.New("Error user id convert")
    }

    u, findUserError := users.FindById(userId)
    if findUserError != nil || u == nil  {
        return nil, findUserError
    }

    return u, nil
}