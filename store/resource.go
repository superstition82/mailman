package store

import (
	"context"
	"errors"
	"fmt"
	"strings"
)

type Resource struct {
	ID int `json:"id"`

	// Standard fields
	CreatedTs int64 `json:"createdTs"`
	UpdatedTs int64 `json:"updatedTs"`

	// Domain specific fields
	Filename     string `json:"filename"`
	Blob         []byte `json:"-"`
	ExternalLink string `json:"externalLink"`
	InternalPath string `json:"internalPath"`
	Type         string `json:"type"`
	Size         int64  `json:"size"`

	// Related fields
	LinkedTemplateAmount int `json:"linkedTemplateAmount"`
}

const createResource = `
	INSERT INTO resource (filename, blob, type, size, internal_path, external_link)
	VALUES ($1, $2, $3, $4, $5, $6)
	RETURNING id, filename, blob, type, size, internal_path, external_link, created_ts, updated_ts
`

type CreateResourceParams struct {
	Filename     string
	Blob         []byte
	InternalPath string
	ExternalLink string
	Type         string
	Size         int64
}

func (s *Store) CreateResource(ctx context.Context, create CreateResourceParams) (*Resource, error) {
	row := s.db.QueryRowContext(
		ctx,
		createResource,
		create.Filename,
		create.Blob,
		create.Type,
		create.Size,
		create.InternalPath,
		create.ExternalLink,
	)
	var resource Resource
	err := row.Scan(
		&resource.ID,
		&resource.Filename,
		&resource.Blob,
		&resource.Type,
		&resource.Size,
		&resource.InternalPath,
		&resource.ExternalLink,
		&resource.CreatedTs,
		&resource.UpdatedTs,
	)
	return &resource, err
}

type ResourcePatch struct {
	ID int `json:"-"`

	// Standard fields
	UpdatedTs *int64

	// Domain specific fields
	Filename *string `json:"filename"`
}

func (s *Store) PatchResource(ctx context.Context, patch ResourcePatch) (*Resource, error) {
	set, args := []string{}, []any{}

	if v := patch.UpdatedTs; v != nil {
		set, args = append(set, "updated_ts = ?"), append(args, *v)
	}
	if v := patch.Filename; v != nil {
		set, args = append(set, "filename = ?"), append(args, *v)
	}

	args = append(args, patch.ID)
	fields := []string{"id", "filename", "type", "size", "created_ts", "updated_ts", "internal_path", "external_link"}
	query := `
		UPDATE resource
		SET ` + strings.Join(set, ", ") + `
		WHERE id = ?
		RETURNING ` + strings.Join(fields, ", ")
	var resource Resource
	dests := []any{
		&resource.ID,
		&resource.Filename,
		&resource.Type,
		&resource.Size,
		&resource.CreatedTs,
		&resource.UpdatedTs,
		&resource.InternalPath,
		&resource.ExternalLink,
	}

	if err := s.db.QueryRowContext(ctx, query, args...).Scan(dests...); err != nil {
		return nil, err
	}

	return &resource, nil
}

func (s *Store) FindResource(ctx context.Context, id int) (*Resource, error) {
	fields := []string{"id", "filename", "type", "size", "created_ts", "updated_ts", "internal_path", "external_link"}
	query := fmt.Sprintf(`
	SELECT %s FROM resource
	WHERE resource.id = %d
`, strings.Join(fields, ", "), id)

	row := s.db.QueryRowContext(ctx, query, id)
	var resource Resource
	err := row.Scan(
		&resource.ID,
		&resource.Filename,
		&resource.Type,
		&resource.Size,
		&resource.CreatedTs,
		&resource.UpdatedTs,
		&resource.InternalPath,
		&resource.ExternalLink,
	)

	return &resource, err
}

type ResourceFind struct {
	ID *int `json:"id"`

	// Domain specific fields
	Filename *string `json:"filename"`
	MemoID   *int
	GetBlob  bool

	// Pagination
	Limit  *int
	Offset *int
}

func (s *Store) FindResourceList(ctx context.Context, find *ResourceFind) ([]*Resource, error) {
	where, args := []string{"1 = 1"}, []any{}
	fields := []string{"resource.id", "resource.filename", "resource.type", "resource.size", "resource.created_ts", "resource.updated_ts", "internal_path", "external_link"}

	if v := find.ID; v != nil {
		where, args = append(where, "resource.id = ?"), append(args, *v)
	}
	if v := find.Filename; v != nil {
		where, args = append(where, "resource.filename = ?"), append(args, *v)
	}
	if v := find.MemoID; v != nil {
		where, args = append(where, "resource.id in (SELECT resource_id FROM memo_resource WHERE memo_id = ?)"), append(args, *v)
	}

	query := fmt.Sprintf(`
		SELECT
			COUNT(DISTINCT template_resource.template_id) AS linked_template_amount,
			%s
		FROM resource
		LEFT JOIN template_resource ON resource.id = template_resource.resource_id
		WHERE %s
		GROUP BY resource.id
		ORDER BY resource.id DESC
	`, strings.Join(fields, ", "), strings.Join(where, " AND "))
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

	resourceList := make([]*Resource, 0)
	for rows.Next() {
		var resourceRaw Resource
		dests := []any{
			&resourceRaw.LinkedTemplateAmount,
			&resourceRaw.ID,
			&resourceRaw.Filename,
			&resourceRaw.Type,
			&resourceRaw.Size,
			&resourceRaw.CreatedTs,
			&resourceRaw.UpdatedTs,
			&resourceRaw.InternalPath,
			&resourceRaw.ExternalLink,
		}
		if find.GetBlob {
			dests = append(dests, &resourceRaw.Blob)
		}
		if err := rows.Scan(dests...); err != nil {
			return nil, err
		}
		resourceList = append(resourceList, &resourceRaw)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return resourceList, nil
}

type ResourceDelete struct {
	ID int
}

func (s *Store) DeleteResource(ctx context.Context, delete *ResourceDelete) error {
	where, args := []string{"id = ?"}, []any{delete.ID}

	stmt := `DELETE FROM resource WHERE ` + strings.Join(where, " AND ")
	result, err := s.db.ExecContext(ctx, stmt, args...)
	if err != nil {
		return err
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return errors.New("resource not found")
	}

	return nil
}
