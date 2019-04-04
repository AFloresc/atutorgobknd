package rbac

import (
	"strconv"

	jwt "github.com/dgrijalva/jwt-go"
)

// Token contains authentication data
type Token struct {
	jwtToken *jwt.Token
	userID   int64
	roles    []Role
}

// NewToken initializes a token from a jwt.Token
func NewToken(jwtToken *jwt.Token) (token Token, err error) {
	token = Token{
		jwtToken: jwtToken,
	}

	claims := jwtToken.Claims.(jwt.MapClaims)

	token.userID, err = strconv.ParseInt(claims["user_id"].(string), 10, 64)
	if err != nil {
		return
	}

	roles := claims["roles"].([]interface{})
	for _, role := range roles {
		token.roles = append(token.roles, Role(role.(string)))
	}
	return
}

// UserID returns the user id.
func (t Token) UserID() int64 {
	return t.userID
}

// Roles returns the list of roles
func (t Token) Roles() []Role {
	return t.roles
}
