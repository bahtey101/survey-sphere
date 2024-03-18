package database

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbase *gorm.DB

func Init() *gorm.DB {
	host := "194.226.49.210"
	user := "sudo_user"
	password := "sudo_nikita_plus_dimon"
	dbname := "survey_sphere_db"
	port := "5432"

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbname, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database")
	}

	return db
}

func GetDataBase() *gorm.DB {
	var sleep = time.Duration(1)

	for dbase == nil {
		fmt.Printf("Database is not available. Wait for %d sec.\n", sleep)
		time.Sleep(time.Second)
		dbase = Init()
	}

	return dbase
}
