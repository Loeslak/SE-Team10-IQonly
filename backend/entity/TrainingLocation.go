package entity

import (
	"gorm.io/gorm"
)

type TrainingLocation struct {
	gorm.Model
	Building string
	Floor    string
	Room     string

	Schedule []Schedule `gorm:"foreignKey:TrainingLocationID"`
}
