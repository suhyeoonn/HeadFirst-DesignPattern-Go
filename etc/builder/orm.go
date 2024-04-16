package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetORM() *gorm.DB{
	err := godotenv.Load(".env")
    
    if err != nil {
        log.Fatal("Error loading .env file")
    }
	
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PW")
	database := os.Getenv("DB_DATABASE")

	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, database)
  	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}