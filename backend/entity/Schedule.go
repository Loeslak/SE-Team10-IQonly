package entity

import (
	"time"

	"gorm.io/gorm"
)

type Schedule struct {
	gorm.Model
	ScheduleName    string
	Trainingstart   time.Time
	TrainingEnd     time.Time
	DifficultyLevel int
	Describtion     string

	WorkoutID *uint
	Workout   Workout

	TrainingLocationID *uint
	TrainingLocation   TrainingLocation

	TrainerID *uint
	Trainer   Trainer

	TrainerBookingID *uint
	TrainerBooking   TrainerBooking
}
