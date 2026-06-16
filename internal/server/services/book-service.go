package bookService

import (
	"errors"
)

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var ErrBookNotFound = errors.New("book not found")

type BookService struct {
	books []Book
}

func NewBookService() *BookService {
	return &BookService{
		books: []Book{},
	}
}

func (s *BookService) Create(book Book) (Book, error) {
	s.books = append(s.books, book)
	return book, nil
}

func (s *BookService) GetAll() ([]Book, error) {
	return s.books, nil
}

func (s *BookService) GetByID(id string) (Book, error) {
	for _, book := range s.books {
		if book.ID == id {
			return book, nil
		}
	}
	return Book{}, ErrBookNotFound
}

func (s *BookService) Update(id string, updatedBook Book) (Book, error) {
	for idx, book := range s.books {
		if book.ID == id {
			s.books[idx] = updatedBook
			return updatedBook, nil
		}
	}
	return Book{}, ErrBookNotFound
}

func (s *BookService) Delete(id string) error {
	for i, book := range s.books {
		if book.ID == id {
			s.books = append(s.books[:i], s.books[i+1:]...)
			return nil
		}
	}
	return ErrBookNotFound
}
