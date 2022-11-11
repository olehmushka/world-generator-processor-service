package entities

import (
	"time"

	"github.com/google/uuid"
)

type Subfamily struct {
	ID                uuid.UUID     `json:"id"`
	Slug              string        `json:"slug"`
	FamilySlug        string        `json:"family_slug"`
	Origin            Origin        `json:"origin"`
	ExtendedSubfamily *RawSubfamily `json:"extended_subfamily"`
	CreatorUserID     uuid.UUID     `json:"creator_user_id"`
	CreatedAt         time.Time     `json:"created_at"`
	ModifiedAt        time.Time     `json:"modified_at"`
}

type RawSubfamily struct {
	Slug              string        `json:"slug"`
	FamilySlug        string        `json:"family_slug"`
	ExtendedSubfamily *RawSubfamily `json:"extended_subfamily"`
}
