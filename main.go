package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("hello from Ordering Items Project")
	godotenv.Load()
	dbName := os.Getenv("GO_ORDER")

	connection := dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(connection), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}
