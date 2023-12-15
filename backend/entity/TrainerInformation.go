package entity

import (
	"gorm.io/gorm"
)

type TrainerInformation struct {
	gorm.Model
	TrainerFirstName string
	TrainerLastName  string
	TrainerBirthday  string
	TrainerAge       int
	TrainerWeight    float32
	TrainerHeight    float32
	TrainerPhone     string
	TrainerProfile   string

	GenderID *uint
	Gender   Gender

	TrainerBooking []TrainerBooking `gorm:"foreignKey:TrainerInformationID"`
	Trainer        []Trainer        `gorm:"foreignKey:TrainerInformationID"`
}
