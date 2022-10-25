package domain

import (
	"gorm.io/gorm"
	"time"
)

type Locations struct {
	ID     int32  `json:"id" binding:"required"`
	name   string `json:"name" binding:"required"`
	coords point  `json:"coords" binding:"required"`
}

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
	Driver          User       `json:"driver" binding:"required" gorm:"references:ID"`
	Passenger       User       `json:"passenger" binding:"required" gorm:"references:ID"`
	Start           Locations  `json:"start" binding:"required" gorm:"references:ID"`
	End             Locations  `json:"end" binding:"required" gorm:"references:ID"`
	When            time.Time  `json:"when" binding:"required"`
	Status          statusEnum `json:"status" binding:"required"`
	DriverRating    int32
	PassengerRating int32
}
