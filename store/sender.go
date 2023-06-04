package store

import "context"

type Sender struct {
	ID int `json:"id"`

	// Domain specific fields
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Email    string `json:"email"`
	Password string `json:"password"`

	// Standard fields
	CreatedTs int64 `json:"created_ts"`
	UpdatedTs int64 `json:"updated_ts"`
}

const createSender = `
	INSERT INTO sender (host, port, email, password)
	VALUES (?, ?, ?, ?)
	RETURNING id, host, port, email, password, created_ts, updated_ts
`

type CreateSenderParams struct {
	Host     string
	Port     int
	Email    string
	Password string
}

func (s *Store) CreateSender(ctx context.Context, arg CreateSenderParams) (Sender, error) {
	row := s.db.QueryRowContext(ctx, createSender, arg.Host, arg.Port, arg.Email, arg.Password)
	var sender Sender
	err := row.Scan(
		&sender.ID,
		&sender.Host,
		&sender.Port,
		&sender.Email,
		&sender.Password,
		&sender.CreatedTs,
		&sender.UpdatedTs,
	)
	return sender, err
}

const getSender = `
	SELECT id, host, port, email, password, created_ts, updated_ts
	FROM sender
	WEHRE id = ? LIMIT 1	
`

func (s *Store) GetSender(ctx context.Context, id int) (Sender, error) {
	row := s.db.QueryRowContext(ctx, getSender, id)
	var sender Sender
	err := row.Scan(
		&sender.ID,
		&sender.Host,
		&sender.Port,
		&sender.Email,
		&sender.Password,
		&sender.CreatedTs,
		&sender.UpdatedTs,
	)
	return sender, err
}

const listSenders = `
	SELECT id, host, port, email, password, created_ts, updated_ts
	FROM sender
	ORDER BY id LIMIT ? OFFSET ?
`

type ListSendersParams struct {
	Limit  int
	Offset int
}

func (s *Store) ListSenders(ctx context.Context, arg ListSendersParams) ([]Sender, error) {
	rows, err := s.db.QueryContext(ctx, listSenders, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []Sender
	for rows.Next() {
		var sender Sender
		if err := rows.Scan(
			&sender.ID,
			&sender.Host,
			&sender.Port,
			&sender.Email,
			&sender.Password,
			&sender.CreatedTs,
			&sender.UpdatedTs,
		); err != nil {
			return nil, err
		}
		items = append(items, sender)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const deleteSender = `
	DELETE FROM sender
	WHERE id = ?
`

func (s *Store) DeleteSender(ctx context.Context, id int) error {
	_, err := s.db.ExecContext(ctx, deleteSender, id)
	return err
}
