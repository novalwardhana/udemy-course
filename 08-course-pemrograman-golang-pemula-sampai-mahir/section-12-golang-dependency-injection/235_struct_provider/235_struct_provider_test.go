package simple

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCategoryClass(t *testing.T) {
	service := InitializedCategoryClass()
	assert.NotNil(t, service)
}

func TestDatabaseRepository(t *testing.T) {
	service := InitializedDatabase()
	assert.NotNil(t, service)
}
