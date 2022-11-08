package domain

type Locations struct {
	Id     int32  `json:"id" binding:"required"`
	Name   string `json:"name" binding:"required"`
	Coords Point  `json:"coords" binding:"required"`
}
