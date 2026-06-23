// handler.go
package book

import (
	"encoding/json"
	"minecraft/internal/common"
	"net/http"
	"strconv"

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
//	@Summary		Create a new book
//	@Description	Creates a new book in DB
//	@Tags			books
//	@Accept			json
//	@Produce		json
//	@Param			book body book.BookCreateRequestDto true "Book details"
//	@Success		201	{object}	book.BookCreateRequestDto
//	@Failure		400
//	@Failure		500
//	@Router			/books [post]
func (c *Handler) CreateBook(w http.ResponseWriter, r *http.Request) {
	var book BookCreateRequestDto

	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := validate.Struct(book); err != nil {
		http.Error(w, "Validation failed: "+err.Error(), http.StatusBadRequest)
		return
	}

	createdBook, err := c.service.Create(book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdBook)
}

// GetBooks godoc
//
//	@Summary		List all books
//	@Description	Returns a paginated list of books with optional search
//	@Tags			books
//	@Produce		json
//	@Param			offset query integer false "offset amount" default(0)
//	@Param			limit query integer false "Items per page" default(20)
//	@Param			search query string false "Search by author or title"
//	@Success		200		book.BookListResponse
//	@Failure		400
//	@Failure		500
//	@Router			/books [get]
func (c *Handler) GetBooks(w http.ResponseWriter, r *http.Request) {
	var params BookListRequest

	if value := r.URL.Query().Get("offset"); value != "" {
		offset, err := strconv.Atoi(value)
		if err != nil {
			// http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		params.Offset = &offset
	}

	if value := r.URL.Query().Get("limit"); value != "" {
		limit, err := strconv.Atoi(value)
		if err != nil {
			// http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		params.Limit = &limit
	}

	if search, exists := r.URL.Query()["search"]; exists {
		params.Search = &search[0]
	}

	if err := params.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	books, err := c.service.GetAll(params)
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
//	@Summary		Update a book
//	@Description	Partially updates an existing book
//	@Tags			books
//	@Accept			json
//	@Produce		json
//	@Param			id			path		string		false	"Book ID"
//	@Param			updateDto	body		book.BookUpdateRequest	true	"Book update fields"
//	@Success		200			{object}	book.BookUpdateResponse
//	@Failure		404
//	@Failure		500
//	@Router			/books/{id} [put]
func (c *Handler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	var q common.BaseItemRequestDto
	var updatedBook BookUpdateRequest

	id, parseError := uuid.Parse(r.URL.Query().Get("id"))
	if parseError != nil {
		http.Error(w, "ID must be a uuid", http.StatusBadRequest)
	}
	q.Id = id

	if err := json.NewDecoder(r.Body).Decode(&updatedBook); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if _, err := c.service.Update(q.Id, updatedBook); err != nil {
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
//	@Router			/books/{id} [delete]
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
