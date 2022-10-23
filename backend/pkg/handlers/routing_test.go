package handlers

/*
func TestRounting(t *testing.T) {
	router := SetRoutes()

	t.Run("ping endpoint success", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/ping", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "pong", w.Body.String())
	})

	t.Run("get all users endpoint success", func(t *testing.T) {
		expectedResponse := []domain.User{
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

		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/users", nil)
		router.ServeHTTP(w, req)

		var response []domain.User
		json.Unmarshal(w.Body.Bytes(), &response)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, expectedResponse, response)
	})

	t.Run("get one user endpoint success", func(t *testing.T) {
		expectedResponse := domain.User{
			Id:          25,
			Name:        "",
			Nickname:    "",
			IsDriver:    false,
			DriverStats: nil,
			PassengerStats: domain.PassengerStats{
				RatingAvg:   0,
				RatingCount: 0,
				TripCount:   0,
			},
			Car: nil,
		}

		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/users/25", nil)
		router.ServeHTTP(w, req)

		var response domain.User
		json.Unmarshal(w.Body.Bytes(), &response)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, expectedResponse, response)
	})

}
*/
