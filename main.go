package main

import (
	"fmt"
	"backend-fp-alpro/handlers"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func main() {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	db := connectDatabase()
	defer db.Close()
	fmt.Println("Database connected!")

	router.GET("/restaurants", handlers.GetRestaurants(db))
	router.GET("/restaurants/search", handlers.SearchRestaurants(db))
	router.GET("/restaurants/:id", handlers.GetRestaurantByID(db))
	router.POST("/restaurants", handlers.CreateRestaurant(db))
	router.PUT("/restaurants/:id", handlers.UpdateRestaurant(db))
	router.DELETE("/restaurants/:id", handlers.DeleteRestaurant(db))

	router.GET("/api/tempat/:id/menu", handlers.GetMenus(db)) 
	router.POST("/api/tempat/:id/menu", handlers.CreateMenu(db))

	router.GET("/reviews", handlers.GetReviews(db))
	router.POST("/reviews", handlers.AuthMiddleware(), handlers.CreateReview(db))
	router.GET("/reviews/:id", handlers.GetReviewByID(db))
	router.PUT("/reviews/:id", handlers.AuthMiddleware(), handlers.UpdateReview(db))
	router.DELETE("/reviews/:id", handlers.AuthMiddleware(), handlers.DeleteReview(db))

	router.GET("/restaurants/:id/reviews", handlers.GetReviewsByRestaurant(db))
	router.GET("/restaurants/:id/rating", handlers.GetRestaurantRating(db))

	router.POST("/register", handlers.RegisterUser(db))
	router.POST("/login", handlers.LoginUser(db))
	router.GET("/profile", handlers.AuthMiddleware(), handlers.GetProfile(db))

	router.POST("/favorites/:restaurant_id", handlers.AuthMiddleware(), handlers.AddFavorite(db))
	router.GET("/favorites", handlers.AuthMiddleware(), handlers.GetFavorites(db))
	router.DELETE("/favorites/:restaurant_id", handlers.AuthMiddleware(), handlers.DeleteFavorite(db))
	
	router.Run(":8080")
}