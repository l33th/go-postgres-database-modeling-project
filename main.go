package main

import (
	// Third party packages
	"go-postgres-database-modeling/db"

	// Default packages
	"log"
)

// main is the entry point
func main() {
	log.Println("Hello, World!")
	db.Connect()
}