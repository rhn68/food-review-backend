package handlers

import (
	"database/sql"
	"backend-fp-alpro/models"
	"github.com/gin-gonic/gin"
)

func AddFavorite(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		restaurantID := c.Param("restaurant_id")
		userID, exists := c.Get("user_id")
		if !exists {
			c.JSON(401, gin.H{
				"message": "User not found",
			})
			return
		}
		userIDInt := userID.(int)
		var restaurantExists int
		err := db.QueryRow(
			"SELECT id FROM restaurants WHERE id = $1",
			restaurantID,
		).Scan(&restaurantExists)
		if err != nil {
			c.JSON(404, gin.H{
				"message": "Restaurant not found",
			})
			return
		}
		var favoriteExists int
		err = db.QueryRow(
			`SELECT id FROM favorites
			 WHERE user_id = $1 AND restaurant_id = $2`,
			userIDInt,
			restaurantID,
		).Scan(&favoriteExists)
		if err == nil {
			c.JSON(400, gin.H{
				"message": "Restaurant already in favorites",
			})
			return
		}
		_, err = db.Exec(
			`INSERT INTO favorites (user_id, restaurant_id)
			 VALUES ($1, $2)`,
			userIDInt,
			restaurantID,
		)
		if err != nil {
			c.JSON(500, gin.H{
				"message": "Failed to add favorite",
			})
			return
		}
		c.JSON(201, gin.H{
			"message": "Restaurant added to favorites",
		})
	}
}

func GetFavorites(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, exists := c.Get("user_id")
		if !exists {
			c.JSON(401, gin.H{
				"message": "User not found",
			})
			return
		}
		userIDInt := userID.(int)
		rows, err := db.Query(
			`SELECT r.id, r.name,
					COALESCE(r.description, ''),
					r.location,
					COALESCE(r.category, ''),
					COALESCE(r.latitude, 0),
					COALESCE(r.longitude, 0),
					COALESCE(r.image, ''),
					COALESCE(AVG(rv.rating), 0)
			FROM favorites f
			JOIN restaurants r ON f.restaurant_id = r.id
			LEFT JOIN reviews rv ON r.id = rv.restaurant_id
			WHERE f.user_id = $1
			GROUP BY r.id
			ORDER BY r.id`,
			userIDInt,
		)
		if err != nil {
			c.JSON(500, gin.H{
				"message": "Failed to get favorites",
			})
			return
		}
		defer rows.Close()
		restaurants := []models.Restaurant{}
		for rows.Next() {
			var restaurant models.Restaurant
			err := rows.Scan(
				&restaurant.ID,
				&restaurant.Name,
				&restaurant.Description,
				&restaurant.Location,
				&restaurant.Category,
				&restaurant.Latitude,
				&restaurant.Longitude,
				&restaurant.Image,
				&restaurant.RatingRata2,
			)
			if err != nil {
				c.JSON(500, gin.H{
					"message": "Failed to read favorite data",
				})
				return
			}
			restaurants = append(restaurants, restaurant)
		}
		c.JSON(200, restaurants)
	}
}

func DeleteFavorite(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		restaurantID := c.Param("restaurant_id")
		userID, exists := c.Get("user_id")
		if !exists {
			c.JSON(401, gin.H{
				"message": "User not found",
			})
			return
		}
		userIDInt := userID.(int)
		result, err := db.Exec(
			`DELETE FROM favorites
			 WHERE user_id = $1 AND restaurant_id = $2`,
			userIDInt,
			restaurantID,
		)
		if err != nil {
			c.JSON(500, gin.H{
				"message": "Failed to delete favorite",
			})
			return
		}
		rowsAffected, err := result.RowsAffected()
		if err != nil {
			c.JSON(500, gin.H{
				"message": "Failed to check delete",
			})
			return
		}
		if rowsAffected == 0 {
			c.JSON(404, gin.H{
				"message": "Favorite not found",
			})
			return
		}
		c.JSON(200, gin.H{
			"message": "Restaurant removed from favorites",
		})
	}
}