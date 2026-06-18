// handler.go
package book

import (
	"encoding/json"
	"minecraft/internal/common"
	"net/http"

	"github.com/google/uuid"
)

func SetupBooksRouter(handler *Handler) {
	http.HandleFunc("GET /books", handler.GetBooks)
	http.HandleFunc("GET /books/{id}", handler.GetBook)
	http.HandleFunc("POST /books", handler.CreateBook)
	http.HandleFunc("PUT /books/{id}", handler.UpdateBook)
	http.HandleFunc("DELETE /books/{id}", handler.DeleteBook)
}

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}

// CreateBook godoc
//
//	@Summary		request to create a book
//	@Description	create new book in DB
//	@Tags			books
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	book.Book
//	@Failure		400
//	@Failure		500
//	@Router			/books [post]
func (c *Handler) CreateBook(w http.ResponseWriter, r *http.Request) {
	var book Book
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

// GetBooks godoc
//
//	@Summary		get books list
//	@Description	get array of books by params
//	@Tags			books
//	@Produce		json
//	@Param			ids		query	[]string	false	"Book IDs"
//	@Param			name	query	string		false	"Book name"
//	@Param			author	query	string		false	"Book author"
//	@Param			isbn	query	string		false	"Book isbn"
//	@Success		200		{array}	book.Book
//	@Failure		400
//	@Failure		500
//	@Router			/books [get]
func (c *Handler) GetBooks(w http.ResponseWriter, r *http.Request) {
	books, err := c.service.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// GetBook godoc
//
//	@Summary		get single
//	@Description	get book by params
//	@Tags			books
//	@Produce		json
//	@Param			id	path		string	false	"Book ID"
//	@Success		200	{object}	book.Book
//	@Failure		400
//	@Failure		404
//	@Failure		500
//	@Router			/books/{id} [get]
func (c *Handler) GetBook(w http.ResponseWriter, r *http.Request) {
	var q common.BaseItemRequestDto

	id, parseError := uuid.Parse(r.URL.Query().Get("id"))
	if parseError != nil {
		http.Error(w, "ID must be a uuid", http.StatusBadRequest)
	}
	q.Id = id

	book, err := c.service.GetByID(q.Id)
	if err != nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(book)
}

// UpdateBook godoc
//
//	@Summary		update Book by id
//	@Description	update book in DB
//	@Tags			books
//	@Accept			json
//	@Produce		json
//	@Param			id			path		string		false	"Book ID"
//	@Param			updateDto	body		book.Book	true	"Book update fields"
//	@Success		200			{object}	book.Book
//	@Failure		404
//	@Failure		500
//	@Router			/books [put]
func (c *Handler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	var q common.BaseItemRequestDto

	id, parseError := uuid.Parse(r.URL.Query().Get("id"))
	if parseError != nil {
		http.Error(w, "ID must be a uuid", http.StatusBadRequest)
	}
	q.Id = id

	var updatedBook Book
	err := json.NewDecoder(r.Body).Decode(&updatedBook)

	if err != nil {
		http.Error(w, "Invalid request Payload", http.StatusBadRequest)
		return
	}

	_, err = c.service.Update(q.Id, updatedBook)
	if err != nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedBook)
}

// DeleteBook godoc
//
//	@Summary		delete Book by id
//	@Description	delete book in DB
//	@Tags			books
//	@Param			id	path		string	false	"Book ID"
//	@Success		200	{object}	book.Book
//	@Failure		404
//	@Failure		500
//	@Router			/books [delete]
func (c *Handler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	var q common.BaseItemRequestDto

	id, parseError := uuid.Parse(r.URL.Query().Get("id"))
	if parseError != nil {
		http.Error(w, "ID must be a uuid", http.StatusBadRequest)
	}
	q.Id = id

	err := c.service.Delete(q.Id)
	if err != nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}
