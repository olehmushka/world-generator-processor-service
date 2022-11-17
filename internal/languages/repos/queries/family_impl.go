package queries

import (
	"context"
	"fmt"
	"strings"
	"world_generator_processor_service/core/storage"
	"world_generator_processor_service/internal/languages/entities"
	"world_generator_processor_service/internal/languages/repos"

	"github.com/olehmushka/golang-toolkit/wrapped_error"
	"github.com/olehmushka/postgres-golang"
	"go.uber.org/fx"
)

type familyRepo struct {
	client postgres.ClientReader
}

func NewFamilyRepo(client postgres.ClientReader) Family {
	return &familyRepo{client: client}
}

var FamilyModule = fx.Options(
	fx.Provide(NewFamilyRepo),
)

func buildQueryFamilyBySlugQuery() string {
	return fmt.Sprintf(`
			SELECT
			id,
			slug,
			origin,
			creator_user_id,
			created_at,
			modified_at
		FROM %s.%s
		WHERE slug=$1
		`, repos.SchemaName, repos.FamiliesTableName)
}

func (r *familyRepo) QueryFamilyBySlug(ctx context.Context, slug string) (*entities.Family, error) {
	query := buildQueryFamilyBySlugQuery()

	return r.queryFamily(ctx, r.client, query, slug)
}

func (r *familyRepo) QueryFamilyBySlugTx(ctx context.Context, client postgres.Reader, slug string) (*entities.Family, error) {
	query := buildQueryFamilyBySlugQuery()

	return r.queryFamily(ctx, client, query, slug)
}

func buildQueryFamiliesQuery(origin entities.Origin, paginate storage.Pagination) (string, []any) {
	args := make([]any, 0, 2)
	statements := make([]string, 0, 2)
	if !origin.IsZero() {
		statements = append(statements, "WHERE origin=$1")
		args = append(args, origin)
	}
	statements = append(statements, fmt.Sprintf("LIMIT $%d", len(statements)+1))
	args = append(args, paginate.Limit)
	statements = append(statements, fmt.Sprintf("OFFSET $%d", len(statements)+1))
	args = append(args, paginate.Offset)
	query := fmt.Sprintf(`
		SELECT
			id,
			slug,
			origin,
			creator_user_id,
			created_at,
			modified_at
		FROM %s.%s
		%s
	`, repos.SchemaName, repos.FamiliesTableName, strings.Join(statements, " "))

	return query, args
}

func (r *familyRepo) QueryFamilies(ctx context.Context, origin entities.Origin, paginate storage.Pagination) ([]*entities.Family, error) {
	q, args := buildQueryFamiliesQuery(origin, paginate)

	conn, err := r.client.AcquireConn(ctx)
	if err != nil {
		return nil, wrapped_error.NewInternalServerError(err, "can not acquire conn for quering families")
	}

	rows, err := r.client.QueryRows(ctx, conn, q, args...)
	if err != nil {
		return nil, wrapped_error.NewInternalServerError(err, "can not query families")
	}

	var out []*entities.Family
	for rows.Next() {
		var f entities.Family
		if err = rows.Scan(
			&f.ID,
			&f.Slug,
			&f.Origin,
			&f.CreatorUserID,
			&f.CreatedAt,
			&f.ModifiedAt,
		); err != nil {
			return nil, wrapped_error.NewBadRequestError(err, "can not scan families")
		}
		out = append(out, &f)
	}
	if err = rows.Err(); err != nil {
		return nil, wrapped_error.NewBadRequestError(err, "scanning families failed")
	}

	return out, nil
}

func (r *familyRepo) CountFamilies(ctx context.Context, origin entities.Origin) (int, error) {
	var whereStatement string
	args := make([]any, 0, 1)
	if !origin.IsZero() {
		whereStatement = "WHERE origin=$1"
		args = append(args, origin)
	}
	query := fmt.Sprintf(`
		SELECT COUNT(*)
		FROM %s.%s
		%s
	`, repos.SchemaName, repos.FamiliesTableName, whereStatement)

	return r.countFamilies(ctx, r.client, query, args...)
}

func (r *familyRepo) countFamilies(ctx context.Context, client postgres.Reader, q string, args ...any) (int, error) {
	return client.CountRows(ctx, q, args...)
}

func (r *familyRepo) queryFamily(ctx context.Context, client postgres.Reader, q string, args ...any) (*entities.Family, error) {
	row, err := client.QueryRow(ctx, q, args...)
	if err != nil {
		return nil, wrapped_error.NewInternalServerError(err, "can not query family")
	}

	var out entities.Family
	if err = row.Scan(
		&out.ID,
		&out.Slug,
		&out.Origin,
		&out.CreatorUserID,
		&out.CreatedAt,
		&out.ModifiedAt,
	); err == nil {
		return &out, nil
	}
	if wrapped_error.Equals(err, postgres.ErrNoRows) {
		return nil, nil
	}

	return nil, wrapped_error.NewBadRequestError(err, "can not scan family")
}
