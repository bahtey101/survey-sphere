package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	host := "194.226.49.210"
	user := "sudo_user"
	password := "sudo_nikita_plus_dimon"
	dbname := "survey_sphere_db"
	port := "5432"

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbname, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	print(err)
	if err != nil {
		panic("Failed to connect to database")
	}

	return db
}
