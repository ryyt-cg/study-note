package main

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	gormConfig := &gorm.Config{}

	db, err := gorm.Open(sqlite.Open("test.db"), gormConfig)
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	// CREATE TABLE `products` (`id` integer PRIMARY KEY AUTOINCREMENT,`created_at` datetime,`updated_at` datetime,`deleted_at` datetime,`code` text,`price` integer)
	err = db.AutoMigrate(&Product{})
	if err != nil {
		panic("failed to migrate schema")
	}

	newProduct := Product{Code: "D42", Price: 100}
	// Create new Product & assign ID number
	db.Create(&newProduct)

	// Read
	var product Product
	db.First(&product, newProduct.ID)     // find product with integer primary key
	db.First(&product, "code = ?", "D42") // find product with code D42

	var products []Product
	db.Find(&products)

	// Update - update product's price to 200
	db.Model(&product).Update("Price", 200)
	// Update - update multiple fields
	db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Soft Delete - mark all products as deleted.
	for _, p := range products {
		db.Delete(&p, p.ID)
	}

	db.Find(&products)

	fmt.Println("delete new product & exit")
}
