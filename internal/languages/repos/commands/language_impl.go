package commands

import (
	"context"
	"fmt"
	"strings"
	"world_generator_processor_service/internal/languages/entities"
	"world_generator_processor_service/internal/languages/repos"

	"github.com/google/uuid"
	"github.com/olehmushka/golang-toolkit/wrapped_error"
	"github.com/olehmushka/postgres-golang"
	"go.uber.org/fx"
)

type languageRepo struct {
	client postgres.ClientWriter
}

func NewLanguageRepo(client postgres.ClientWriter) Language {
	return &languageRepo{client: client}
}

var LanguageModule = fx.Options(
	fx.Provide(NewLanguageRepo),
)

func (r *languageRepo) Create(ctx context.Context, lang *entities.Language) error {
	return r.create(ctx, r.client, lang)
}

func (r *languageRepo) CreateTx(ctx context.Context, client postgres.Writer, lang *entities.Language) error {
	return r.create(ctx, client, lang)
}

func (r *languageRepo) create(ctx context.Context, client postgres.Writer, lang *entities.Language) error {
	fieldsNames := []string{
		"slug",
		"family_slug",
		"subfamily_slug",
		"wordbase_slug",
		"female_own_names",
		"male_own_names",
		"words",
		"min",
		"max",
		"dupl",
		"m",
		"origin",
	}
	args := []any{
		lang.Slug,
		lang.FamilySlug,
		lang.SubfamilySlug,
		lang.WordbaseSlug,
		lang.FemaleOwnNames,
		lang.MaleOwnNames,
		lang.Words,
		lang.Min,
		lang.Max,
		lang.Dupl,
		lang.M,
		lang.Origin,
	}
	if lang.CreatorUserID != uuid.Nil {
		fieldsNames = append(fieldsNames, "creator_user_id")
		args = append(args, lang.CreatorUserID)
	}
	values := make([]string, 0, len(fieldsNames))
	for i := range fieldsNames {
		values = append(values, fmt.Sprintf("$%d", i+1))
	}
	query := fmt.Sprintf(`
		INSERT INTO %s.%s (%s)
		VALUES (%s)
		RETURNING id
		`,
		repos.SchemaName,
		repos.FamiliesTableName,
		strings.Join(fieldsNames, ", "),
		strings.Join(values, ", "),
	)
	id, err := client.InsertRow(
		ctx,
		query,
		args...,
	)
	if err != nil {
		fmt.Printf("ErrorError: %+v\nLang: %+v\n\n", err, lang)
		return wrapped_error.NewInternalServerError(err, "can not create language")
	}
	lang.ID = id

	return nil
}
