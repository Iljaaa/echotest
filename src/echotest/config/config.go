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

// 
type Configuration struct {

	// secret key for jwt middleware
	JwtSecret string

	// db confid
    DB   dbConfig 
}

var conf Configuration;

func init() {
	conf = Configuration{
		JwtSecret: "my_super_secret",
	}
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
	fmt.Printf("GetConfig.: %+v\n", conf)
    return conf
}