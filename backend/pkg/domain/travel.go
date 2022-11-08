package domain

import (
	"time"

	"gorm.io/gorm"
)

type statusEnum string

const (
	PENDING   statusEnum = "pending"
	ACCEPTED  statusEnum = "accepted"
	REJECTED  statusEnum = "rejected"
	CANCELLED statusEnum = "cancelled"
)

type Travel struct {
	gorm.Model
	Id              int32      `json:"id" binding:"required" gorm:"primaryKey"`
	Seat            int32      `json:"seat" binding:"required" gorm:"primaryKey"`
	Driver          *User      `json:"driver" binding:"required" gorm:"foreignKey:Id;references:Driver"`
	Passenger       *User      `json:"passenger" binding:"required" gorm:"foreignKey:Id;references:Passenger"`
	Start           *Locations `json:"start" binding:"required" gorm:"foreignKey:Id;references:Start"`
	End             *Locations `json:"end" binding:"required" gorm:"foreignKey:Id;references:End"`
	When            time.Time  `json:"when" binding:"required"`
	Status          statusEnum `json:"status" binding:"required"`
	DriverRating    int32      `json:"driverRating"`
	PassengerRating int32      `json:"passengerRating"`
}
