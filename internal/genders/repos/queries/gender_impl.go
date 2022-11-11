package queries

import (
	"context"
	"fmt"
	"strings"
	"world_generator_processor_service/core/storage"
	"world_generator_processor_service/internal/genders/entities"
	"world_generator_processor_service/internal/genders/repos"

	"github.com/olehmushka/golang-toolkit/wrapped_error"
	"github.com/olehmushka/postgres-golang"
	"go.uber.org/fx"
)

type genderRepo struct {
	client postgres.ClientReader
}

func NewGenderRepo(client postgres.ClientReader) Gender {
	return &genderRepo{client: client}
}

var GenderModule = fx.Options(
	fx.Provide(NewGenderRepo),
)

func buildQueryGenderByNameQuery() string {
	return fmt.Sprintf(`
			SELECT
			id,
			name,
			origin,
			creator_user_id,
			created_at,
			modified_at
		FROM %s.%s
		WHERE name=$1
		`, repos.SchemaName, repos.GendersTableName)
}

func (r *genderRepo) QueryGenderByNameTx(ctx context.Context, client postgres.Reader, name string) (entities.Gender, error) {
	query := buildQueryGenderByNameQuery()

	return r.queryGender(ctx, client, query, name)
}

func (r *genderRepo) QueryGenderByName(ctx context.Context, name string) (entities.Gender, error) {
	query := buildQueryGenderByNameQuery()

	return r.queryGender(ctx, r.client, query, name)
}

func buildQueryGendersQuery(origin entities.Origin, paginate storage.Pagination) (string, []any) {
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
			name,
			origin,
			creator_user_id,
			created_at,
			modified_at
		FROM %s.%s
		%s
	`, repos.SchemaName, repos.GendersTableName, strings.Join(statements, " "))

	return query, args
}

func (r *genderRepo) QueryGenders(ctx context.Context, origin entities.Origin, paginate storage.Pagination) ([]entities.Gender, error) {
	q, args := buildQueryGendersQuery(origin, paginate)

	conn, err := r.client.AcquireConn(ctx)
	if err != nil {
		return nil, wrapped_error.NewInternalServerError(err, "can not acquire conn for quering genders")
	}

	rows, err := r.client.QueryRows(ctx, conn, q, args...)
	if err != nil {
		return nil, wrapped_error.NewInternalServerError(err, "can not query genders")
	}

	var out []entities.Gender
	for rows.Next() {
		var ga entities.Gender
		if err = rows.Scan(
			&ga.ID,
			&ga.Name,
			&ga.Origin,
			&ga.CreatorUserID,
			&ga.CreatedAt,
			&ga.ModifiedAt,
		); err != nil {
			return nil, wrapped_error.NewBadRequestError(err, "can not scan genders")
		}
		out = append(out, ga)
	}
	if err = rows.Err(); err != nil {
		return nil, wrapped_error.NewBadRequestError(err, "scanning genders failed")
	}

	return out, nil
}

func (r *genderRepo) CountGenders(ctx context.Context, origin entities.Origin) (int, error) {
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
	`, repos.SchemaName, repos.GendersTableName, whereStatement)

	return r.countGenders(ctx, r.client, query, args...)
}

func (r *genderRepo) countGenders(ctx context.Context, client postgres.Reader, q string, args ...any) (int, error) {
	return client.CountRows(ctx, q, args...)
}

func (r *genderRepo) queryGender(ctx context.Context, client postgres.Reader, q string, args ...any) (entities.Gender, error) {
	row, err := client.QueryRow(ctx, q, args...)
	if err != nil {
		return entities.Gender{}, wrapped_error.NewInternalServerError(err, "can not query gender")
	}

	var out entities.Gender
	if err = row.Scan(
		&out.ID,
		&out.Name,
		&out.Origin,
		&out.CreatorUserID,
		&out.CreatedAt,
		&out.ModifiedAt,
	); err == nil {
		return out, nil
	}
	if wrapped_error.Equals(err, postgres.ErrNoRows) {
		return entities.Gender{}, nil
	}

	return entities.Gender{}, wrapped_error.NewBadRequestError(err, "can not scan gender")
}
