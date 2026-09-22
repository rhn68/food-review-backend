package models

type Review struct {
	ID           int     `json:"id"`
	RestaurantID int     `json:"tempat_id"`
	Reviewer     string  `json:"nama_pengulas"`
	Rating       float64 `json:"rating"`
	Comment      string  `json:"komentar"`
	CreatedAt    string  `json:"created_at"`
	Image        string  `json:"foto_url"`
}
