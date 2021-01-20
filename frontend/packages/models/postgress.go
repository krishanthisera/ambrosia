package models

import (

	//"github.com/jinzhu/gorm"
	"fmt"

	"gorm.io/driver/postgres"

	"gorm.io/gorm"

	_ "github.com/lib/pq"
)

var DB *gorm.DB

type PsqlConnect struct {
	Host     string
	Port     int64
	User     string
	Password string
	Dbname   string
}

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
