// repository.go
package book

import (
	"database/sql"

	"github.com/google/uuid"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(book *BookDB) (BookDB, error) {
	return BookDB{}, nil
}
func (r *Repository) GetAll(search string, offset, limit int) ([]BookDB, error) {
	return nil, nil
}
func (r *Repository) GetById(id uuid.UUID) (BookDB, error)              { return BookDB{}, nil }
func (r *Repository) Update(id uuid.UUID, book *BookDB) (BookDB, error) { return BookDB{}, nil }
func (r *Repository) Delete(id uuid.UUID) error                         { return nil }
