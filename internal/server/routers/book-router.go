package bookRouter

import (
	"net/http"

	bookController "minecraft/internal/server/controllers"
)

func SetupBooksRouter(controller *bookController.BookController) {
	http.HandleFunc("/books", controller.GetBooks)
	http.HandleFunc("/books/create", controller.CreateBook)
	http.HandleFunc("/books/get", controller.GetBook)
	http.HandleFunc("/books/update", controller.UpdateBook)
	http.HandleFunc("/books/delete", controller.DeleteBook)
}
