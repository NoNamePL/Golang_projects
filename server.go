package main

import (
	"database/sql"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	_ "github.com/lib/pq"
	"log"
	"os"
)

type todo struct {
	Item string
}

func main() {

	connStr := "postgres://postgres:postgrespw@localhost:32768/fiber_DB?sslmode=disable"

	// Connect to database

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return indexHandler(c, db)
	}) // add This

	app.Post("/", func(c *fiber.Ctx) error {
		return postHandler(c, db)
	}) // add This

	app.Put("/update", func(c *fiber.Ctx) error {
		return putHandler(c, db)
	}) // add This

	app.Delete("/delete", func(c *fiber.Ctx) error {
		return deleteHandler(c, db)
	}) // Add this

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	app.Static("/", "/public") // add this before starting the app
	log.Fatalln(app.Listen(fmt.Sprintf(":%s", port)))
}

func indexHandler(c *fiber.Ctx, db *sql.DB) error {
	var res string
	var todos []string

	rows, err := db.Query("Select *FROM todos")

	defer rows.Close()
	if err != nil {
		log.Fatal(err)
		c.JSON("An error occured")
	}

	for rows.Next() {
		rows.Scan(&res)
		todos = append(todos, res)
	}
	return c.Render("index", fiber.Map{
		"Todos": todos,
	})

}

func postHandler(c *fiber.Ctx, db *sql.DB) error {
	newTodo := todo{}
	if err := c.BodyParser(&newTodo); err != nil {
		log.Printf("An error occured: %v", err)
		return c.SendString(err.Error())
	}
	fmt.Printf("%v", newTodo)
	if newTodo.Item != "" {
		_, err := db.Exec("INSERT INTO  todos VALUES ($1)", newTodo.Item)
		if err != nil {
			log.Fatal("An error occured while exciting query: %v", err)
		}
	}
	return c.Redirect("/")
}

func putHandler(c *fiber.Ctx, db *sql.DB) error {
	olditem := c.Query("olditem")
	newitem := c.Query("newitem")
	db.Exec("UPDATE todos Set item=$1 WHERE item=$2", newitem, olditem)
	return c.Redirect("/")
}

func deleteHandler(c *fiber.Ctx, db *sql.DB) error {
	todoToDelete := c.Query("item")
	db.Exec("DELETE from todos WHERE item=$1", todoToDelete)
	return c.SendString("deleted")
}
