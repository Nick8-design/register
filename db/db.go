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
	psw:=os.Getenv("dbpas")
	if psw==""{
		log.Fatal("error missing db url")
	}

	 url:="postgresql://neondb_owner:"+psw+"@ep-long-bonus-a8t1sjn8-pooler.eastus2.azure.neon.tech/neondb?sslmode=require"



	Db,err=gorm.Open(postgres.Open(url),&gorm.Config{});

	if err!=nil{
		log.Fatal("error Connecting to db")
	}

	Db.AutoMigrate(&models.RegModel{})
	fmt.Println("Successfully connected and intiallized db")

}