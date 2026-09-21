package handlers

import (
"os"
"strings"
"time"
"github.com/gin-gonic/gin"
"github.com/golang-jwt/jwt/v5"
"github.com/joho/godotenv"

)

var jwtKey []byte

func init() {
err := godotenv.Load()
if err != nil {
panic("Failed to load .env")
}

jwtSecret := os.Getenv("JWT_SECRET")
if jwtSecret == "" {
	panic("JWT_SECRET is not set")
}

jwtKey = []byte(jwtSecret)

}

func GenerateToken(userID int, email string) (string, error) {
claims := jwt.MapClaims{
"user_id": userID,
"email": email,
"exp": time.Now().Add(24 * time.Hour).Unix(),
}

token := jwt.NewWithClaims(
	jwt.SigningMethodHS256,
	claims,
)

return token.SignedString(jwtKey)

}

func AuthMiddleware() gin.HandlerFunc {
return func(c *gin.Context) {
authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(401, gin.H{
			"message": "Authorization token is required",
		})
		c.Abort()
		return
	}
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	token, err := jwt.Parse(
		tokenString,
		func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		},
	)
	if err != nil || !token.Valid {
		c.JSON(401, gin.H{
			"message": "Invalid or expired token",
		})
		c.Abort()
		return
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.JSON(401, gin.H{
			"message": "Invalid token claims",
		})
		c.Abort()
		return
	}
	userID := int(claims["user_id"].(float64))
	c.Set("user_id", userID)
	c.Next()
}
}