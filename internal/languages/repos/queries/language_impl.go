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

type languageRepo struct {
	client postgres.ClientReader
}

func NewLanguageRepo(client postgres.ClientReader) Language {
	return &languageRepo{client: client}
}

var LanguageModule = fx.Options(
	fx.Provide(NewLanguageRepo),
)

func buildQueryLanguageBySlugQuery() string {
	return fmt.Sprintf(`
		SELECT
			id,
			slug,
			family_slug,
			subfamily_slug,
			wordbase_slug,
			female_own_names,
			male_own_names,
			words,
			min,
			max,
			dupl,
			m,
			origin,
			creator_user_id,
			created_at,
			modified_at
		FROM %s.%s
		WHERE slug=$1
		`, repos.SchemaName, repos.LanguagesTableName)
}

func (r *languageRepo) QueryLanguageBySlug(ctx context.Context, slug string) (*entities.Language, error) {
	query := buildQueryLanguageBySlugQuery()

	return r.queryLanguage(ctx, r.client, query, slug)
}

func (r *languageRepo) QueryLanguageBySlugTx(ctx context.Context, client postgres.Reader, slug string) (*entities.Language, error) {
	query := buildQueryLanguageBySlugQuery()

	return r.queryLanguage(ctx, client, query, slug)
}

func buildQueryLanguagesQuery(origin entities.Origin, paginate storage.Pagination) (string, []any) {
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
			subfamily_slug,
			wordbase_slug,
			female_own_names,
			male_own_names,
			words,
			min,
			max,
			dupl,
			m,
			origin,
			creator_user_id,
			created_at,
			modified_at
		FROM %s.%s
		%s
	`, repos.SchemaName, repos.LanguagesTableName, strings.Join(statements, " "))

	return query, args
}

func (r *languageRepo) QueryLanguages(ctx context.Context, origin entities.Origin, paginate storage.Pagination) ([]*entities.Language, error) {
	q, args := buildQueryLanguagesQuery(origin, paginate)

	conn, err := r.client.AcquireConn(ctx)
	if err != nil {
		return nil, wrapped_error.NewInternalServerError(err, "can not acquire conn for quering languages")
	}

	rows, err := r.client.QueryRows(ctx, conn, q, args...)
	if err != nil {
		return nil, wrapped_error.NewInternalServerError(err, "can not query languages")
	}

	var out []*entities.Language
	for rows.Next() {
		var lang entities.Language
		if err = rows.Scan(
			&lang.ID,
			&lang.Slug,
			&lang.FamilySlug,
			&lang.SubfamilySlug,
			&lang.WordbaseSlug,
			&lang.FemaleOwnNames,
			&lang.MaleOwnNames,
			&lang.Words,
			&lang.Min,
			&lang.Max,
			&lang.Dupl,
			&lang.M,
			&lang.Origin,
			&lang.CreatorUserID,
			&lang.CreatedAt,
			&lang.ModifiedAt,
		); err != nil {
			return nil, wrapped_error.NewBadRequestError(err, "can not scan languages")
		}
		out = append(out, &lang)
	}
	if err = rows.Err(); err != nil {
		return nil, wrapped_error.NewBadRequestError(err, "scanning languages failed")
	}

	return out, nil
}

func (r *languageRepo) CountLanguages(ctx context.Context, origin entities.Origin) (int, error) {
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
	`, repos.SchemaName, repos.LanguagesTableName, whereStatement)

	return r.countLanguages(ctx, r.client, query, args...)
}

func (r *languageRepo) countLanguages(ctx context.Context, client postgres.Reader, q string, args ...any) (int, error) {
	return client.CountRows(ctx, q, args...)
}

func (r *languageRepo) queryLanguage(ctx context.Context, client postgres.Reader, q string, args ...any) (*entities.Language, error) {
	row, err := client.QueryRow(ctx, q, args...)
	if err != nil {
		return nil, wrapped_error.NewInternalServerError(err, "can not query language")
	}

	var out entities.Language
	if err = row.Scan(
		&out.ID,
		&out.Slug,
		&out.FamilySlug,
		&out.SubfamilySlug,
		&out.WordbaseSlug,
		&out.FemaleOwnNames,
		&out.MaleOwnNames,
		&out.Words,
		&out.Min,
		&out.Max,
		&out.Dupl,
		&out.M,
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

	return nil, wrapped_error.NewBadRequestError(err, "can not scan language")
}
