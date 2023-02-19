package book

import (
	"FiberPlusGorm_RestAPI/database"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

type Book struct {
	gorm.Model
	Title  string `json:"title"`
	Author string `json:"author"`
	Rating int    `json:"rating"`
}

func GetBooks(c *fiber.Ctx) {
	//Записываем в переменную ссылку на нашу БД
	db := database.DBConn
	var books []Book
	//ищем все книги
	db.Find(&books)
	//выводим все книги
	err := c.JSON(books)
	if err != nil {
		panic("you can't get these books")
	}
}

func GetBook(c *fiber.Ctx) {
	//создаем id как в main.go(или что там указано в routes) для поиска
	id := c.Params("id")
	db := database.DBConn

	var book Book
	// ищем отдельную книгу
	db.Find(&book, id)
	//выводим
	err := c.JSON(book)
	if err != nil {
		panic("you can't get this book")
	}
}

func NewBook(c *fiber.Ctx) {
	db := database.DBConn
	// создаем книгу
	book := new(Book)
	if err := c.BodyParser(book); err != nil {
		c.Status(503).Send(err)
		return
	}
	// записываем созданную книгу в БД
	db.Create(&book)
	//	выводим книгу
	err := c.JSON(book)
	if err != nil {
		panic("you can't create a book")
	}
}

func DeleteBook(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var book Book
	db.First(&book, id)

	if book.Title == "" {
		c.Status(500).Send("No book found with given ID ")
		return
	}

	db.Delete(&book)
	c.Send("book successfully deleted ")
}
