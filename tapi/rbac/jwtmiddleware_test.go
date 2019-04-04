package rbac

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewJWTMiddlware(t *testing.T) {
	jwtMW := NewJWTMiddlware(NewAuthenticationKeyForTest())
	assert := assert.New(t)
	assert.NotNil(jwtMW)
}
