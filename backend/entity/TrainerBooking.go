package entity

import (
	"gorm.io/gorm"
)

type TrainerBooking struct {
	gorm.Model

	MemberID *uint
	Member   Member

	TrainerInformationID *uint
	TrainerInformation   TrainerInformation

	Schedule []Schedule `gorm:"foreignKey:TrainerBookingID"`
}
