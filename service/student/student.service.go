package student

import (
	"github.com/rizalherniawan/backend-test-psn/dto/request"
	"github.com/rizalherniawan/backend-test-psn/model"
)

type StudentService interface {
	Create(payload request.StudentRequest)
	FindBySubject(subject string) ([]model.Student, error)
	FindBySubjectAndStudentId(subject string, studentId string) (model.Student, error)
	DeleteStudentById(studentId string) error
}
