package entities

import (
	"time"

	"github.com/google/uuid"
)

type Subbase struct {
	ID            uuid.UUID `json:"id"`
	Slug          string    `json:"slug"`
	BaseSlug      string    `json:"base_slug"`
	Origin        Origin    `json:"origin"`
	CreatorUserID uuid.UUID `json:"creator_user_id"`
	CreatedAt     time.Time `json:"created_at"`
	ModifiedAt    time.Time `json:"modified_at"`
}
