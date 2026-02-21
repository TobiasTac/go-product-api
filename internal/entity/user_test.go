package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("Tobias Alves", "tobias.vida@live.com", "123123")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, "Tobias Alves", user.Name)
	assert.Equal(t, "tobias.vida@live.com", user.Email)
}
func TestUser_ValidatePassword(t *testing.T) {
	user, err := NewUser("Tobias Alves", "tobias.vida@live.com", "123123")
	assert.Nil(t, err)
	assert.True(t, user.ValidatePassword("123123"))
	assert.False(t, user.ValidatePassword("123456"))
	assert.NotEqual(t, "123123", user.Password)
}
