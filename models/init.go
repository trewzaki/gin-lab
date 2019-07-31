package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func InitialDatabase() {
	var connectErr error

	// connectionString := fmt.Sprintf("host=myhost port=myport user=gorm dbname=gorm password=mypassword")
	connectionString := fmt.Sprintf("host=35.240.191.214 port=5432 user=postgres dbname=fillgoods-lab password=IkDD43cIamwavO7s")
	DB, connectErr = gorm.Open("postgres", connectionString)
	if connectErr != nil {
		fmt.Println("Connect database failed.")
		return
	}

	autoMigrateDatabase()
	// defer db.Close()
	fmt.Println("Database connection successfully...")
}

func autoMigrateDatabase() {
	DB.AutoMigrate(&User{})
}
