package db

import (
	// Third party packages
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"

	// Default packages
	"log"
	"time"
)

// ProductItem creates a product_item table
type ProductItem struct {
	RefPtr    int      `sql:"-"`
	tableName struct{} `sql:"product_items_collection"`
	ID        int      `sql:"id,pk"`
	Name      string   `sql:"name,unique"`
	Desc      string   `sql:"desc"`
	Image     string   `sql:"image"`
	Price     float64  `sql:"price,type:real"`

	Features struct {
		Name string
		Desc string
		Imp  int
	} `sql:"features,type:jsonb"`

	CreatedAt time.Time `sql:"created_at"`
	UpdatedAt time.Time `sql:"updated_at"`
	IsActive  bool      `sql:"is_active"`
}

// CreateProductItemsTable creates a new table
func CreateProductItemsTable(db *pg.DB) error {
	options := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	// Create table
	createErr := db.CreateTable(&ProductItem{}, options)
	if createErr != nil {
		log.Printf("Error while creating table productItems, ERROR: %v\n", createErr)
	}
	log.Println("Table ProductItems created successfully!")
	return nil
}