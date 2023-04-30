package test

import (
	"belajar-golang-restful-api/simple"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFileConnection(t *testing.T) {
	connect, cleanup := simple.InitializedConnection("database")
	assert.NotNil(t, connect)

	cleanup()
}
