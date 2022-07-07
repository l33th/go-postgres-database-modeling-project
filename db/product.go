package db

import "time"

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
