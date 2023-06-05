package store

import (
	"context"
)

type User struct {
	ID int

	// Standard fields
	CreatedTs int64
	UpdatedTs int64

	// Domain specific fields
	Username     string
	Role         string
	PasswordHash string
}

const createUser = `
INSERT INTO user (username, role, password_hash) 
VALUES (?, ?, ?)
RETURNING id, username, role, password_hash, created_ts, updated_ts
`

type CreateUserParams struct {
	Username     string
	Role         string
	PasswordHash string
}

func (s *Store) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := s.db.QueryRowContext(ctx, createUser, arg.Username, arg.Role, arg.PasswordHash)
	var user User
	err := row.Scan(&user.ID, &user.Username, &user.Role, &user.PasswordHash, &user.CreatedTs, &user.UpdatedTs)

	return user, err
}

const deleteUser = `
DELETE FROM user
WHERE id = ?
`

func (s *Store) DeleteUser(ctx context.Context, id int) error {
	_, err := s.db.ExecContext(ctx, deleteUser, id)

	return err
}

const getUser = `
SELECT id, username, role, password_hash, created_ts, updated_ts FROM user
WHERE id = ? LIMIT 1
`

func (s *Store) GetUser(ctx context.Context, id int) (User, error) {
	row := s.db.QueryRowContext(ctx, getUser, id)
	var user User
	err := row.Scan(&user.ID, &user.Username, &user.Role, &user.PasswordHash, &user.CreatedTs, &user.UpdatedTs)

	return user, err
}

const listUseres = `
SELECT id, username, role, password_hash, created_ts, updated_ts FROM user
ORDER BY id LIMIT ? OFFSET ?
`

type ListUsersParams struct {
	Limit  int64
	Offset int64
}

func (s *Store) ListUsers(ctx context.Context, arg ListRecipientsParams) ([]User, error) {
	rows, err := s.db.QueryContext(ctx, listUseres, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Username, &user.Role, &user.PasswordHash, &user.CreatedTs, &user.UpdatedTs); err != nil {
			return nil, err
		}
		items = append(items, user)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}
