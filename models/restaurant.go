package models

type Restaurant struct {
    ID          int     `json:"id"`
    Name        string  `json:"name"`
    Description string  `json:"description"`
    Location    string  `json:"location"`
    Category    string  `json:"category"`
    Latitude    float64 `json:"latitude"`
    Longitude   float64 `json:"longitude"`
    Image       string  `json:"image"`
    Rating      float64 `json:"rating"`
}