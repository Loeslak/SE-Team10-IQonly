package entity

import (
	"gorm.io/gorm"
)

type Gender struct {
	gorm.Model
	Sex string

	TrainerInformation []TrainerInformation `gorm:"foreignKey:GenderID"`
}
