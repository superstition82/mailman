package store

import (
	"context"
	"database/sql"
)

type Recipient struct {
	ID int `json:"id"`

	// Domain specific fields
	Email     string `json:"email"`
	Reachable string `json:"reachable"`

	// Standard fields
	CreatedTs int64 `json:"createdTs"`
	UpdatedTs int64 `json:"updatedTs"`
}

const createRecipient = `
	INSERT INTO recipient (email, reachable) 
	VALUES (?, ?)
	RETURNING id, email, reachable, created_ts, updated_ts
`

type CreateRecipientParams struct {
	Email     string
	Reachable string
}

func (s *Store) CreateRecipient(ctx context.Context, arg CreateRecipientParams) (Recipient, error) {
	row := s.db.QueryRowContext(ctx, createRecipient, arg.Email, arg.Reachable)
	var recipient Recipient
	err := row.Scan(&recipient.ID, &recipient.Email, &recipient.Reachable, &recipient.CreatedTs, &recipient.UpdatedTs)
	return recipient, err
}

const getRecipient = `
	SELECT id, email, reachable, created_ts, updated_ts 
	FROM recipient
	WHERE id = ? LIMIT 1
`

func (s *Store) GetRecipient(ctx context.Context, id int) (Recipient, error) {
	row := s.db.QueryRowContext(ctx, getRecipient, id)
	var recipient Recipient
	err := row.Scan(&recipient.ID, &recipient.Email, &recipient.Reachable, &recipient.CreatedTs, &recipient.UpdatedTs)
	return recipient, err
}

const listAllRecipients = `
	SELECT id, email, reachable, created_ts, updated_ts  
	FROM recipient
	ORDER BY id
`

func (s *Store) FindRecepientList(ctx context.Context) ([]Recipient, error) {
	rows, err := s.db.QueryContext(ctx, listAllRecipients)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Recipient
	for rows.Next() {
		var recipient Recipient
		if err := rows.Scan(&recipient.ID, &recipient.Email, &recipient.Reachable, &recipient.CreatedTs, &recipient.UpdatedTs); err != nil {
			return nil, err
		}
		items = append(items, recipient)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}

const listRecipients = `
	SELECT id, email, reachable, created_ts, updated_ts  
	FROM recipient
	ORDER BY id LIMIT ? OFFSET ?
`

type ListRecipientsParams struct {
	Limit  int
	Offset int
}

func (s *Store) ListRecipients(ctx context.Context, arg ListRecipientsParams) ([]Recipient, error) {
	rows, err := s.db.QueryContext(ctx, listRecipients, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []Recipient
	for rows.Next() {
		var recipient Recipient
		if err := rows.Scan(&recipient.ID, &recipient.Email, &recipient.Reachable, &recipient.CreatedTs, &recipient.UpdatedTs); err != nil {
			return nil, err
		}
		items = append(items, recipient)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}

const updateRecipient = `
	UPDATE recipient
	SET 
		email = COALESCE(?, email),
		reachable = COALESCE(?, reachable)
	WHERE
		id = ?
	RETURNING id, email, reachable, created_ts, updated_ts
`

type UpdateRecipientParams struct {
	ID int `json:"id"`

	// Domain specific fields
	Email     string `json:"email"`
	Reachable string `json:"reachable"`
}

func (s *Store) UpdateRecipient(ctx context.Context, arg UpdateRecipientParams) (Recipient, error) {
	row := s.db.QueryRowContext(ctx, updateRecipient, arg.Email, arg.Reachable, arg.ID)
	var recipient Recipient
	err := row.Scan(&recipient.ID, &recipient.Email, &recipient.Reachable, &recipient.CreatedTs, &recipient.UpdatedTs)
	return recipient, err
}

const deleteRecipient = `
	DELETE FROM recipient
	WHERE id = ?
`

func (s *Store) DeleteRecipient(ctx context.Context, id int) error {
	_, err := s.db.ExecContext(ctx, deleteRecipient, id)
	return err
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
