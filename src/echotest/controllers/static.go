package controllers

import "fmt";
import "net/http"
import "errors"

func HelloWorld() (int, string, error) {
	fmt.Printf("hellow world controler")
    return http.StatusOK, "Hello, World", nil
}

// exmaple with error
func Error() (int, string, error) {
	return 400, "", errors.New("test error")
}