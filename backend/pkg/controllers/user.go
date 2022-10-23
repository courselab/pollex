package controllers

import "github.com/courselab/pollex/pollex-backend/pkg/domain"

type User interface {
	GetUsers() []domain.User
	GetUser(userId int32) (domain.User, error)
	CreateUser(user domain.User) (*domain.User, error)
	UpdateUser(userId int32, user domain.User) (domain.User, error)
	DeleteUser(userId int32) error
	PatchUser(userId int32, user domain.User) (domain.User, error)
}

type user struct {
}

type Params struct {
}

func NewUserController(p *Params) User {
	return &user{}
}

func (u *user) GetUsers() []domain.User {

	// get users from database

	users := []domain.User{
		{
			Id:          1223,
			Name:        "Igor Takeo Passenger",
			Nickname:    "igortakeo_passenger",
			IsDriver:    false,
			DriverStats: nil,
			PassengerStats: domain.PassengerStats{
				RatingAvg:   10,
				RatingCount: 10,
				TripCount:   50,
			},
			Car: nil,
		},
		{
			Id:       1224,
			Name:     "Igor Takeo Driver",
			Nickname: "igortakeo_driver",
			IsDriver: true,
			DriverStats: &domain.DriverStats{
				RatingAvg:   10,
				RatingCount: 10,
				TripCount:   50,
			},
			PassengerStats: domain.PassengerStats{
				RatingAvg:   10,
				RatingCount: 10,
				TripCount:   50,
			},
			Car: &domain.Car{
				Model:        "Volkswagen-Gol",
				Color:        "Grey",
				LicensePlate: "NAM-2876",
			},
		},
	}

	return users
}

func (u *user) GetUser(userId int32) (domain.User, error) {
	// get user from database
	var user domain.User
	user.Id = userId

	return user, nil
}

func (u *user) CreateUser(user domain.User) (*domain.User, error) {
	//create user in the database

	return &user, nil
}

func (u *user) UpdateUser(userId int32, user domain.User) (domain.User, error) {
	//update user in the database

	return user, nil
}

func (u *user) DeleteUser(userId int32) error {
	//delete user in the database

	return nil
}

func (u *user) PatchUser(userId int32, user domain.User) (domain.User, error) {
	// update some user attributes in the database

	return user, nil
}
