package main

import (
	"kt_go/db"
	"kt_go/handles"
	"log"
	

	"github.com/gofiber/fiber/v2"
)

func init(){
	db.ConnectDb()
}

func main(){
	app := fiber.New()
    app.Post("/register",handles.Register) 
	app.Post("/login",handles.Login) 
	app.Get("/ping",handles.Ping) 

    log.Fatal(app.Listen(":8080"))
}