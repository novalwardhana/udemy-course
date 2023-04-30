// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package simple

// Injectors from injector.go:

func InitializedStudentService(log bool) (*StudentService, error) {
	studentRepository := NewStudentRepository(log)
	studentService, err := NewStudentService(studentRepository)
	if err != nil {
		return nil, err
	}
	return studentService, nil
}

func InitializedStudentController(log bool) (*StudentController, error) {
	studentRepository := NewStudentRepository(log)
	studentService, err := NewStudentService(studentRepository)
	if err != nil {
		return nil, err
	}
	studentController, err := NewStudentController(studentService)
	if err != nil {
		return nil, err
	}
	return studentController, nil
}
