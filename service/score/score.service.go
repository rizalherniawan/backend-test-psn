package score

import (
	"github.com/rizalherniawan/backend-test-psn/dto/request"
)

type ScoreService interface {
	Create(studentId string, payload request.ScoreRequest) error
	UpdateScoreBySubjectAndStucentId(subject string, studentId string, payload request.ScoreRequest) error
}
