package db

import (
	"fmt"
	"kt_go/models"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func ConnectDb(){
	err:=godotenv.Load()
	if err!=nil{
		log.Fatal("error loading .env ",err)
	}
	dsn:=os.Getenv("DB_URL")

	if dsn==""{
		log.Fatal("error missing db url")
	}

	Db,err=gorm.Open(postgres.Open(dsn),&gorm.Config{});

	if err!=nil{
		log.Fatal("error Connecting to db")
	}

	Db.AutoMigrate(&models.RegModel{})
	fmt.Println("Successfully connected and intiallized db")

}