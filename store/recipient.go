package store

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
)

type Recipient struct {
	ID int `json:"id"`

	// Standard fields
	CreatedTs int64 `json:"createdTs"`
	UpdatedTs int64 `json:"updatedTs"`

	// Domain specific fields
	Email     string `json:"email"`
	Reachable string `json:"reachable"`
}

type RecipientCreate struct {
	Email     string
	Reachable string
}

func (s *Store) CreateRecipient(ctx context.Context, create *RecipientCreate) (*Recipient, error) {
	query := `
		INSERT INTO recipient (email, reachable) 
		VALUES ($1, $2)
		RETURNING id, email, reachable, created_ts, updated_ts
	`
	row := s.db.QueryRowContext(
		ctx,
		query,
		create.Email,
		create.Reachable,
	)

	var recipient Recipient
	err := row.Scan(
		&recipient.ID,
		&recipient.Email,
		&recipient.Reachable,
		&recipient.CreatedTs,
		&recipient.UpdatedTs,
	)
	return &recipient, err
}

type RecipientFind struct {
	ID *int `json:"id"`

	// Pagination
	Limit  *int
	Offset *int
}

func (s *Store) FindRecipientList(ctx context.Context, find *RecipientFind) ([]*Recipient, error) {
	where, args := []string{"1 = 1"}, []any{}

	if v := find.ID; v != nil {
		where, args = append(where, "recipient.id = ?"), append(args, *v)
	}

	fields := []string{"recipient.id", "recipient.email", "recipient.reachable", "recipient.created_ts", "recipient.updated_ts"}
	query := fmt.Sprintf(`
		SELECT %s
		FROM recipient
		WHERE %s
		ORDER BY id DESC
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

	recipientList := make([]*Recipient, 0)
	for rows.Next() {
		var recipient Recipient
		dests := []any{
			&recipient.ID,
			&recipient.Email,
			&recipient.Reachable,
			&recipient.CreatedTs,
			&recipient.UpdatedTs,
		}
		if err := rows.Scan(dests...); err != nil {
			return nil, err
		}
		recipientList = append(recipientList, &recipient)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return recipientList, nil
}

func (s *Store) FindRecipient(ctx context.Context, find *RecipientFind) (*Recipient, error) {
	list, err := s.FindRecipientList(ctx, find)
	if err != nil {
		return nil, err
	}
	if len(list) == 0 {
		return nil, errors.New("not found")
	}
	recipient := list[0]

	return recipient, nil
}

type RecipientPatch struct {
	ID int `json:"-"`

	// Standard fields
	UpdatedTs *int64

	// Domain specific fields
	Email     *string `json:"filename"`
	Reachable *string `json:"reachable"`
}

func (s *Store) PatchRecipient(ctx context.Context, patch *RecipientPatch) (*Recipient, error) {
	set, args := []string{}, []any{}

	if v := patch.UpdatedTs; v != nil {
		set, args = append(set, "updated_ts = ?"), append(args, *v)
	}
	if v := patch.Email; v != nil {
		set, args = append(set, "email = ?"), append(args, *v)
	}
	if v := patch.Reachable; v != nil {
		set, args = append(set, "reachable = ?"), append(args, *v)
	}

	args = append(args, patch.ID)
	fields := []string{"recipient.id", "recipient.email", "recipient.reachable", "recipient.created_ts", "recipient.updated_ts"}
	query := `
		UPDATE recipient
		SET ` + strings.Join(set, ", ") + `
		WHERE id = ?
		RETURNING ` + strings.Join(fields, ", ")
	var recipient Recipient
	dests := []any{
		&recipient.ID,
		&recipient.Email,
		&recipient.Reachable,
		&recipient.CreatedTs,
		&recipient.UpdatedTs,
	}

	if err := s.db.QueryRowContext(ctx, query, args...).Scan(dests...); err != nil {
		return nil, err
	}

	return &recipient, nil
}

const deleteRecipient = `
	DELETE FROM recipient
	WHERE id = ?
`

type RecipientDelete struct {
	ID int
}

func (s *Store) DeleteRecipient(ctx context.Context, delete *RecipientDelete) error {
	where, args := []string{"id = ?"}, []any{delete.ID}

	stmt := `DELETE FROM recipient WHERE ` + strings.Join(where, " AND ")
	result, err := s.db.ExecContext(ctx, stmt, args...)
	if err != nil {
		return err
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return errors.New("recipient not found")
	}

	return nil
}

func (s *Store) DeleteBulkRecipient(ctx context.Context, ids []int) error {
	tx, err := s.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}
	for _, id := range ids {
		_, err = tx.Exec(deleteRecipient, id)
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
