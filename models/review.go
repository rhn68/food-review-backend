package models

type Review struct {
	ID           int     `json:"id"`
	RestaurantID int     `json:"restaurant_id"`
	Reviewer     string  `json:"reviewer"`
	Rating       float64 `json:"rating"`
	Comment      string  `json:"comment"`
}