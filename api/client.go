package api

import (
	"fmt"
	"probemail/util"
)

type Client struct {
	ID int `json:"id"`

	// Standard fields
	CreatedAt int64 `json:"createdAt"`
	UpdatedAt int64 `json:"updatedAt"`

	// Domain specific fields
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateClientRequest struct {
	// Domain specific fields
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (create CreateClientRequest) Validate() error {
	if create.Email != "" {
		if len(create.Email) > 256 {
			return fmt.Errorf("email is too long, maximum length is 256")
		}
		if !util.ValidateEmail(create.Email) {
			return fmt.Errorf("invalid email format")
		}
	}
	if len(create.Password) > 512 {
		return fmt.Errorf("password is too long, maximum length is 512")
	}

	return nil
}

type FindClientRequest struct {
	ID *int `json:"id"`

	// Domain specific fields
	Email string `json:"email"`
}

type DeleteClientRequest struct {
	ID int
}
