package entity

import (
	"gorm.io/gorm"
)

type Workout = struct {
	gorm.Model
	WorkoutName string

	Schedule []Schedule `gorm:"foreignKey:WorkoutID"`
}
