package entity

import (
	"gorm.io/gorm"
)

type Member struct {
	gorm.Model
	MemberUsername string
	MemberEmail    string
	MemberPassword string
	MemberStatus   bool

	TrainerBooking []TrainerBooking `gorm:"foreignKey:MemberID"`
}
