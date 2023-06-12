package store

import (
	"context"
	"errors"
	"strings"
)

type TemplateResource struct {
	TemplateID int   `json:"templateId"`
	ResourceID int   `json:"resourceId"`
	CreatedTs  int64 `json:"createdTs"`
	UpdatedTs  int64 `json:"updatedTs"`
}

type TemplateResourceUpsert struct {
	TemplateID int `json:"-"`
	ResourceID int
	UpdatedTs  *int64
}

func (s *Store) UpsertTemplateResource(ctx context.Context, upsert *TemplateResourceUpsert) (*TemplateResource, error) {
	set := []string{"template_id", "resource_id"}
	args := []any{upsert.TemplateID, upsert.ResourceID}
	placeholder := []string{"?", "?"}

	if v := upsert.UpdatedTs; v != nil {
		set, args, placeholder = append(set, "updated_ts"), append(args, v), append(placeholder, "?")
	}

	query := `
		INSERT INTO template_resource (
			` + strings.Join(set, ", ") + `
		)
		VALUES (` + strings.Join(placeholder, ",") + `)
		ON CONFLICT(template_id, resource_id) DO UPDATE 
		SET
			updated_ts = EXCLUDED.updated_ts
		RETURNING template_id, resource_id, created_ts, updated_ts
	`
	var templateResource TemplateResource
	if err := s.db.QueryRowContext(ctx, query, args...).Scan(
		&templateResource.TemplateID,
		&templateResource.ResourceID,
		&templateResource.CreatedTs,
		&templateResource.UpdatedTs,
	); err != nil {
		return nil, err
	}

	return &templateResource, nil
}

type TemplateResourceDelete struct {
	TemplateID *int
	ResourceID *int
}

func (s *Store) DeleteTemplateResource(ctx context.Context, delete *TemplateResourceDelete) error {
	where, args := []string{}, []any{}

	if v := delete.TemplateID; v != nil {
		where, args = append(where, "template_id = ?"), append(args, *v)
	}
	if v := delete.ResourceID; v != nil {
		where, args = append(where, "resource_id = ?"), append(args, *v)
	}

	stmt := `DELETE FROM template_resource WHERE ` + strings.Join(where, " AND ")
	result, err := s.db.ExecContext(ctx, stmt, args...)
	if err != nil {
		return err
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return errors.New("template resource not found")
	}

	return nil
}
