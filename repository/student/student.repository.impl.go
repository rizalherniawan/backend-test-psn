package student

import (
	"errors"

	"github.com/rizalherniawan/backend-test-psn/model"
	"gorm.io/gorm"
)

type StudentRepositoryImpl struct {
	db *gorm.DB
}

func NewStudentRepositoryImpl(db *gorm.DB) StudentRepository {
	return &StudentRepositoryImpl{db: db}
}

func (s *StudentRepositoryImpl) CreateStudent(student model.Student) {
	result := s.db.Table("students").Create(&student)
	err := result.Error

	if err != nil {
		panic(err.Error())
	}
}

func (s *StudentRepositoryImpl) FindById(id int) (model.Student, error) {
	var student model.Student
	result := s.db.Table("students").First(&student, id)

	if result.Error != nil {
		return student, errors.New("student not found")
	}

	return student, nil
}

func (s *StudentRepositoryImpl) FindBySubject(subject string) ([]model.Student, error) {
	var students []model.Student
	res := s.db.
		Joins("left join student_scores on students.id = student_scores.student_id").
		Where("student_scores.subjects = ?", subject).
		Preload("StudentScores", "student_scores.subjects = ?", subject).
		Find(&students)

	if res.RowsAffected == 0 {
		return students, errors.New("subject not found")
	}

	return students, nil
}

func (s *StudentRepositoryImpl) FindBySubjectAndStudentId(subject string, studentId string) (model.Student, error) {
	var student model.Student
	res := s.db.
		Joins("left join student_scores on students.id = student_scores.student_id").
		Where("student_scores.subjects = ? and students.id = ?", subject, studentId).
		Preload("StudentScores", "student_scores.subjects = ?", subject).
		Find(&student)

	if res.RowsAffected == 0 {
		return student, errors.New("student with this subject is not found")
	}

	return student, nil
}

func (s *StudentRepositoryImpl) DeleteStudent(student model.Student) {
	s.db.Delete(&student)
}
