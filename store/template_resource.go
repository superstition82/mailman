package store

import (
	"context"
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
