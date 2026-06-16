package baseItemDto

import (
	"time"

	"github.com/google/uuid"
)

type BaseItemResponseDto struct {
	Id        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type BaseItemRequestDto struct {
	Id uuid.UUID `json:"id,omitempty" validate:"uuid"`
}

type BasePaginationDto struct {
	Offset int `json:"offset,omitempty"`
	Limit  int `json:"limit,omitempty"`
}
