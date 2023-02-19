package main

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price int
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	//Migrate the scheme

	db.AutoMigrate(&Product{})

	//Create

	db.Create(&Product{
		Code:  "D42",
		Price: 100,
	})

	//Read
	var product Product

	db.First(&product, 1)                 //find product with integer primary key
	db.First(&product, "code = ?", "D42") // find product with codeD42

	//Update - update product's price to 200
	db.Model(&product).Update("Price", 200)
	//Update - update multiply fields
	db.Model(&product).Updates(Product{Code: "F42", Price: 200}) // non-zero fields
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	//fmt.Println(db)
	fmt.Println(product.Code)
	//Delete - delete product
	db.Delete(&product, 1)
}
