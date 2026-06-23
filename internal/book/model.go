// model.go
package book

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Book struct {
	ID        uuid.UUID `json:"id"`
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	ISBN      string    `json:"isbn"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type BookDB struct {
	ID        uuid.UUID      `db:"id"`
	Title     string         `db:"title"`
	Author    string         `db:"author"`
	ISBN      sql.NullString `db:"isbn"`
	CreatedAt time.Time      `db:"created_at"`
	UpdatedAt time.Time      `db:"updated_at"`
}
