package entities

import (
	"time"

	genderEntities "world_generator_processor_service/internal/genders/entities"

	"github.com/google/uuid"
)

type Culture struct {
	ID                       uuid.UUID                       `json:"id"`
	Slug                     string                          `json:"slug"`
	BaseSlug                 string                          `json:"base_slug"`
	SubbaseSlug              string                          `json:"subbase_slug"`
	EthosSlug                string                          `json:"ethos_slug"`
	LangSlug                 string                          `json:"lang_slug"`
	DominatedGenderName      genderEntities.Gender           `json:"dominated_gender_name"`
	DominatedGenderInfluence genderEntities.Influence        `json:"dominated_gender_influence"`
	MartialCustom            genderEntities.GenderAcceptance `json:"martial_custom"`
	Origin                   Origin                          `json:"origin"`
	CreatorUserID            uuid.UUID                       `json:"creator_user_id"`
	CreatedAt                time.Time                       `json:"created_at"`
	ModifiedAt               time.Time                       `json:"modified_at"`
}
