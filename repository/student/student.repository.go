package student

import (
	"github.com/rizalherniawan/backend-test-psn/model"
)

type StudentRepository interface {
	CreateStudent(student model.Student)
	FindById(id int) (model.Student, error)
	FindBySubject(subject string) ([]model.Student, error)
	FindBySubjectAndStudentId(subject string, studentId string) (model.Student, error)
	DeleteStudent(student model.Student)
}
