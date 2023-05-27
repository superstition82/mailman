package core

// Role is the type of a role.
type Role string

const (
	Host       Role = "HOST"
	Admin      Role = "ADMIN"
	NormalUser Role = "USER"
)

func (role Role) String() string {
	switch role {
	case Host:
		return "HOST"
	case Admin:
		return "ADMIN"
	case NormalUser:
		return "USER"
	}
	return "USER"
}

type User struct {
	ID        int   `json:"id"`
	CreatedAt int64 `json:"createdAt"`
	UpdatedAt int64 `json:"updatedAt"`

	// Domain specific fields
	Username     string `json:"username"`
	Role         Role   `json:"role"`
	PasswordHash string `json:"-"`
}

type UserCreate struct {
	// Domain specific fields
	Username     string `json:"username"`
	Role         Role   `json:"role"`
	Password     string `json:"password"`
	PasswordHash string
}

type UserPatch struct {
	ID int `json:"-"`

	// Standard fields
	UpdatedAt *int64

	// Domain specific fields
	Username     *string `json:"username"`
	Password     *string `json:"password"`
	PasswordHash *string
}
