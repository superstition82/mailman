package core

import (
	"github.com/golang-jwt/jwt/v4"
)

const (
	// Context section
	// The key name used to store user id in the context
	// user id is extracted from the jwt token subject field.
	userIDContextKey = "user-id"
)

// Claims creates a struct that will be encoded to a JWT.
// We add jwt.RegisteredClaims as an embedded type, to provide fields such as name.
type Claims struct {
	Name string `json:"name"`
	jwt.RegisteredClaims
}

func GetUserIDContextKey() string {
	return userIDContextKey
}

// GenerateTokensAndSetCookies generates jwt token and saves it to the http-only cookie.
