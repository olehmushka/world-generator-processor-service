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

type subfamilyRepo struct {
	client postgres.ClientWriter
}

func NewSubfamilyRepo(client postgres.ClientWriter) Subfamily {
	return &subfamilyRepo{client: client}
}

var SubfamilyModule = fx.Options(
	fx.Provide(NewSubfamilyRepo),
)

func (r *subfamilyRepo) Create(ctx context.Context, sf *entities.Subfamily) error {
	return r.create(ctx, r.client, sf)
}

func (r *subfamilyRepo) CreateTx(ctx context.Context, client postgres.Writer, sf *entities.Subfamily) error {
	return r.create(ctx, client, sf)
}

func (r *subfamilyRepo) create(ctx context.Context, client postgres.Writer, sf *entities.Subfamily) error {
	fieldsNames := []string{"slug", "origin", "family_slug"}
	args := []any{sf.Slug, sf.Origin, sf.FamilySlug}
	if sf.CreatorUserID != uuid.Nil {
		fieldsNames = append(fieldsNames, "creator_user_id")
		args = append(args, sf.CreatorUserID)
	}
	if sf.ExtendedSubfamily != nil {
		fieldsNames = append(fieldsNames, "extended_subfamily")
		args = append(args, sf.ExtendedSubfamily)
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
		repos.SubfamiliesTableName,
		strings.Join(fieldsNames, ", "),
		strings.Join(values, ", "),
	)
	id, err := client.InsertRow(
		ctx,
		query,
		args...,
	)
	if err != nil {
		return wrapped_error.NewInternalServerError(err, "can not create subfamily")
	}
	sf.ID = id

	return nil
}
