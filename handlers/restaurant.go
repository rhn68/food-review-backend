package handlers

import (
	"database/sql"
	"strconv"
	"github.com/gin-gonic/gin"
	"backend-fp-alpro/models"
)

func GetRestaurants(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		rows, err := db.Query(
			`SELECT r.id, r.name,
					COALESCE(r.description, ''),
					r.location,
					COALESCE(r.category, ''),
					COALESCE(r.latitude, 0),
					COALESCE(r.longitude, 0),
					COALESCE(r.image, ''),
					COALESCE(AVG(rv.rating), 0)
			FROM restaurants r
			LEFT JOIN reviews rv ON r.id = rv.restaurant_id
			GROUP BY r.id
			ORDER BY r.id`,
		)
		if err != nil {
			c.JSON(500, gin.H{
				"message": "Failed to get restaurants",
			})
			return
		}
		defer rows.Close()
		var restaurants []models.Restaurant
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
				&restaurant.Rating,
			)
			if err != nil {
				c.JSON(500, gin.H{
					"message": "Failed to read restaurant data",
				})
				return
			}
			restaurants = append(restaurants, restaurant)
		}
		c.JSON(200, restaurants)
	}
}

func CreateRestaurant(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var restaurant models.Restaurant
		err := c.ShouldBindJSON(&restaurant)
		if err != nil {
			c.JSON(400, gin.H{
				"message": "Invalid JSON",
			})
			return
		}
		if restaurant.Name == "" {
			c.JSON(400, gin.H{
				"message": "Name cannot be empty",
			})
			return
		}
		if restaurant.Location == "" {
			c.JSON(400, gin.H{
				"message": "Location cannot be empty",
			})
			return
		}
		err = db.QueryRow(
			`INSERT INTO restaurants (
			name,
			description,
			location,
			category,
			latitude,
			longitude,
			image
			)
			VALUES ($1, $2, $3, $4, $5, $6, $7)
			RETURNING id`,
			restaurant.Name,
			restaurant.Description,
			restaurant.Location,
			restaurant.Category,
			restaurant.Latitude,
			restaurant.Longitude,
			restaurant.Image,
		).Scan(&restaurant.ID)
		if err != nil {
			c.JSON(500, gin.H{
				"message": "Failed to add restaurant",
			})
			return
		}
		c.JSON(201, restaurant)
	}
}

func GetRestaurantByID(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var restaurant models.Restaurant
		err := db.QueryRow(
		`SELECT r.id, r.name,
				COALESCE(r.description, ''),
				r.location,
				COALESCE(r.category, ''),
				COALESCE(r.latitude, 0),
				COALESCE(r.longitude, 0),
				COALESCE(r.image, ''),
				COALESCE(AVG(rv.rating), 0)
		FROM restaurants r
		LEFT JOIN reviews rv ON r.id = rv.restaurant_id
		WHERE r.id = $1
		GROUP BY r.id`,
			id,
		).Scan(
			&restaurant.ID,
			&restaurant.Name,
			&restaurant.Description,
			&restaurant.Location,
			&restaurant.Category,
			&restaurant.Latitude,
			&restaurant.Longitude,
			&restaurant.Image,
			&restaurant.Rating,
		)
		if err != nil {
			c.JSON(404, gin.H{
				"message": "Restaurant not found",
			})
			return
		}
		c.JSON(200, restaurant)
	}
}

func UpdateRestaurant(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var restaurant models.Restaurant
		err := c.ShouldBindJSON(&restaurant)
		if err != nil {
			c.JSON(400, gin.H{
				"message": "Invalid JSON",
			})
			return
		}
		if restaurant.Name == "" {
			c.JSON(400, gin.H{
				"message": "Name cannot be empty",
			})
			return
		}
		if restaurant.Location == "" {
			c.JSON(400, gin.H{
				"message": "Location cannot be empty",
			})
			return
		}
		result, err := db.Exec(
			`UPDATE restaurants
			SET name = $1,
				description = $2,
				location = $3,
				category = $4,
				latitude = $5,
				longitude = $6,
				image = $7
			WHERE id = $8`,
			restaurant.Name,
			restaurant.Description,
			restaurant.Location,
			restaurant.Category,
			restaurant.Latitude,
			restaurant.Longitude,
			restaurant.Image,
			id,
		)
		if err != nil {
			c.JSON(500, gin.H{
				"message": "Failed to update restaurant",
			})
			return
		}
		rowsAffected, err := result.RowsAffected()
		if err != nil {
			c.JSON(500, gin.H{
				"message": "Failed to check update",
			})
			return
		}
		if rowsAffected == 0 {
			c.JSON(404, gin.H{
				"message": "Restaurant not found",
			})
			return
		}
		restaurant.ID, _ = strconv.Atoi(id)
		c.JSON(200, restaurant)
	}
}

func DeleteRestaurant(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		result, err := db.Exec(
			"DELETE FROM restaurants WHERE id = $1",
			id,
		)
		if err != nil {
			c.JSON(500, gin.H{
				"message": "Failed to delete restaurant",
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
				"message": "Restaurant not found",
			})
			return
		}
		c.JSON(200, gin.H{
			"message": "Restaurant deleted",
		})
	}
}

func SearchRestaurants(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		name := c.Query("name")
		if name == "" {
			c.JSON(400, gin.H{
				"message": "Search name cannot be empty",
			})
			return
		}
		rows, err := db.Query(
			`SELECT r.id, r.name,
					COALESCE(r.description, ''),
					r.location,
					COALESCE(r.category, ''),
					COALESCE(r.latitude, 0),
					COALESCE(r.longitude, 0),
					COALESCE(r.image, ''),
					COALESCE(AVG(rv.rating), 0)
			FROM restaurants r
			LEFT JOIN reviews rv ON r.id = rv.restaurant_id
			WHERE r.name ILIKE $1
			GROUP BY r.id
			ORDER BY r.id`,
			"%"+name+"%",
		)
		if err != nil {
			c.JSON(500, gin.H{
				"message": "Failed to search restaurants",
			})
			return
		}
		defer rows.Close()
		var restaurants []models.Restaurant
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
				&restaurant.Rating,
			)
			if err != nil {
				c.JSON(500, gin.H{
					"message": "Failed to read restaurant data",
				})
				return
			}
			restaurants = append(restaurants, restaurant)
		}
		c.JSON(200, restaurants)
	}
}