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

type subfamilyRepo struct {
	client postgres.ClientReader
}

func NewSubfamilyRepo(client postgres.ClientReader) Subfamily {
	return &subfamilyRepo{client: client}
}

var SubfamilyModule = fx.Options(
	fx.Provide(NewSubfamilyRepo),
)

func buildQuerySubfamilyBySlugQuery() string {
	return fmt.Sprintf(`
		SELECT
			id,
			slug,
			family_slug,
			origin,
			extended_subfamily,
			creator_user_id,
			created_at,
			modified_at
		FROM %s.%s
		WHERE slug=$1
		`, repos.SchemaName, repos.SubfamiliesTableName)
}

func (r *subfamilyRepo) QuerySubfamilyBySlug(ctx context.Context, slug string) (*entities.Subfamily, error) {
	query := buildQuerySubfamilyBySlugQuery()

	return r.querySubfamily(ctx, r.client, query, slug)
}

func (r *subfamilyRepo) QuerySubfamilyBySlugTx(ctx context.Context, client postgres.Reader, slug string) (*entities.Subfamily, error) {
	query := buildQuerySubfamilyBySlugQuery()

	return r.querySubfamily(ctx, client, query, slug)
}

func buildQuerySubfamiliesQuery(origin entities.Origin, paginate storage.Pagination) (string, []any) {
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
			family_slug,
			origin,
			extended_subfamily,
			creator_user_id,
			created_at,
			modified_at
		FROM %s.%s
		%s
	`, repos.SchemaName, repos.SubfamiliesTableName, strings.Join(statements, " "))

	return query, args
}

func (r *subfamilyRepo) QuerySubfamilies(ctx context.Context, origin entities.Origin, paginate storage.Pagination) ([]*entities.Subfamily, error) {
	q, args := buildQuerySubfamiliesQuery(origin, paginate)

	conn, err := r.client.AcquireConn(ctx)
	if err != nil {
		return nil, wrapped_error.NewInternalServerError(err, "can not acquire conn for quering subfamilies")
	}

	rows, err := r.client.QueryRows(ctx, conn, q, args...)
	if err != nil {
		return nil, wrapped_error.NewInternalServerError(err, "can not query subfamilies")
	}

	var out []*entities.Subfamily
	for rows.Next() {
		var sf entities.Subfamily
		if err = rows.Scan(
			&sf.ID,
			&sf.Slug,
			&sf.FamilySlug,
			&sf.Origin,
			&sf.ExtendedSubfamily,
			&sf.CreatorUserID,
			&sf.CreatedAt,
			&sf.ModifiedAt,
		); err != nil {
			return nil, wrapped_error.NewBadRequestError(err, "can not scan subfamilies")
		}
		out = append(out, &sf)
	}
	if err = rows.Err(); err != nil {
		return nil, wrapped_error.NewBadRequestError(err, "scanning subfamilies failed")
	}

	return out, nil
}

func (r *subfamilyRepo) CountSubfamilies(ctx context.Context, origin entities.Origin) (int, error) {
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
	`, repos.SchemaName, repos.SubfamiliesTableName, whereStatement)

	return r.countSubfamilies(ctx, r.client, query, args...)
}

func (r *subfamilyRepo) countSubfamilies(ctx context.Context, client postgres.Reader, q string, args ...any) (int, error) {
	return client.CountRows(ctx, q, args...)
}

func (r *subfamilyRepo) querySubfamily(ctx context.Context, client postgres.Reader, q string, args ...any) (*entities.Subfamily, error) {
	row, err := client.QueryRow(ctx, q, args...)
	if err != nil {
		return nil, wrapped_error.NewInternalServerError(err, "can not query subfamily")
	}

	var out entities.Subfamily
	if err = row.Scan(
		&out.ID,
		&out.Slug,
		&out.FamilySlug,
		&out.Origin,
		&out.ExtendedSubfamily,
		&out.CreatorUserID,
		&out.CreatedAt,
		&out.ModifiedAt,
	); err == nil {
		return &out, nil
	}
	if wrapped_error.Equals(err, postgres.ErrNoRows) {
		return nil, nil
	}

	return nil, wrapped_error.NewBadRequestError(err, "can not scan subfamily")
}
