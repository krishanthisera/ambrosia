package models

import (
	"fmt"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"

	"gorm.io/gorm"
)

var DB *gorm.DB

func DbConnect() {

	dsn := "host=localhost user=user password=123 dbname=topics port=5432 sslmode=disable "
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}
	fmt.Println("DB Connected")
	database.AutoMigrate(&Topic{})

	DB = database
}
