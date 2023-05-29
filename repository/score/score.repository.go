package score

import "github.com/rizalherniawan/backend-test-psn/model"

type ScoreRepository interface {
	Create(score model.StudentScore)
	UpdateScore(existingScore model.StudentScore)
}
