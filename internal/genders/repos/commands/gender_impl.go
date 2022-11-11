package commands

import (
	"context"
	"fmt"
	"strings"
	"world_generator_processor_service/internal/genders/entities"
	"world_generator_processor_service/internal/genders/repos"

	"github.com/google/uuid"
	"github.com/olehmushka/golang-toolkit/wrapped_error"
	"github.com/olehmushka/postgres-golang"
	"go.uber.org/fx"
)

type genderRepo struct {
	client postgres.ClientWriter
}

func NewGenderRepo(client postgres.ClientWriter) Gender {
	return &genderRepo{client: client}
}

var GenderModule = fx.Options(
	fx.Provide(NewGenderRepo),
)

func (r *genderRepo) Create(ctx context.Context, g entities.Gender) error {
	return r.create(ctx, r.client, g)
}

func (r *genderRepo) CreateTx(ctx context.Context, client postgres.Writer, g entities.Gender) error {
	return r.create(ctx, client, g)
}

func (r *genderRepo) create(ctx context.Context, client postgres.Writer, g entities.Gender) error {
	fieldsNames := []string{"name", "origin"}
	args := []any{g.Name, g.Origin}
	if g.CreatorUserID != uuid.Nil {
		fieldsNames = append(fieldsNames, "creator_user_id")
		args = append(args, g.CreatorUserID)
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
		repos.GendersTableName,
		strings.Join(fieldsNames, ", "),
		strings.Join(values, ", "),
	)
	id, err := client.InsertRow(
		ctx,
		query,
		args...,
	)
	if err != nil {
		return wrapped_error.NewInternalServerError(err, "can not create gender")
	}
	g.ID = id

	return nil
}
