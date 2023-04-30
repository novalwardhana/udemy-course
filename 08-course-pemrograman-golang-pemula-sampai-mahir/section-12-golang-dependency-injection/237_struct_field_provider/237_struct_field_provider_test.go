package simple

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCategory(t *testing.T) {
	category := InitializeCategoryRepository()
	assert.NotNil(t, category)
}

func TestMU(t *testing.T) {
	mu := mu()
	assert.NotNil(t, mu)
}

func TestChelsea(t *testing.T) {
	chelsea := chelsea()
	assert.NotNil(t, chelsea)
}

func TestArsenal(t *testing.T) {
	arsenal := arsenal()
	assert.NotNil(t, arsenal)
}
