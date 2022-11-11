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

type genderAcceptanceRepo struct {
	client postgres.ClientReader
}

func NewGenderAcceptanceRepo(client postgres.ClientReader) GenderAcceptance {
	return &genderAcceptanceRepo{client: client}
}

var GenderAcceptaceModule = fx.Options(
	fx.Provide(NewGenderAcceptanceRepo),
)

func buildQueryGenderAcceptanceByNameQuery() string {
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
		`, repos.SchemaName, repos.GenderAcceptancesTableName)
}

func (r *genderAcceptanceRepo) QueryGenderAcceptanceByNameTx(ctx context.Context, client postgres.Reader, name string) (entities.GenderAcceptance, error) {
	query := buildQueryGenderAcceptanceByNameQuery()

	return r.queryGenderAcceptance(ctx, client, query, name)
}

func (r *genderAcceptanceRepo) QueryGenderAcceptanceByName(ctx context.Context, name string) (entities.GenderAcceptance, error) {
	query := buildQueryGenderAcceptanceByNameQuery()

	return r.queryGenderAcceptance(ctx, r.client, query, name)
}

func buildQueryGenderAcceptancesQuery(origin entities.Origin, paginate storage.Pagination) (string, []any) {
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
	`, repos.SchemaName, repos.GenderAcceptancesTableName, strings.Join(statements, " "))

	return query, args
}

func (r *genderAcceptanceRepo) QueryGenderAcceptances(ctx context.Context, origin entities.Origin, paginate storage.Pagination) ([]entities.GenderAcceptance, error) {
	q, args := buildQueryGenderAcceptancesQuery(origin, paginate)

	conn, err := r.client.AcquireConn(ctx)
	if err != nil {
		return nil, wrapped_error.NewInternalServerError(err, "can not acquire conn for quering gender_acceptances")
	}

	rows, err := r.client.QueryRows(ctx, conn, q, args...)
	if err != nil {
		return nil, wrapped_error.NewInternalServerError(err, "can not query gender_acceptances")
	}

	var out []entities.GenderAcceptance
	for rows.Next() {
		var ga entities.GenderAcceptance
		if err = rows.Scan(
			&ga.ID,
			&ga.Name,
			&ga.Origin,
			&ga.CreatorUserID,
			&ga.CreatedAt,
			&ga.ModifiedAt,
		); err != nil {
			return nil, wrapped_error.NewBadRequestError(err, "can not scan gender_acceptances")
		}
		out = append(out, ga)
	}
	if err = rows.Err(); err != nil {
		return nil, wrapped_error.NewBadRequestError(err, "scanning gender_acceptances failed")
	}

	return out, nil
}

func (r *genderAcceptanceRepo) CountGenderAcceptances(ctx context.Context, origin entities.Origin) (int, error) {
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
	`, repos.SchemaName, repos.GenderAcceptancesTableName, whereStatement)

	return r.countGenderAcceptances(ctx, r.client, query, args...)
}

func (r *genderAcceptanceRepo) countGenderAcceptances(ctx context.Context, client postgres.Reader, q string, args ...any) (int, error) {
	return client.CountRows(ctx, q, args...)
}

func (r *genderAcceptanceRepo) queryGenderAcceptance(ctx context.Context, client postgres.Reader, q string, args ...any) (entities.GenderAcceptance, error) {
	row, err := client.QueryRow(ctx, q, args...)
	if err != nil {
		return entities.GenderAcceptance{}, wrapped_error.NewInternalServerError(err, "can not query gender_acceptance")
	}

	var out entities.GenderAcceptance
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
		return entities.GenderAcceptance{}, nil
	}

	return entities.GenderAcceptance{}, wrapped_error.NewBadRequestError(err, "can not scan gender_acceptance")
}
