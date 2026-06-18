// service.go
package book

import (
	"errors"

	"github.com/google/uuid"
)

var ErrBookNotFound = errors.New("book not found")

type Service struct {
	books []Book
	repo  *Repository
}

func NewService() *Service {
	return &Service{
		books: []Book{},
	}
}

func (s *Service) Create(book Book) (Book, error) {
	s.books = append(s.books, book)
	return book, nil
}

func (s *Service) GetAll() ([]Book, error) {
	return s.books, nil
}

func (s *Service) GetByID(id uuid.UUID) (Book, error) {
	for _, book := range s.books {
		if book.ID == id {
			return book, nil
		}
	}
	return Book{}, ErrBookNotFound
}

func (s *Service) Update(id uuid.UUID, updatedBook Book) (Book, error) {
	for idx, book := range s.books {
		if book.ID == id {
			s.books[idx] = updatedBook
			return updatedBook, nil
		}
	}
	return Book{}, ErrBookNotFound
}

func (s *Service) Delete(id uuid.UUID) error {
	for i, book := range s.books {
		if book.ID == id {
			s.books = append(s.books[:i], s.books[i+1:]...)
			return nil
		}
	}
	return ErrBookNotFound
}
