package jwt

import "fmt"
import "time"
import "encoding/json"
import "net/http"
import "errors"
// import "strconv"

// echo 
import "github.com/labstack/echo/v4"

// jwt
import "github.com/golang-jwt/jwt/v4"

import "github.com/Iljaaa/echotest/src/config"

var jwtSecretKey string

func init (){
	fmt.Println("Init jwt")
	jwtSecretKey = config.GetConfig().JwtSecret
	fmt.Printf("jwtSecretKey %s\n", jwtSecretKey)
}

//
// jwt middleware check
//
func JwtMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {

		data, err := getRequestData(c)
		if err != nil {
			return c.JSON(http.StatusOK, map[string]interface{}{
				"result": false,
				"error" : "wrong data",
			})
		}

		token, exists := data["token"]
		if !exists {
			return c.JSON(http.StatusOK, map[string]interface{}{
				"result": false,
				"error" : "token not presents",
			})
		}

		// choeck token 
		// genToken, _ := CreateToken()
        // fmt.Printf("genToken: %s\n", genToken) 

		isValid, checkErrors := checkToken(token.(string))
		if isValid == false || checkErrors != nil {
			c.Logger().Error(checkErrors)
			return c.JSON(http.StatusOK, map[string]interface{}{
				"result": false,
				"error" : fmt.Sprintf("token error (%s)", checkErrors),
			})
		}

        // token is ok
        return next(c)
    }
}

//
// get request data
//
func getRequestData(c echo.Context) (map[string]interface{}, error) {

	jsonBody := make(map[string]interface{})

	err := json.NewDecoder(c.Request().Body).Decode(&jsonBody)
	if err != nil {
		return jsonBody, err
	}
	
	return jsonBody, nil
}

//
// create token
//
func CreateToken () (string, error) {
	fmt.Printf("unix %T\n", time.Now().Unix())
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo": "bar",
		"nbf": time.Now().Unix(), // <- not defore https://pkg.go.dev/github.com/golang-jwt/jwt/v4#RegisteredClaims
	})

	// mySigningKey := []byte("my_secret_key")

	tokenString, err := token.SignedString([]byte(jwtSecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

//
// Check token
//
func checkToken (tokenString string) (bool, error) {
	// sample token string taken from the New example
	// tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIiLCJuYmYiOjE0NDQ0Nzg0MDB9.u1riaD1rW97opCoAuRCTy4w58Br-Zk-bh7vLiRIsrpU"

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New(fmt.Sprintf("Unexpected signing method: %v", token.Header["alg"])) //  fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(jwtSecretKey), nil
		// return []byte("my_secret_key"), nil
	})

	if err != nil {
		return false, err
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
	// if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// parseTime, err := time.Parse("Jan 02, 2006", "Sep 30, 2021")

		// fmt.Printf("nbf: %T\n", claims["nbf"].(float64))
		// n, _ := strconv.ParseInt(claims["nbf"].(string), 10, 64);

		// parseTime := time.Unix(claims["nbf"].(int64), 0)
		// parseTime, err := time.Parse(time.RFC1123, claims["nbf"].(string))
		
		// fmt.Printf("parseTime: %v\n", parseTime)
		// fmt.Printf("err: %v\n", err)

		// getiing additional params
		// fmt.Println(claims["foo"], claims["nbf"])

		return true, nil

	} 
	
	return false, errors.New("Somthing wrong witch claims")
}