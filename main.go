package main

import (
	"fmt"
	"log"
	"sharath/database"
	"sharath/api"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

func initDatabase(){
	var err error
	database.DBConn , err = gorm.Open("mysql",database.DNS)
	if err!=nil{
		panic("Failed to connect Database")
	}
	fmt.Println("Database connected successfully !")
	database.DBConn.AutoMigrate(&api.Notes{})
	database.DBConn.AutoMigrate(&api.User{})
	fmt.Println("Database Migrated !")
}

func setupRoutes(app *fiber.App){
	app.Post("/signup",api.RegisterNewUser)
	app.Post("/login",api.UserLogin)
	app.Get("/:sid/notes",api.GetAllNotes)
	app.Post("/:sid/notes",api.CreateNote)
	app.Delete("/:sid/:id/notes",api.DeleteNote)
}

func main(){
	app:=fiber.New()
	initDatabase()
	setupRoutes(app)
	log.Fatal(app.Listen(":8088"))
	defer database.DBConn.Close()
}