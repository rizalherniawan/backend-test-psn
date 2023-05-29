package model

import "time"

type Student struct {
	Id            int            `json:"id" gorm:"id;primaryKey"`
	Name          string         `json:"name" gorm:"name;not null"`
	Age           int            `json:"age" gorm:"age;not null"`
	DateCreated   time.Time      `gorm:"date_created;autoUpdateTime:milli"`
	CreatedBy     string         `json:"createdBy" gorm:"created_by;not null"`
	StudentScores []StudentScore `json:"studentScores" gorm:"foreignKey:StudentId;references:Id"`
}
