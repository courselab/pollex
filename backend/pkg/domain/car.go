package domain

type Car struct {
	Model        string `json:"model" binding:"required"`
	Color        string `json:"color" binding:"required"`
	LicensePlate string `json:"licensePlate" binding:"required"`
}
