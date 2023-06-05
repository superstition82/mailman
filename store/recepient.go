package store

import (
	"context"
	"database/sql"
)

type Recepient struct {
	ID int `json:"id"`

	// Domain specific fields
	Email     string `json:"email"`
	Reachable string `json:"reachable"`

	// Standard fields
	CreatedTs int64 `json:"createdTs"`
	UpdatedTs int64 `json:"updatedTs"`
}

const createRecepient = `
	INSERT INTO recepient (email, reachable) 
	VALUES (?, ?)
	RETURNING id, email, reachable, created_ts, updated_ts
`

type CreateRecepientParams struct {
	Email     string
	Reachable string
}

func (s *Store) CreateRecepient(ctx context.Context, arg CreateRecepientParams) (Recepient, error) {
	row := s.db.QueryRowContext(ctx, createRecepient, arg.Email, arg.Reachable)
	var recepient Recepient
	err := row.Scan(&recepient.ID, &recepient.Email, &recepient.Reachable, &recepient.CreatedTs, &recepient.UpdatedTs)
	return recepient, err
}

const getRecepient = `
	SELECT id, email, reachable, created_ts, updated_ts 
	FROM recepient
	WHERE id = ? LIMIT 1
`

func (s *Store) GetRecepient(ctx context.Context, id int) (Recepient, error) {
	row := s.db.QueryRowContext(ctx, getRecepient, id)
	var recepient Recepient
	err := row.Scan(&recepient.ID, &recepient.Email, &recepient.Reachable, &recepient.CreatedTs, &recepient.UpdatedTs)
	return recepient, err
}

const listAllRecepients = `
	SELECT id, email, reachable, created_ts, updated_ts  
	FROM recepient
	ORDER BY id
`

func (s *Store) ListAllRecepients(ctx context.Context) ([]Recepient, error) {
	rows, err := s.db.QueryContext(ctx, listAllRecepients)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Recepient
	for rows.Next() {
		var recepient Recepient
		if err := rows.Scan(&recepient.ID, &recepient.Email, &recepient.Reachable, &recepient.CreatedTs, &recepient.UpdatedTs); err != nil {
			return nil, err
		}
		items = append(items, recepient)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}

const listRecepients = `
	SELECT id, email, reachable, created_ts, updated_ts  
	FROM recepient
	ORDER BY id LIMIT ? OFFSET ?
`

type ListRecepientsParams struct {
	Limit  int
	Offset int
}

func (s *Store) ListRecepients(ctx context.Context, arg ListRecepientsParams) ([]Recepient, error) {
	rows, err := s.db.QueryContext(ctx, listRecepients, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []Recepient
	for rows.Next() {
		var recepient Recepient
		if err := rows.Scan(&recepient.ID, &recepient.Email, &recepient.Reachable, &recepient.CreatedTs, &recepient.UpdatedTs); err != nil {
			return nil, err
		}
		items = append(items, recepient)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}

const updateRecepient = `
	UPDATE recepient
	SET 
		email = COALESCE(?, email),
		reachable = COALESCE(?, reachable)
	WHERE
		id = ?
	RETURNING id, email, reachable, created_ts, updated_ts
`

type UpdateRecepientParams struct {
	ID int `json:"id"`

	// Domain specific fields
	Email     string `json:"email"`
	Reachable string `json:"reachable"`
}

func (s *Store) UpdateRecepient(ctx context.Context, arg UpdateRecepientParams) (Recepient, error) {
	row := s.db.QueryRowContext(ctx, updateRecepient, arg.Email, arg.Reachable, arg.ID)
	var recepient Recepient
	err := row.Scan(&recepient.ID, &recepient.Email, &recepient.Reachable, &recepient.CreatedTs, &recepient.UpdatedTs)
	return recepient, err
}

const deleteRecepient = `
	DELETE FROM recepient
	WHERE id = ?
`

func (s *Store) DeleteRecepient(ctx context.Context, id int) error {
	_, err := s.db.ExecContext(ctx, deleteRecepient, id)
	return err
}

func (s *Store) DeleteBulkRecepient(ctx context.Context, ids []int) error {
	tx, err := s.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}

	for _, id := range ids {
		_, err = tx.Exec(deleteRecepient, id)
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
