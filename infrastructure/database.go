package infrastructure

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func load() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Err load .env")
	}
}

type connectData struct {
	load()
	USER string fmt.Sprintf("$s", os.Getenv("USER_NAME"))
	PASSWORD string fmt.Sprintf("$s", os.Getenv("PASSWORD"))
}
