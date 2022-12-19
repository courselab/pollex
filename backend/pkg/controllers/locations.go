package controllers

import "github.com/courselab/pollex/pollex-backend/pkg/domain"

type Locations interface {
	GetLocations() []domain.Location
}

type locations struct {
}

type LocationsParams struct {
}

func NewLocationsController(p *LocationsParams) Locations {
	return &locations{}
}

func (l *locations) GetLocations() []domain.Location {
	locations := []domain.Location{
		{
			Id:        1,
			Name:      "LocationTest",
			Coords:    "CoordsTest",
			Thumbnail: "THumbnailTest",
		},

		{
			Id:        2,
			Name:      "LocationTest2",
			Coords:    "CoordsTest2",
			Thumbnail: "THumbnailTest2",
		},
	}

	return locations
}
