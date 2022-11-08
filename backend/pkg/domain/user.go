package domain

import "gorm.io/gorm"

type PassengerStats struct {
	RatingAvg   int32 `json:"ratingAvg" binding:"required"`
	RatingCount int32 `json:"ratingCount" binding:"required"`
	TripCount   int32 `json:"tripCount" binding:"required"`
}

type DriverStats struct {
	RatingAvg   int32 `json:"ratingAvg" binding:"required"`
	RatingCount int32 `json:"ratingCount" binding:"required"`
	TripCount   int32 `json:"tripCount" binding:"required"`
}

type User struct {
	gorm.Model
	ID             string          `json:"id" binding:"required" gorm:"primaryKey"`
	Name           string          `json:"name" binding:"required"`
	Nickname       string          `json:"nickname"`
	IsDriver       bool            `json:"isDriver" gorm:"column:isDriver"`
	DriverStats    *DriverStats    `json:"driverStats" gorm:"embedded;embeddedPrefix:driver_"`
	PassengerStats *PassengerStats `json:"passengerStats" binding:"required" gorm:"embedded;embeddedPrefix:passenger_"`
	Car            *Car            `json:"car" gorm:"embedded;embeddedPrefix:car_"`
}
