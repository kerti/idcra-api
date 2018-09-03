package context

import (
	"fmt"
	"log"
	"time"

	// Blank import for MySQL
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func OpenDB(config *Config) (*sqlx.DB, error) {
	log.Println("Database is connecting... ")
	db, err := sqlx.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", config.DBUser, config.DBPassword, config.DBHost, config.DBPort, config.DBName))

	if err != nil {
		log.Println("Error: " + err.Error())
		panic(err.Error())
	}

	if err = db.Ping(); err != nil {
		log.Println("Error: " + err.Error())
		log.Println("Retry database connection in 5 seconds... ")
		time.Sleep(time.Duration(5) * time.Second)
		return OpenDB(config)
	}
	log.Println("Database is connected ")
	return db, nil
}
