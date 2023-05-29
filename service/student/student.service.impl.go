package student

import (
	"strconv"

	"github.com/rizalherniawan/backend-test-psn/dto/request"
	"github.com/rizalherniawan/backend-test-psn/model"
	"github.com/rizalherniawan/backend-test-psn/repository/student"
)

type StudentServiceImpl struct {
	studentRepository student.StudentRepository
}

func NewStudentServiceImpl(studentRepository student.StudentRepository) StudentService {
	return &StudentServiceImpl{studentRepository: studentRepository}
}

func (s *StudentServiceImpl) Create(payload request.StudentRequest) {
	student := model.Student{
		Name:      payload.Name,
		Age:       payload.Age,
		CreatedBy: payload.CreatedBy,
	}
	s.studentRepository.CreateStudent(student)
}

func (s *StudentServiceImpl) FindBySubject(subject string) ([]model.Student, error) {
	return s.studentRepository.FindBySubject(subject)
}

func (s *StudentServiceImpl) FindBySubjectAndStudentId(subject string, studentId string) (model.Student, error) {
	return s.studentRepository.FindBySubjectAndStudentId(subject, studentId)
}

func (s *StudentServiceImpl) DeleteStudentById(studentId string) error {

	id, _ := strconv.Atoi(studentId)
	res, err := s.studentRepository.FindById(id)

	if err != nil {
		return err
	}

	s.studentRepository.DeleteStudent(res)

	return nil
}
