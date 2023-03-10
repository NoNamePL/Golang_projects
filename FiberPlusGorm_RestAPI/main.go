package main

import (
	"FiberPlusGorm_RestAPI/book"
	"FiberPlusGorm_RestAPI/database"
	"fmt"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
)

func initDataBase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "books.db")
	if err != nil {
		panic("failed to connect to database")
	}
	fmt.Println("Database successfully opened")

	database.DBConn.AutoMigrate(&book.Book{})
	fmt.Println("Database Migrated")
}

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/book", book.GetBooks)
	app.Get("/api/v1/book/:id", book.GetBook)
	app.Post("/api/v1/book", book.NewBook)
	app.Delete("/api/v1/book/:id", book.DeleteBook)
}

func main() {
	app := fiber.New()
	initDataBase()
	defer database.DBConn.Close()

	setupRoutes(app)

	err := app.Listen(3000)
	if err != nil {
		log.Fatal(err)
	}
}
