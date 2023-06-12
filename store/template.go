package store

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Template struct {
	ID int `json:"id"`

	// Standard fields
	CreatedTs int64 `json:"createdTs"`
	UpdatedTs int64 `json:"updatedTs"`

	// Domain specific fields
	Subject string `json:"subject"`
	Body    string `json:"body"`

	// Composed fields
	ResourceIDList []int `json:"resourceIdList"`
}

type TemplateCreate struct {
	Subject string
	Body    string
}

func (s *Store) CreateTemplate(ctx context.Context, create *TemplateCreate) (*Template, error) {
	query := `
		INSERT INTO template (subject, body)
		VALUES ($1, $2)
		RETURNING id, subject, body, created_ts, updated_ts
	`

	row := s.db.QueryRowContext(
		ctx,
		query,
		create.Subject,
		create.Body,
	)
	var template Template
	err := row.Scan(
		&template.ID,
		&template.Subject,
		&template.Body,
		&template.CreatedTs,
		&template.UpdatedTs,
	)
	return &template, err
}

type TemplateFind struct {
	ID *int `json:"id"`

	// Pagination
	Limit  *int
	Offset *int
}

func (s *Store) FindTemplateList(ctx context.Context, find *TemplateFind) ([]*Template, error) {
	where, args := []string{"1 = 1"}, []any{}

	if v := find.ID; v != nil {
		where, args = append(where, "template.id = ?"), append(args, *v)
	}
	query := fmt.Sprintf(`
		SELECT 
			template.id AS id,
			template.created_ts AS created_ts,
			template.updated_ts AS updated_ts,
			template.subject AS subject,
			template.body AS body,
			GROUP_CONCAT(template_resource.resource_id) AS resource_id_list
		FROM 
			template
		LEFT JOIN
			template_resource ON template.id = template_resource.template_id
		WHERE %s
		GROUP BY template.id
		ORDER BY id DESC
	`, strings.Join(where, " AND "))
	if find.Limit != nil {
		query = fmt.Sprintf("%s LIMIT %d", query, *find.Limit)
		if find.Offset != nil {
			query = fmt.Sprintf("%s OFFSET %d", query, *find.Offset)
		}
	}

	rows, err := s.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	templateList := make([]*Template, 0)
	for rows.Next() {
		var template Template
		var templateResourceIDList sql.NullString
		dests := []any{
			&template.ID,
			&template.CreatedTs,
			&template.UpdatedTs,
			&template.Subject,
			&template.Body,
			&templateResourceIDList,
		}
		if err := rows.Scan(dests...); err != nil {
			return nil, err
		}
		if templateResourceIDList.Valid {
			idStringList := strings.Split(templateResourceIDList.String, ",")
			template.ResourceIDList = make([]int, 0, len(idStringList))
			for _, idString := range idStringList {
				id, err := strconv.Atoi(idString)
				if err != nil {
					return nil, err
				}
				template.ResourceIDList = append(template.ResourceIDList, id)
			}
		}
		templateList = append(templateList, &template)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return templateList, nil
}

func (s *Store) FindTemplate(ctx context.Context, find *TemplateFind) (*Template, error) {
	list, err := s.FindTemplateList(ctx, find)
	if err != nil {
		return nil, err
	}
	if len(list) == 0 {
		return nil, errors.New("not found")
	}
	template := list[0]

	return template, nil
}

type TemplatePatch struct {
	ID int `json:"-"`

	// Standard fields
	UpdatedTs *int64

	// Domain specific fields
	Subject *string `json:"subject"`
	Body    *string `json:"body"`
}

func (s *Store) PatchTemplate(ctx context.Context, patch *TemplatePatch) (*Template, error) {
	set, args := []string{}, []any{}
	if v := patch.UpdatedTs; v != nil {
		set, args = append(set, "updated_ts = ?"), append(args, *v)
	}
	if v := patch.Subject; v != nil {
		set, args = append(set, "subject = ?"), append(args, *v)
	}
	if v := patch.Body; v != nil {
		set, args = append(set, "body = ?"), append(args, *v)
	}
	args = append(args, patch.ID)
	fields := []string{"id", "subject", "body", "created_ts", "updated_ts"}
	query := `
		UPDATE template
		SET ` + strings.Join(set, ", ") + `
		WHERE id = ?
		RETURNING ` + strings.Join(fields, ", ")
	var template Template
	dests := []any{
		&template.ID,
		&template.Subject,
		&template.Body,
		&template.CreatedTs,
		&template.UpdatedTs,
	}
	if err := s.db.QueryRowContext(ctx, query, args...).Scan(dests...); err != nil {
		return nil, err
	}

	return &template, nil
}

type TemplateDelete struct {
	ID int
}

const deleteTemplate = `
	DELETE FROM template
	WHERE id = ?
`

func (s *Store) DeleteTemplate(ctx context.Context, delete *TemplateDelete) error {
	where, args := []string{"id = ?"}, []any{delete.ID}

	stmt := `DELETE FROM template WHERE ` + strings.Join(where, " AND ")
	result, err := s.db.ExecContext(ctx, stmt, args...)
	if err != nil {
		return err
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return errors.New("template not found")
	}

	return nil
}

func (s *Store) DeleteBulkTemplate(ctx context.Context, ids []int) error {
	tx, err := s.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}

	for _, id := range ids {
		_, err = tx.Exec(deleteTemplate, id)
		if err != nil {
			tx.Rollback()
			return nil
		}
	}

	err = tx.Commit()
	if err != nil {
		return nil
	}

	return nil
}
