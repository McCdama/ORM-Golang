package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	gorm.Model
	Email    string `gorm:"column:email" json:"email"`
	Password string `gorm:"column:password" json:"-"`
}

func (User) TableName() string {
	return "users"
}

type UserInfo struct {
	User      User   `gorm:"foreignkey:u_id;association_foreignkey:id"`
	UID       uint   `gorm:"column:u_id" json:"-"`
	FirstName string `gorm:"column:first_name" json:"first_name"`
	LastName  string `gorm:"column:last_name" json:"last_name"`
	Phone     string `gorm:"column:phone" json:"phone"`
	Address   string `gorm:"column:address" json:"address"`
}

func (UserInfo) TableName() string {
	return "user_info"
}

type Order struct {
	gorm.Model
	OrderType int
	Dummy     string
}
type Ajassem struct {
	gorm.Model
	Dummy int
}

func main() {
	db, err := gorm.Open("mysql", "root:0000@/test_gorm?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect")
	}
	defer db.Close()

	/* var hasTable = db.HasTable(&User{})
	if hasTable == true {
		fmt.Println("Table exist :) ")
	} */
	db.AutoMigrate(&User{}, &Order{}, &Ajassem{}, &UserInfo{})
	//	fmt.Println("Creating")
	//	db.CreateTable(&Ajassem{})
	//fmt.Println("Dropping")
	//db.DropTable(&User{}, &Ajassem{}, &Order{})

	fmt.Println("Modify")
	// change column order_type's data type to `text` for model `Order`
	db.Model(&Order{}).ModifyColumn("order_type", "text")
	//db.Model(&User{}).ModifyColumn("name", "long_text")

	db.Model(&Order{}).DropColumn("dummy")
	db.Model(&UserInfo{}).AddForeignKey("u_id", "users(id)", "RESTRICT", "RESTRICT")

}
