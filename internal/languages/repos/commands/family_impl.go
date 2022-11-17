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

type familyRepo struct {
	client postgres.ClientWriter
}

func NewFamilyRepo(client postgres.ClientWriter) Family {
	return &familyRepo{client: client}
}

var FamilyModule = fx.Options(
	fx.Provide(NewFamilyRepo),
)

func (r *familyRepo) Create(ctx context.Context, f *entities.Family) error {
	return r.create(ctx, r.client, f)
}

func (r *familyRepo) CreateTx(ctx context.Context, client postgres.Writer, f *entities.Family) error {
	return r.create(ctx, client, f)
}

func (r *familyRepo) create(ctx context.Context, client postgres.Writer, f *entities.Family) error {
	fieldsNames := []string{"slug", "origin"}
	args := []any{f.Slug, f.Origin}
	if f.CreatorUserID != uuid.Nil {
		fieldsNames = append(fieldsNames, "creator_user_id")
		args = append(args, f.CreatorUserID)
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
		return wrapped_error.NewInternalServerError(err, "can not create family")
	}
	f.ID = id

	return nil
}
