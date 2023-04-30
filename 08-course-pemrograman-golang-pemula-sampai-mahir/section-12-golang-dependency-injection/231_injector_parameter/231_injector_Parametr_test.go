package simple

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStudentServiceSuccess(t *testing.T) {
	studentService, err := InitializedStudentService(false)
	assert.Nil(t, err)
	assert.NotNil(t, studentService)
}

func TestStudentServiceFailed(t *testing.T) {
	studentService, err := InitializedStudentService(true)
	assert.Nil(t, studentService)
	assert.NotNil(t, err)
}

func TestStudentControllerSuccess(t *testing.T) {
	studentController, err := InitializedStudentController(false)
	assert.Nil(t, err)
	assert.NotNil(t, studentController)
}

func TestStudentControllerFailed(t *testing.T) {
	studentController, err := InitializedStudentController(true)
	assert.Nil(t, studentController)
	assert.NotNil(t, err)
}
