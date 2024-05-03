package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"gin-books-api/cmd/controller"
	"gin-books-api/cmd/controller/model"
)

var books = []*model.Book{
	{ID: "1", ISBN: "9789574157327", Title: "腹語術", Author: "夏宇", Price: 340.00, Stock: 180},
	{ID: "2", ISBN: "9789574373383", Title: "脊椎之軸", Author: "夏宇", Price: 777.00, Stock: 80},
	{ID: "3", ISBN: "9789571375922", Title: "大裂", Author: "胡遷", Price: 380.00, Stock: 120},
	{ID: "4", ISBN: "9789864061471", Title: "遠處的拉莫", Author: "胡遷", Price: 360.00, Stock: 140},
}

var log = logrus.New()

func init() {
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	log.SetLevel(logrus.InfoLevel)
}

func main() {
	model.Books = books

	router := gin.Default()
	router.GET("/books", controller.GetBooksHandler)
	router.GET("/books/:id", controller.GetBookByIDHandler)
	router.POST("/books", controller.CreateBookHandler)
	router.DELETE("/books/:id", controller.DeleteBookHandler)
	router.PATCH("/books/:id", controller.UpdateBookHandler)
	router.PATCH("/books/checkout", controller.CheckoutBookHandler)

	log.Fatal(router.Run("localhost:8080"))
}
