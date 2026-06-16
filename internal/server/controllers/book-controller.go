package bookController

import (
	"encoding/json"
	"net/http"

	bookService "minecraft/internal/server/services"
)

type BookController struct {
	service *bookService.BookService
}

func NewBookController(service *bookService.BookService) *BookController {
	return &BookController{
		service: service,
	}
}

func (c *BookController) CreateBook(w http.ResponseWriter, r *http.Request) {
	var book bookService.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	createdBook, err := c.service.Create(book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdBook)
}

func (c *BookController) GetBooks(w http.ResponseWriter, r *http.Request) {
	books, err := c.service.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func (c *BookController) GetBook(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	book, err := c.service.GetByID(id)
	if err != nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(book)
}

func (c *BookController) UpdateBook(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	var updatedBook bookService.Book
	err := json.NewDecoder(r.Body).Decode(&updatedBook)

	if err != nil {
		http.Error(w, "Invalid request Payload", http.StatusBadRequest)
		return
	}

	_, err = c.service.Update(id, updatedBook)
	if err != nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedBook)
}

func (c *BookController) DeleteBook(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	err := c.service.Delete(id)
	if err != nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}
