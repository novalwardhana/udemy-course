package simple

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserService(t *testing.T) {
	service := InitializedUserService()
	assert.NotNil(t, service)
}

func TestCategoryClassService(t *testing.T) {
	service := InitializedCategoryClass(true)
	assert.NotNil(t, service)
}

func TestAdminController(t *testing.T) {
	service := InitializedAdminController(true)
	assert.NotNil(t, service)
}
