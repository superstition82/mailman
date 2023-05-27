package auth

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

const (
	issuer                   = "probemail"
	keyID                    = "v1"
	AccessTokenAudienceName  = "user.access-token"
	RefreshTokenAudienceName = "user.refresh-token"
	apiTokenDuration         = 2 * time.Hour
	accessTokenDuration      = 24 * time.Hour
	refreshTokenDuration     = 7 * 24 * time.Hour
	RefreshThresholdDuration = 1 * time.Hour
	CookieExpDuration        = refreshTokenDuration - 1*time.Minute
	AccessTokenCookieName    = "access-token"
	RefreshTokenCookieName   = "refresh-token"
	UserIDCookieName         = "user"
)

type claimsMessage struct {
	Name string `json:"name"`
	jwt.RegisteredClaims
}

// GenerateAPIToken generates an API token.
func GenerateAPIToken(userName string, userID int, secret string) (string, error) {
	expirationTime := time.Now().Add(apiTokenDuration)
	return generateToken(userName, userID, AccessTokenAudienceName, expirationTime, []byte(secret))
}

// GenerateAccessToken generates an access token for web.
func GenerateAccessToken(userName string, userID int, secret string) (string, error) {
	expirationTime := time.Now().Add(accessTokenDuration)
	return generateToken(userName, userID, AccessTokenAudienceName, expirationTime, []byte(secret))
}

// GenerateRefreshToken generates a refresh token for web.
func GenerateRefreshToken(userName string, userID int, secret string) (string, error) {
	expirationTime := time.Now().Add(refreshTokenDuration)
	return generateToken(userName, userID, RefreshTokenAudienceName, expirationTime, []byte(secret))
}

func generateToken(username string, userID int, aud string, expirationTime time.Time, secret []byte) (string, error) {
	// Create the JWT claims, which includes the username and expiry time.
	claims := &claimsMessage{
		Name: username,
		RegisteredClaims: jwt.RegisteredClaims{
			Audience: jwt.ClaimStrings{aud},
			// In JWT, the expiry time is expressed as unix milliseconds.
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    issuer,
			Subject:   strconv.Itoa(userID),
		},
	}

	// Declare the token with the HS256 algorithm used for signing, and the claims.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token.Header["kid"] = keyID

	// Create the JWT string.
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
