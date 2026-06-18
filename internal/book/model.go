// model.go
package book

import "github.com/google/uuid"

type Book struct {
	ID     uuid.UUID `json:"id"`
	Title  string    `json:"title"`
	Author string    `json:"author"`
	ISBN   string    `json:"isbn"`
}
