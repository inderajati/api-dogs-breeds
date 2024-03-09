package models

type Dogs struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	Name     string `json:"name"`
	Breed    string `json:"breed"`
	SubBreed string `json:"sub-breed"`
}

type Response struct {
	Message string `json:"message"`
}