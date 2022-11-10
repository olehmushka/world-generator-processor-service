package entities

import (
	"time"

	"github.com/google/uuid"
)

type GenderAcceptance struct {
	ID            uuid.UUID `json:"id"`
	Name          string    `json:"name"`
	Origin        Origin    `json:"origin"`
	CreatorUserID uuid.UUID `json:"creator_user_id"`
	CreatedAt     time.Time `json:"created_at"`
	ModifiedAt    time.Time `json:"modified_at"`
}
