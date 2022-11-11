package entities

import (
	"time"

	"github.com/google/uuid"
)

type Tradition struct {
	ID                  uuid.UUID `json:"id"`
	Slug                string    `json:"slug"`
	Description         string    `json:"description"`
	PreferredEthosSlugs []string  `json:"preferred_ethos_slugs"`
	Type                Type      `json:"type"`
	OmitTraditionSlugs  []string  `json:"omit_tradition_slugs"`
	OmitGenderDominance []string  `json:"omit_gender_dominance"`
	OmitEthosSlugs      []string  `json:"omit_ethos_slugs"`
	Origin              Origin    `json:"origin"`
	CreatorUserID       uuid.UUID `json:"creator_user_id"`
	CreatedAt           time.Time `json:"created_at"`
	ModifiedAt          time.Time `json:"modified_at"`
}
