//go:build wireinject

// + wireinject

package simple

import "github.com/google/wire"

func InitializedStudentService(log bool) (*StudentService, error) {
	wire.Build(NewStudentRepository, NewStudentService)
	return nil, nil
}

func InitializedStudentController(log bool) (*StudentController, error) {
	wire.Build(NewStudentRepository, NewStudentService, NewStudentController)
	return nil, nil
}
