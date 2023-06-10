package store

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
)

type Sender struct {
	ID int `json:"id"`

	// Standard fields
	CreatedTs int64 `json:"createdTs"`
	UpdatedTs int64 `json:"updatedTs"`

	// Domain specific fields
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Email    string `json:"email"`
	Password string `json:"password"`
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

func (s *Store) CreateSender(ctx context.Context, arg CreateSenderParams) (*Sender, error) {
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
	return &sender, err
}

type SenderFind struct {
	ID *int `json:"id"`

	// Pagination
	Limit  *int
	Offset *int
}

func (s *Store) FindSenderList(ctx context.Context, find *SenderFind) ([]*Sender, error) {
	where, args := []string{"1 = 1"}, []any{}

	if v := find.ID; v != nil {
		where, args = append(where, "sender.id = ?"), append(args, *v)
	}

	fields := []string{"sender.id", "sender.host", "sender.port", "sender.email", "sender.password", "sender.created_ts", "sender.updated_ts"}
	query := fmt.Sprintf(`
		SELECT %s
		FROM sender
		WHERE %s
		ORDER BY sender.id DESC
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

	senderList := make([]*Sender, 0)
	for rows.Next() {
		var sender Sender
		dests := []any{
			&sender.ID,
			&sender.Host,
			&sender.Port,
			&sender.Email,
			&sender.Password,
			&sender.CreatedTs,
			&sender.UpdatedTs,
		}
		if err := rows.Scan(dests...); err != nil {
			return nil, err
		}
		senderList = append(senderList, &sender)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return senderList, nil
}

func (s *Store) FindSender(ctx context.Context, find *SenderFind) (*Sender, error) {
	list, err := s.FindSenderList(ctx, find)
	if err != nil {
		return nil, err
	}
	if len(list) == 0 {
		return nil, errors.New("not found")
	}

	sender := list[0]

	return sender, nil
}

const deleteSender = `
	DELETE FROM sender
	WHERE id = ?
`

type SenderDelete struct {
	ID int
}

func (s *Store) DeleteSender(ctx context.Context, delete *SenderDelete) error {
	where, args := []string{"id = ?"}, []any{delete.ID}

	stmt := `DELETE FROM sender WHERE ` + strings.Join(where, " AND ")
	result, err := s.db.ExecContext(ctx, stmt, args...)
	if err != nil {
		return err
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return errors.New("sender not found")
	}

	return nil
}

type SenderBulkDelete struct {
	IDs []int
}

func (s *Store) DeleteBulkSender(ctx context.Context, bulkDelete *SenderBulkDelete) error {
	tx, err := s.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}

	for _, id := range bulkDelete.IDs {
		_, err = tx.Exec(deleteSender, id)
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
