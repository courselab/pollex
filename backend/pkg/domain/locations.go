package domain

import (
	geom "github.com/twpayne/go-geom"
	"gorm.io/gorm"
)

type Locations struct {
	gorm.Model
	Id     int32      `json:"id" binding:"required" gorm:"primaryKey"`
	Name   string     `json:"name" binding:"required"`
	Coords geom.Point `json:"coords" binding:"required"`
}
