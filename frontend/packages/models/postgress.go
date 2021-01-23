package models

import (

	//"github.com/jinzhu/gorm"
	"fmt"
	"os"

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

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("AMB_DB_HOST"), os.Getenv("AMB_DB_USER"), os.Getenv("AMB_DB_PASS"),
		os.Getenv("AMB_DB_NAME"), os.Getenv("AMB_DB_PORT"))
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}
	fmt.Println("DB Connected")
	database.AutoMigrate(&Topic{})

	DB = database
}
