package handlers

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
	Id             int32          `json:"id" binding:"required"`
	Name           string         `json:"name" binding:"required"`
	Nickname       string         `json:"nickname" binding:"required"`
	IsDriver       bool           `json:"isDriver" binding:"required"`
	DriverStatus   *DriverStats   `json:"driverStats"`
	PassengerStats PassengerStats `json:"passengerStats" binding:"required"`
	Car            *Car           `json:"car"`
}
