package main

import (
	"fmt"
	"log"
	"os"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

func load() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Err load .env")
	}
}

func main() {
	load()
	userName := fmt.Sprintf(os.Getenv("USER_NAME"))
	password := fmt.Sprintf(os.Getenv("PASSWORD"))
	dbName := fmt.Sprintf(os.Getenv("DATABASE_NAME"))
	fmt.Println(userName, password, dbName)
}
