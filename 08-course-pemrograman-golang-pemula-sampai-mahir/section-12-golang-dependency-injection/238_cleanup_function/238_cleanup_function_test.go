package simple

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserController(t *testing.T) {
	controller, cleanup := InitializedUserController("Noval")

	assert.NotNil(t, controller)
	cleanup()
}

func TestAdminController(t *testing.T) {
	controller, cleanup, error := InitializedAdminController("my admin")
	assert.NotNil(t, controller)
	assert.Nil(t, error)
	cleanup()
}
