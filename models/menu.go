package models

type Menu struct {
	ID           int    `json:"id"`
	RestaurantID int    `json:"tempat_id"`
	Nama         string `json:"nama"`
	Harga        int    `json:"harga"`
	Badge        string `json:"badge"`
}
