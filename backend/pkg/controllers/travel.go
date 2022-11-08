package controllers

import "github.com/courselab/pollex/pollex-backend/pkg/domain"

type Travel interface {
	GetTravels() []domain.Travel
	GetTravel(travelId int32) (*domain.Travel, error)
	CreateTravel(travel domain.Travel) (*domain.Travel, error)
	UpdateTravel(travelId int32, travel domain.Travel) (*domain.Travel, error)
	DeleteTravel(travelId int32) error
	PatchTravel(travelId int32, travel domain.Travel) (*domain.Travel, error)
}

type travel struct {
}

type Params struct {
}

func NewTravelController(p *Params) Travel {
	return &travel{}
}

func (t *travel) getTravels() []domain.Travel {
	//TODO: get travels from database
	var travels []domain.Travel

	return travels
}

func (t *travel) GetTravel(travelId int32) (*domain.Travel, error) {
	//TODO: get travel from database

	var travel domain.Travel
	travel.Id = travelId

	return &travel, nil
}

func (t *travel) CreateTravel(travel domain.Travel) (*domain.Travel, error) {
	//TODO: create travel in the database

	return &travel, nil
}

func (t *travel) UpdateTravel(travelId int32, travel domain.Travel) (*domain.Travel, error) {
	//TODO: update travel in the database

	return &travel, nil
}

func (t *travel) DeleteTravel(travelId int32) error {
	//TODO: delete travel in the database

	return nil
}

func (t *travel) PatchTravel(travelId int32, travel domain.Travel) (*domain.Travel, error) {
	//TODO: update some travel attributes in the database

	return &travel, nil
}
