package model

import (
	"time"
)

type StudentScore struct {
	Id           int        `json:"id" gorm:"id;primaryKey"`
	Subjects     string     `json:"subject" gorm:"subjects;not null"`
	StudentId    int        `json:"studentId" gorm:"student_id;not null"`
	Score        float64    `json:"score" gorm:"score;not null"`
	DateCreated  time.Time  `gorm:"date_created;autoUpdateTime:milli"`
	CreatedBy    string     `json:"createdBy" gorm:"created_by;not null"`
	DateModified *time.Time `json:"date_modified;dateModified"`
	ModifiedBy   *string    `json:"modified_by;modifiedBy"`
}
