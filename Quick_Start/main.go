package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	db, err := gorm.Open("mysql", "user:password@/test_gorm?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect")
	}
	defer db.Close()

	db.AutoMigrate(&Product{})
	db.Create(&Product{
		Code:  "Mohed",
		Price: 97916,
	})

	var product Product
	db.First(&product, 1) // with id
	fmt.Println(product)

	var productCode Product
	db.First(&productCode, "price = ?", "97916")
	fmt.Println(productCode)

	db.Model(&productCode).Update("Price", 1995)
	db.Delete(&productCode)
}
