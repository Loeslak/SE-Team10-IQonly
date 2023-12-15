package entity

import (
	"gorm.io/gorm"
)

type Trainer struct {
	gorm.Model
	TrainerUsername string
	TrainerEmail    string
	TrainerPassword string

	TrainerInformationID *uint
	TrainerInformation   TrainerInformation

	Schedule []Schedule `gorm:"foreignKey:TrainerID"`
}
