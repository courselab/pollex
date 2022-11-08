package domain

import "time"

type Travel struct {
	Id              int32      `json:"id" binding:"required"`
	Seat            int32      `json:"seat" binding:"required"`
	Driver          *User      `json:"driver" binding:"required" gorm:"foreignKey:Id;references:Driver"`
	Passenger       []User     `json:"passenger" binding:"required" gorm:"foreignKey:Id;references:Passenger"`
	Start           *Locations `json:"start" binding:"required" gorm:"foreignKey:Id;references:Start"`
	End             *Locations `json:"end" binding:"required" gorm:"foreignKey:Id;references:End"`
	When            time.Time  `json:"when" binding:"required"`
	DriverRating    int32      `json:"driverRating"`
	PassengerRating int32      `json:"passengerRating"`
}
