package main

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type TestUser struct {
	gorm.Model
}

/* func (v TestUser) TableName() string {
	fmt.Printf("v: %p %v\n", &v, v)
	return "ajassematy"
} */
type Animal struct {
	AnimalId int64     `gorm:"column:beast_id;primary_key"` // set column name to `beast_id`
	Birthday time.Time `gorm:"column:day_of_the_beast"`     // set column name to `day_of_the_beast`
	Age      int64     `gorm:"column:age_of_the_beast"`     // set column name to `age_of_the_beast`
}

func main() {
	db, err := gorm.Open("mysql", "root:0000@/test_gorm?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect")
	}
	defer db.Close()
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "McCdama_" + defaultTableName
	}

	db.SingularTable(true) // Reihenfolge wichig
	db.AutoMigrate(&TestUser{}, &Animal{})
}
