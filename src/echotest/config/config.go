package config

import "fmt"

// Db config
type dbConfig struct {
	HOST		string
	PORT		string
	USER		string
	PASSWORD	string
}

type Configuration struct {
    DB   dbConfig
}

var conf Configuration;

func init() {
	conf = Configuration{}
	conf.DB = dbConfig{
		HOST: "asdasdas",
		PORT: "123",
		USER: "user",
		PASSWORD: "password",
	}
	// conf = Configuration{
	// 	db: dbConfig{
	// 		host: "asdasdas",
	// 	}
	// }
}

func GetConfig() Configuration {

	fmt.Printf("aaa: %+v\n", conf)
    return conf
}