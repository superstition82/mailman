// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package db

import ()

type Recepient struct {
	ID     int64
	Email  string
	Status string
}

type Sender struct {
	ID       int64
	Email    string
	Password string
}

type Template struct {
	ID      int64
	Subject string
	Body    string
}
