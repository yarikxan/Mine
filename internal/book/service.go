// service.go
package book

import (
	"database/sql"
	"errors"
	"minecraft/internal/common"

	"github.com/google/uuid"
)

var ErrBookNotFound = errors.New("book not found")

type Service struct {
	repo *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{
		repo: repository,
	}
}

func (s *Service) Create(book BookCreateRequestDto) (*bookCreateResponseDto, error) {
	dbBook := &BookDB{
		Title:  book.Title,
		Author: book.Author,
		ISBN:   sql.NullString{String: book.Isbn, Valid: book.Isbn != ""},
	}

	createdBook, err := s.repo.Create(dbBook)

	if err != nil {
		return nil, err
	}

	return &bookCreateResponseDto{
		Id:        createdBook.ID,
		Title:     createdBook.Title,
		Author:    createdBook.Title,
		Isbn:      createdBook.ISBN.String,
		CreatedAt: createdBook.CreatedAt,
		UpdatedAt: createdBook.UpdatedAt,
	}, nil
}

func (s *Service) GetAll(params BookListRequest) (*BookListResponse, error) {
	books, err := s.repo.GetAll(*params.Search, *params.Offset, *params.Limit)
	if err != nil {
		return nil, err
	}

	result := make([]bookCreateResponseDto, 0, len(books))
	for _, book := range books {
		result = append(result, bookCreateResponseDto{
			Id:        book.ID,
			Title:     book.Title,
			Author:    book.Author,
			Isbn:      book.ISBN.String,
			CreatedAt: book.CreatedAt,
			UpdatedAt: book.UpdatedAt,
		})
	}

	return &BookListResponse{
		pagination: common.BasePaginationDto{},
		items:      result,
	}, err
}

func (s *Service) GetByID(id uuid.UUID) (*Book, error) {
	book, err := s.repo.GetById(id)
	if err != nil {
		return nil, err
	}

	return &Book{
		ID:        book.ID,
		Title:     book.Title,
		Author:    book.Author,
		ISBN:      book.ISBN.String,
		CreatedAt: book.CreatedAt,
		UpdatedAt: book.UpdatedAt,
	}, ErrBookNotFound
}

func (s *Service) Update(id uuid.UUID, updatedBook BookUpdateRequest) (*Book, error) {
	dbBook := &BookDB{
		Title:  *(updatedBook.Title),
		Author: *(updatedBook.Author),
		ISBN:   sql.NullString{String: *(updatedBook.Isbn), Valid: *(updatedBook.Isbn) != ""},
	}

	result, err := s.repo.Update(id, dbBook)

	if err != nil {
		return nil, err
	}

	return &Book{
		ID:        result.ID,
		Title:     result.Title,
		Author:    result.Author,
		ISBN:      result.ISBN.String,
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
	}, ErrBookNotFound
}

func (s *Service) Delete(id uuid.UUID) error {
	err := s.repo.Delete(id)

	if err != nil {
		return ErrBookNotFound
	}
	return nil
}
