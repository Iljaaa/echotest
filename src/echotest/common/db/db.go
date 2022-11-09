package db


// pgx
// import "github.com/jackc/pgx/v5"
import "github.com/jackc/pgx/v5/pgxpool"

import "fmt"
import "os"
import "context"

import "github.com/Iljaaa/echotest/src/config";


func Test(){

}

//
// connection
// 
var PullCon *pgxpool.Pool;


func InitConnection() {

    config := config.GetConfig()

    dbConnectString := fmt.Sprintf ("postgres://%s:%s@%s:%d/%s", config.DB.USER, config.DB.PASSWORD, config.DB.HOST, config.DB.PORT, config.DB.DB)
    fmt.Printf("connect to db: %s\n", dbConnectString)

	pull, err := pgxpool.New(context.Background(), dbConnectString)
	if err != nil {
	  fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	  os.Exit(1)
	}
	
	fmt.Printf("pool: %T %+v\n", pull, pull)

	PullCon = pull

	// Connection = pool

	// defer PullCon.Close()

    /*conn, err := pgx.Connect(context.Background(), dbConnectString)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	Connection = conn*/


	//defer conn.Close(context.Background())

	fmt.Println("connected")
}


func GetPool () *pgxpool.Pool {
	return PullCon
}