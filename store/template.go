package store

import "context"

type Template struct {
	ID int `json:"id"`

	// Domain specific fields
	Subject string `json:"subject"`
	Body    string `json:"body"`

	// Standard fields
	CreatedTs int64 `json:"createdTs"`
	UpdatedTs int64 `json:"updatedTs"`
}

const createTemplate = `
	INSERT INTO template (subject, body)
	VALUES ($1, $2)
	RETURNING id, subject, body, created_ts, updated_ts
`

type CreateTemplateParams struct {
	Subject string
	Body    string
}

func (s *Store) CreateTemplate(ctx context.Context, arg CreateTemplateParams) (Template, error) {
	row := s.db.QueryRowContext(ctx, createTemplate, arg.Subject, arg.Body)
	var template Template
	err := row.Scan(&template.ID, &template.Subject, &template.Body, &template.CreatedTs, &template.UpdatedTs)
	return template, err
}

const getTemplate = `
	SELECT id, subject, body, created_ts, updated_ts
	FROM template
	WHERE id = $1 LIMIT 1
`

func (s *Store) GetTemplate(ctx context.Context, id int) (Template, error) {
	row := s.db.QueryRowContext(ctx, getTemplate, id)
	var template Template
	err := row.Scan(&template.ID, &template.Subject, &template.Body, &template.CreatedTs, &template.UpdatedTs)
	return template, err
}

const listAllTemplates = `
	SELECT id, subject, body, created_ts, updated_ts
	FROM template
	ORDER BY id
`

func (s *Store) ListAllTemplates(ctx context.Context) ([]Template, error) {
	rows, err := s.db.QueryContext(ctx, listAllTemplates)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []Template
	for rows.Next() {
		var template Template
		if err := rows.Scan(&template.ID, &template.Subject, &template.Body, &template.CreatedTs, &template.UpdatedTs); err != nil {
			return nil, err
		}
		items = append(items, template)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}

const updateTemplate = `
	UPDATE template
	SET
		subject = COALESCE($1, subject),
		body = COALESCE($2, body)
	WHERE
		id = $3
	RETURNING id, subject, body, created_ts, updated_ts
`

type UpdateTemplateParams struct {
	ID int `json:"id"`

	// Domain specific fields
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

func (s *Store) UpdateTemplate(ctx context.Context, arg UpdateTemplateParams) (Template, error) {
	row := s.db.QueryRowContext(ctx, updateTemplate, arg.Subject, arg.Body, arg.ID)
	var template Template
	err := row.Scan(&template.ID, &template.Subject, &template.Body, &template.CreatedTs, &template.UpdatedTs)
	return template, err
}

const deleteTemplate = `
	DELETE FROM template
	WHERE id = ?
`

func (s *Store) DeleteTemplate(ctx context.Context, id int) error {
	_, err := s.db.ExecContext(ctx, deleteTemplate, id)
	return err
}
