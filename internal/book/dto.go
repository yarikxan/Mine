// dto.go
package book

import (
	"minecraft/internal/common"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

var validate = validator.New()

type BookCreateRequestDto struct {
	Isbn   string `json:"isbn" validate:"required,min=1,max=13"`
	Title  string `json:"title" validate:"required,min=1,max=255"`
	Author string `json:"author" validate:"required,min=1,max=100"`
}

func (b *BookCreateRequestDto) Validate() error {
	return validate.Struct(b)
}

type bookCreateResponseDto struct {
	Id        uuid.UUID
	Isbn      string
	Title     string
	Author    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type BookListRequest struct {
	Offset *int    `json:"offset,omitempty" validate:"min=0"`
	Limit  *int    `json:"limit,omitempty" validate:"min=0"`
	Search *string `json:"search,omitempty" validate:"min=0"`
}

func (b *BookListRequest) Validate() error {
	return validate.Struct(b)
}

type BookListResponse struct {
	pagination common.BasePaginationDto
	items      []bookCreateResponseDto
}

type BookRequest struct {
	common.BaseItemRequestDto
	Isbn *string `json:"isbn,omitempty" validate:"max=13"`
}

func (b *BookRequest) Validate() error {
	return validate.Struct(b)
}

type BookResponse bookCreateResponseDto

type BookUpdateRequest struct {
	Isbn   *string `json:"isbn,omitempty" validate:"min=1,max=13"`
	Title  *string `json:"title,omitempty" validate:"min=1,max=10"`
	Author *string `json:"author,omitempty" validate:"min=1,max=13"`
}

func (b *BookUpdateRequest) Validate() error {
	return validate.Struct(b)
}

type BookUpdateResponse bookCreateResponseDto

type BookDeleteRequestDto BookRequest

func (b *BookDeleteRequestDto) Validate() error {
	return validate.Struct(b)
}
