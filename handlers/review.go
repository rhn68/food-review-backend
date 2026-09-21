package handlers

import (
	"database/sql"
	"strconv"
	"backend-fp-alpro/models"
	"github.com/gin-gonic/gin"
)

func GetReviews(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		rows, err := db.Query(
			"SELECT id, restaurant_id, reviewer, rating, comment FROM reviews",
		)
		if err != nil {
			c.JSON(500, gin.H{
				"message": "Failed to get reviews",
			})
			return
		}
		defer rows.Close()
		var reviews []models.Review
		for rows.Next() {
			var review models.Review
			err := rows.Scan(
				&review.ID,
				&review.RestaurantID,
				&review.Reviewer,
				&review.Rating,
				&review.Comment,
			)
			if err != nil {
				c.JSON(500, gin.H{
					"message": "Failed to read review data",
				})
				return
			}
			reviews = append(reviews, review)
		}
		c.JSON(200, reviews)
	}
}

func CreateReview(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var review models.Review
		err := c.ShouldBindJSON(&review)
		if err != nil {
			c.JSON(400, gin.H{
				"message": "Invalid JSON",
			})
			return
		}
		userID, exists := c.Get("user_id")
		if !exists {
			c.JSON(401, gin.H{
				"message": "User not found",
			})
			return
		}
		userIDInt := userID.(int)
		var reviewer string
		err = db.QueryRow(
			"SELECT name FROM users WHERE id = $1",
			userIDInt,
		).Scan(&reviewer)
		if err != nil {
			c.JSON(500, gin.H{
				"message": "Failed to get user",
			})
			return
		}
		var restaurantID int
		err = db.QueryRow(
			"SELECT id FROM restaurants WHERE id = $1",
			review.RestaurantID,
		).Scan(&restaurantID)
		if err != nil {
			c.JSON(404, gin.H{
				"message": "Restaurant not found",
			})
			return
		}
		if review.Rating < 1 || review.Rating > 5 {
			c.JSON(400, gin.H{
				"message": "Rating must be between 1 and 5",
			})
			return
		}
		if review.Comment == "" {
			c.JSON(400, gin.H{
				"message": "Comment cannot be empty",
			})
			return
		}
		err = db.QueryRow(
			`INSERT INTO reviews (restaurant_id, user_id, reviewer, rating, comment)
			VALUES ($1, $2, $3, $4, $5)
			RETURNING id`,
			review.RestaurantID,
			userIDInt,
			reviewer,
			review.Rating,
			review.Comment,
		).Scan(&review.ID)
		if err != nil {
			c.JSON(500, gin.H{
				"message": "Failed to add review",
			})
			return
		}
		review.Reviewer = reviewer
		c.JSON(201, review)
	}
}

func GetReviewByID(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var review models.Review
		err := db.QueryRow(
			`SELECT id, restaurant_id, reviewer, rating, comment
			 FROM reviews
			 WHERE id = $1`,
			id,
		).Scan(
			&review.ID,
			&review.RestaurantID,
			&review.Reviewer,
			&review.Rating,
			&review.Comment,
		)
		if err != nil {
			c.JSON(404, gin.H{
				"message": "Review not found",
			})
			return
		}
		c.JSON(200, review)
	}
}

func UpdateReview(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		userID, exists := c.Get("user_id")
		if !exists {
			c.JSON(401, gin.H{
				"message": "User not found",
			})
			return
		}
		userIDInt := userID.(int)
		var review models.Review
		err := c.ShouldBindJSON(&review)
		if err != nil {
			c.JSON(400, gin.H{
				"message": "Invalid JSON",
			})
			return
		}
		var reviewer string
		err = db.QueryRow(
			"SELECT name FROM users WHERE id = $1",
			userIDInt,
		).Scan(&reviewer)
		if err != nil {
			c.JSON(500, gin.H{
				"message": "Failed to get user",
			})
			return
		}
		if review.Rating < 1 || review.Rating > 5 {
			c.JSON(400, gin.H{
				"message": "Rating must be between 1 and 5",
			})
			return
		}
		if review.Comment == "" {
			c.JSON(400, gin.H{
				"message": "Comment cannot be empty",
			})
			return
		}
		var restaurantID int
		err = db.QueryRow(
			"SELECT id FROM restaurants WHERE id = $1",
			review.RestaurantID,
		).Scan(&restaurantID)
		if err != nil {
			c.JSON(404, gin.H{
				"message": "Restaurant not found",
			})
			return
		}
		result, err := db.Exec(
			`UPDATE reviews
			 SET restaurant_id = $1, reviewer = $2, rating = $3, comment = $4
			 WHERE id = $5 AND user_id = $6`,
			review.RestaurantID,
			reviewer,
			review.Rating,
			review.Comment,
			id,
			userIDInt,
		)
		if err != nil {
			c.JSON(500, gin.H{
				"message": "Failed to update review",
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
				"message": "Review not found",
			})
			return
		}
		reviewID, _ := strconv.Atoi(id)
		review.Reviewer = reviewer
		c.JSON(200, gin.H{
			"id":            reviewID,
			"restaurant_id": review.RestaurantID,
			"reviewer":      review.Reviewer,
			"rating":        review.Rating,
			"comment":       review.Comment,
		})
	}
}

func DeleteReview(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		userID, exists := c.Get("user_id")

		if !exists {
			c.JSON(401, gin.H{
				"message": "User not found",
			})
			return
		}

		userIDInt := userID.(int)
		result, err := db.Exec(
			"DELETE FROM reviews WHERE id = $1 AND user_id = $2",
			id,
			userIDInt,
		)
		if err != nil {
			c.JSON(500, gin.H{
				"message": "Failed to delete review",
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
				"message": "Review not found",
			})
			return
		}
		c.JSON(200, gin.H{
			"message": "Review deleted",
		})
	}
}

func GetReviewsByRestaurant(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		restaurantID := c.Param("id")
		rows, err := db.Query(
			`SELECT id, restaurant_id, reviewer, rating, comment
			 FROM reviews
			 WHERE restaurant_id = $1`,
			restaurantID,
		)
		if err != nil {
			c.JSON(500, gin.H{
				"message": "Failed to get reviews",
			})
			return
		}
		defer rows.Close()
		var reviews []models.Review
		for rows.Next() {
			var review models.Review
			err := rows.Scan(
				&review.ID,
				&review.RestaurantID,
				&review.Reviewer,
				&review.Rating,
				&review.Comment,
			)
			if err != nil {
				c.JSON(500, gin.H{
					"message": "Failed to read review data",
				})
				return
			}
			reviews = append(reviews, review)
		}
		c.JSON(200, reviews)
	}
}

func GetRestaurantRating(db *sql.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        restaurantID := c.Param("id")
        var averageRating float64
		err := db.QueryRow(
			`SELECT COALESCE(AVG(rating), 0)
			FROM reviews
			WHERE restaurant_id = $1`,
			restaurantID,
		).Scan(&averageRating)
        if err != nil {
            c.JSON(500, gin.H{
                "message": "Failed to get restaurant rating",
            })
            return
        }
        c.JSON(200, gin.H{
            "restaurant_id": restaurantID,
            "average_rating": averageRating,
        })
    }
}