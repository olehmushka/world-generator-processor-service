package entities

import (
	"time"

	"github.com/google/uuid"
)

type Language struct {
	ID             uuid.UUID `json:"id"`
	Slug           string    `json:"slug"`
	FamilySlug     string    `json:"family_slug"`
	SubfamilySlug  string    `json:"subfamily_slug"`
	WordbaseSlug   string    `json:"wordbase_slug"`
	FemaleOwnNames []string  `json:"female_own_names"`
	MaleOwnNames   []string  `json:"male_own_names"`
	Words          []string  `json:"words"`
	Min            int       `json:"min"`
	Max            int       `json:"max"`
	Dupl           string    `json:"dupl"`
	M              float64   `json:"m"`
	Origin         Origin    `json:"origin"`
	CreatorUserID  uuid.UUID `json:"creator_user_id"`
	CreatedAt      time.Time `json:"created_at"`
	ModifiedAt     time.Time `json:"modified_at"`
}
