package simple

import "errors"

type StudentRepository struct {
	Log bool
}

func NewStudentRepository(log bool) *StudentRepository {
	return &StudentRepository{
		Log: log,
	}
}

type StudentService struct {
	StudentRepository *StudentRepository
}

func NewStudentService(studentRepository *StudentRepository) (*StudentService, error) {
	if studentRepository.Log {
		return nil, errors.New("log error true")
	}
	return &StudentService{
		StudentRepository: studentRepository,
	}, nil
}

type StudentController struct {
	StudentService *StudentService
}

func NewStudentController(studentService *StudentService) (*StudentController, error) {
	if studentService == nil {
		return nil, errors.New("log error true")
	}
	return &StudentController{
		StudentService: studentService,
	}, nil
}
