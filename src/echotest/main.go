package main

import "fmt";
import "github.com/Iljaaa/echotest/src/common";
import "github.com/Iljaaa/echotest/src/config";
// import "echo_test/common";

func main() {
    fmt.Println("startng test script.....")
    common.Test()
    // fmt.Sprintf("asdasdas %T", config.Configuration)
    config := config.GetConfig()
    a := "DB"
    fmt.Printf("%+v\n", config.DB)
    // fmt.Printf("%+v\n", config.db)
    
}