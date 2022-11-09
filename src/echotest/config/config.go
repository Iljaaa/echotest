package config

import "fmt"

// Db config
type dbConfig struct {
	HOST		string
	PORT		int
	USER		string
	PASSWORD	string
	DB			string
}

type Configuration struct {
    DB   dbConfig
}

var conf Configuration;

func init() {
	conf = Configuration{}
	conf.DB = dbConfig{
		HOST: "db",
		PORT: 5432,
		USER: "postgres",
		PASSWORD: "example",
		DB: "sportscools",
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