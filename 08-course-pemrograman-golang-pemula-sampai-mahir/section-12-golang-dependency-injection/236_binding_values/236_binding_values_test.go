package simple

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDatabase1(t *testing.T) {
	db := InitializationDatabase1()
	assert.NotNil(t, db)
}

func TestDatabase2(t *testing.T) {
	db := InitializationDatabase2()
	assert.NotNil(t, db)
}

func TestCategory(t *testing.T) {
	category := InitializedCategory()
	assert.NotNil(t, category)
}

func TestUser(t *testing.T) {
	user := InitializedUserController()
	assert.NotNil(t, user)
}
