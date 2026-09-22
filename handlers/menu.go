package handlers

import (
	"backend-fp-alpro/models"
	"database/sql"
	"github.com/gin-gonic/gin"
)

func GetMenus(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		restaurantID := c.Param("id")
		rows, err := db.Query(
			`SELECT id, restaurant_id, nama, harga, COALESCE(badge, '')
			 FROM menus
			 WHERE restaurant_id = $1
			 ORDER BY id`,
			restaurantID,
		)
		if err != nil {
			c.JSON(500, gin.H{
				"message": "Failed to get menus",
			})
			return
		}
		defer rows.Close()
		var menus []models.Menu
		for rows.Next() {
			var menu models.Menu
			err := rows.Scan(
				&menu.ID,
				&menu.RestaurantID,
				&menu.Nama,
				&menu.Harga,
				&menu.Badge,
			)
			if err != nil {
				c.JSON(500, gin.H{
					"message": "Failed to read menu data",
				})
				return
			}
			menus = append(menus, menu)
		}
		if menus == nil {
			menus = []models.Menu{}
		}
		c.JSON(200, menus)
	}
}

func CreateMenu(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		restaurantID := c.Param("id")
		var menu models.Menu
		err := c.ShouldBindJSON(&menu)
		if err != nil {
			c.JSON(400, gin.H{
				"message": "Invalid JSON",
			})
			return
		}
		if menu.Nama == "" {
			c.JSON(400, gin.H{
				"message": "Menu name cannot be empty",
			})
			return
		}
		if menu.Harga < 0 {
			c.JSON(400, gin.H{
				"message": "Menu price cannot be negative",
			})
			return
		}
		var restaurantExists int
		err = db.QueryRow(
			"SELECT id FROM restaurants WHERE id = $1",
			restaurantID,
		).Scan(&restaurantExists)
		if err != nil {
			c.JSON(404, gin.H{
				"message": "Restaurant not found",
			})
			return
		}
		err = db.QueryRow(
			`INSERT INTO menus (
				restaurant_id,
				nama,
				harga,
				badge
			)
			VALUES ($1, $2, $3, $4)
			RETURNING id`,
			restaurantID,
			menu.Nama,
			menu.Harga,
			menu.Badge,
		).Scan(&menu.ID)
		if err != nil {
			c.JSON(500, gin.H{
				"message": "Failed to add menu",
			})
			return
		}
		menu.RestaurantID = restaurantExists
		c.JSON(201, menu)
	}
}
