package database

import (
	"database/sql"
	"log"
	"go-todo/config" // import the config package
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

// Connect to MySQL database
func Connect() {
	dsn := config.GetDSN() // Get the DSN (Data Source Name) from the config package

	// Connect to the MySQL database using the DSN
	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal("Error pinging the database: ", err)
	}

	log.Println("Connected to MySQL database")
}
