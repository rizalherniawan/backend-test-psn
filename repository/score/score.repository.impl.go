package score

import (
	"github.com/rizalherniawan/backend-test-psn/model"
	"gorm.io/gorm"
)

type ScoreRepositoryImpl struct {
	db *gorm.DB
}

func NewScoreRepositoryImpl(db *gorm.DB) ScoreRepository {
	return &ScoreRepositoryImpl{db: db}
}

func (s *ScoreRepositoryImpl) Create(score model.StudentScore) {
	result := s.db.Table("student_scores").Create(&score)
	err := result.Error

	if err != nil {
		panic(err.Error())
	}
}

func (s *ScoreRepositoryImpl) UpdateScore(existingScore model.StudentScore) {
	s.db.Save(&existingScore)
}
