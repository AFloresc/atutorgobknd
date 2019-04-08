package rbac

import (
	"context"
	"fmt"

	"github.com/atutor/domain"

	jwt "github.com/dgrijalva/jwt-go"
)

// AccessController helps controlling access to the services
type AccessController struct {
	user domain.UserClient
}

// NewAccessController initializes a new access controller
func NewAccessController(user domain.UserClient) *AccessController {
	return &AccessController{
		user: user,
	}
}

func getTokenFromContext(ctx context.Context) (token Token, err error) {
	value := ctx.Value("user")
	if value == nil {
		return Token{}, fmt.Errorf("No token found")
	}
	jwtToken := ctx.Value("user").(*jwt.Token)
	return NewToken(jwtToken)
}

// CheckUserAccess Checks if the user (stored in context) has the specified role and has access to the specified profiles
func (ac *AccessController) CheckUserAccess(ctx context.Context, role Role, profileIDs ...int64) (token Token, err error) {
	token, err = getTokenFromContext(ctx)
	if err != nil {
		return
	}

	if roleInSlice(RoleAdmin, token.Roles()) {
		// Admin has full access to all profiles
		return token, nil
	}

	if !roleInSlice(role, token.Roles()) {
		return token, fmt.Errorf("user has no access to the profile (role)")
	}

	if len(profileIDs) == 0 {
		return token, nil
	}

	profiles, err := ac.profiles.GetUserProfileIDs(ctx, token.UserID())
	if err != nil {
		return
	}
	for _, profileID := range profileIDs {

		if !int64InSlice(profileID, profiles) {
			return token, fmt.Errorf("user has no access to the profile (profile)")
		}
	}

	return
}

func int64InSlice(a int64, list []int64) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func roleInSlice(a Role, list []Role) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
