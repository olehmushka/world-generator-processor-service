package engine

import (
	"testing"
	"world_generator_processor_service/internal/languages/entities"

	engineLang "github.com/olehmushka/world-generator-engine/language"
	"github.com/stretchr/testify/assert"
)

func TestDeserializeRawSubfamily(t *testing.T) {
	tCases := map[string]struct {
		input          *engineLang.Subfamily
		expectedOutput *entities.RawSubfamily
	}{
		"should return nil for nil input": {
			input:          nil,
			expectedOutput: nil,
		},
		"should deserialize a subfamily without an extended subfamily": {
			input: &engineLang.Subfamily{
				Slug:       "germanic",
				FamilySlug: "indo_european",
			},
			expectedOutput: &entities.RawSubfamily{
				Slug:       "germanic",
				FamilySlug: "indo_european",
			},
		},
		"should recursively deserialize a nested extended subfamily": {
			input: &engineLang.Subfamily{
				Slug:       "west_germanic",
				FamilySlug: "indo_european",
				ExtendedSubfamily: &engineLang.Subfamily{
					Slug:       "germanic",
					FamilySlug: "indo_european",
					ExtendedSubfamily: &engineLang.Subfamily{
						Slug:       "proto_indo_european",
						FamilySlug: "indo_european",
					},
				},
			},
			expectedOutput: &entities.RawSubfamily{
				Slug:       "west_germanic",
				FamilySlug: "indo_european",
				ExtendedSubfamily: &entities.RawSubfamily{
					Slug:       "germanic",
					FamilySlug: "indo_european",
					ExtendedSubfamily: &entities.RawSubfamily{
						Slug:       "proto_indo_european",
						FamilySlug: "indo_european",
					},
				},
			},
		},
	}

	for name, tc := range tCases {
		t.Run(name, func(tt *testing.T) {
			out := deserializeRawSubfamily(tc.input)
			assert.Equal(tt, tc.expectedOutput, out)
		})
	}
}

func TestDeserializeSubfamily(t *testing.T) {
	tCases := map[string]struct {
		input          *engineLang.Subfamily
		expectedOutput *entities.Subfamily
	}{
		"should return nil for nil input": {
			input:          nil,
			expectedOutput: nil,
		},
		"should deserialize a subfamily and set native origin": {
			input: &engineLang.Subfamily{
				Slug:       "germanic",
				FamilySlug: "indo_european",
			},
			expectedOutput: &entities.Subfamily{
				Slug:       "germanic",
				FamilySlug: "indo_european",
				Origin:     entities.NativeOrigin,
			},
		},
		"should deserialize nested extended subfamilies as raw subfamilies": {
			input: &engineLang.Subfamily{
				Slug:       "west_germanic",
				FamilySlug: "indo_european",
				ExtendedSubfamily: &engineLang.Subfamily{
					Slug:       "germanic",
					FamilySlug: "indo_european",
				},
			},
			expectedOutput: &entities.Subfamily{
				Slug:       "west_germanic",
				FamilySlug: "indo_european",
				Origin:     entities.NativeOrigin,
				ExtendedSubfamily: &entities.RawSubfamily{
					Slug:       "germanic",
					FamilySlug: "indo_european",
				},
			},
		},
	}

	for name, tc := range tCases {
		t.Run(name, func(tt *testing.T) {
			out := deserializeSubfamily(tc.input)
			assert.Equal(tt, tc.expectedOutput, out)
		})
	}
}
