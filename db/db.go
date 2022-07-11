package db

import (
	// Third party packages
	"github.com/go-pg/pg"

	// Default packages
	"log"
	"os"
)
// Variable db is of type pg.DB
var db *pg.DB

// Connect will establish a database connection
func Connect() {
	// Credentials setup
	options := &pg.Options {
		User: "postgres",
		Password: "lth-dev",
		Addr: "localhost:5432",
	}
	// Connection setup
	db = pg.Connect(options)
	if db == nil {
		log.Printf("Failed to connect to db.\n")
		os.Exit(100)
	}
	log.Println("Connection to db successful.")

	err := CreateProductItemsTable(db)
	if err != nil {
		log.Println("Err: ", err)
	}

	err = db.Close(nil)
	if err != nil {
		log.Printf("Error occured when trying to close the connection: %v\n", err)
		os.Exit(100)
	}
	log.Println("Connection closed successfully.")
}