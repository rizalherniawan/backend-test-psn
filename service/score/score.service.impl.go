package score

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/rizalherniawan/backend-test-psn/dto/request"
	"github.com/rizalherniawan/backend-test-psn/model"
	"github.com/rizalherniawan/backend-test-psn/repository/score"
	"github.com/rizalherniawan/backend-test-psn/repository/student"
)

type ScoreServiceImpl struct {
	scoreRepository   score.ScoreRepository
	studentRepository student.StudentRepository
}

func NewScoreServiceImpl(scoreRepository score.ScoreRepository, studentRepository student.StudentRepository) ScoreService {
	return &ScoreServiceImpl{
		scoreRepository:   scoreRepository,
		studentRepository: studentRepository,
	}
}

func (s *ScoreServiceImpl) Create(studentId string, payload request.ScoreRequest) error {

	id, _ := strconv.Atoi(studentId)

	_, err := s.studentRepository.FindById(id)

	if err != nil {
		return err
	}

	_, err2 := s.studentRepository.FindBySubjectAndStudentId(payload.Subject, studentId)

	if err2 != nil {
		score := model.StudentScore{
			Subjects:  payload.Subject,
			StudentId: id,
			Score:     payload.Score,
			CreatedBy: payload.CreatedBy,
		}

		s.scoreRepository.Create(score)
		return nil
	}

	return errors.New("student with that subject already exist")
}

func (s *ScoreServiceImpl) UpdateScoreBySubjectAndStucentId(subject string, studentId string, payload request.ScoreRequest) error {

	res, err := s.studentRepository.FindBySubjectAndStudentId(subject, studentId)

	if err != nil {
		return err
	}

	fmt.Println(res)

	updateField := res.StudentScores[0]

	now := time.Now()

	updateField.Score = payload.Score
	updateField.ModifiedBy = &payload.ModifiedBy
	updateField.DateModified = &now

	s.scoreRepository.UpdateScore(updateField)

	return nil
}
