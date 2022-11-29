package domain

type Location struct {
	Id        	int32 `json:"id" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Coords 		string `json:"coords" binding:"required"`
	Thumbnail 	string `json:"thumbnail" binding:"required"`
}
