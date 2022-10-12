package controllers

import "github.com/courselab/pollex/pollex-backend/pkg/domain"

func GetUsers() []domain.User {

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

func GetUser(userId int32) (domain.User, error) {
	// get user from database
	var user domain.User
	user.Id = userId

	return user, nil
}

func CreateUser(user domain.User) (domain.User, error) {
	//create user in the database

	return user, nil
}

func UpdateUser(userId int32, user domain.User) (domain.User, error) {
	//update user in the database

	return user, nil
}

func DeleteUser(userId int32) error {
	//delete user in the database

	return nil
}

func PatchUser(userId int32, user domain.User) (domain.User, error) {
	// update some user attributes in the database

	return user, nil
}
