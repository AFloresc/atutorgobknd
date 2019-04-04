package rbac

import (
	"context"

	"github.com/stretchr/testify/assert"

	"testing"

	"github.com/websays-intelligence/profiles"
	"github.com/websays-intelligence/profiles/pmocks"
)

func UserProfilesForTest() (profiles profiles.UserProfilesClient) {
	return pmocks.NewUserProfilesMock()
}
func TestCheckUserAccess(t *testing.T) {
	mock := UserProfilesForTest()

	ctx := context.Background()
	mock.AssociateProfile(ctx, int64(1731), int64(1234))
	controller := NewAccessController(mock)

	t.Run("TestNoTokenInContext", func(t *testing.T) {
		token, err := controller.CheckUserAccess(ctx, "modify", int64(1234))
		assert := assert.New(t)
		assert.NotNil(err)
		assert.Empty(token.UserID())
		assert.Empty(token.Roles())
	})
	jwtToken, err := NewJWTTokenForTest()
	if err != nil {
		t.Fatal(err)
	}

	authContext := context.WithValue(ctx, "user", jwtToken)

	t.Run("TestUserHasAccess", func(t *testing.T) {
		token, err := controller.CheckUserAccess(authContext, "modify", int64(1234))
		assert := assert.New(t)
		assert.Nil(err)
		assert.NotEmpty(token.UserID())
		assert.NotEmpty(token.Roles())
	})

	t.Run("TestUserRoleError", func(t *testing.T) {
		token, err := controller.CheckUserAccess(authContext, "admin", int64(1234))
		assert := assert.New(t)
		assert.NotNil(err)
		assert.NotEmpty(token.UserID())
		assert.NotEmpty(token.Roles())
	})

	t.Run("TestUserProfileError", func(t *testing.T) {
		token, err := controller.CheckUserAccess(authContext, "modify", int64(1235))
		assert := assert.New(t)
		assert.NotNil(err)
		assert.NotEmpty(token.UserID())
		assert.NotEmpty(token.Roles())
	})

	t.Run("TestNoProfile", func(t *testing.T) {
		token, err := controller.CheckUserAccess(authContext, "modify")
		assert := assert.New(t)
		assert.Nil(err)
		assert.NotEmpty(token.UserID())
		assert.NotEmpty(token.Roles())
	})

}
