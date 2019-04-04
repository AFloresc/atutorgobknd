package rbac

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewToken(t *testing.T) {
	jwtToken, err := NewJWTTokenForTest()
	if err != nil {
		t.Error(err)
	}
	token, err := NewToken(jwtToken)
	if err != nil {
		t.Error(err)
	}

	assert := assert.New(t)
	assert.Equal(int64(1731), token.UserID())
	assert.Contains(token.Roles(), RoleLogin)
	assert.Contains(token.Roles(), RoleModify)
	assert.Contains(token.Roles(), RoleShare)
	assert.Contains(token.Roles(), RoleAPI)
	assert.Contains(token.Roles(), RolePublisher)
	assert.Contains(token.Roles(), RoleModifyClippingsAndProfile)
	assert.Contains(token.Roles(), RoleModifyTopics)
	assert.Contains(token.Roles(), RoleTags)
}
