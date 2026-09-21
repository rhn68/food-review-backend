package handlers

import (
    "database/sql"
    "backend-fp-alpro/models"
    "github.com/gin-gonic/gin"
    "golang.org/x/crypto/bcrypt"
)

func RegisterUser(db *sql.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        var user models.User
        err := c.ShouldBindJSON(&user)
        if err != nil {
            c.JSON(400, gin.H{
                "message": "Invalid JSON",
            })
            return
        }
        if user.Name == "" {
            c.JSON(400, gin.H{
                "message": "Name cannot be empty",
            })
            return
        }
        if user.Email == "" {
            c.JSON(400, gin.H{
                "message": "Email cannot be empty",
            })
            return
        }
        if user.Password == "" {
            c.JSON(400, gin.H{
                "message": "Password cannot be empty",
            })
            return
        }
        var existingID int
        err = db.QueryRow(
            "SELECT id FROM users WHERE email = $1",
            user.Email,
        ).Scan(&existingID)
        if err == nil {
            c.JSON(400, gin.H{
                "message": "Email already registered",
            })
            return
        }
        hashedPassword, err := bcrypt.GenerateFromPassword(
            []byte(user.Password),
            bcrypt.DefaultCost,
        )
        if err != nil {
            c.JSON(500, gin.H{
                "message": "Failed to hash password",
            })
            return
        }
        err = db.QueryRow(
            `INSERT INTO users (name, email, password)
            VALUES ($1, $2, $3)
            RETURNING id`,
            user.Name,
            user.Email,
            string(hashedPassword),
        ).Scan(&user.ID)
        if err != nil {
            c.JSON(500, gin.H{
                "message": "Failed to register user",
            })
            return
        }
        c.JSON(201, gin.H{
            "id":    user.ID,
            "name":  user.Name,
            "email": user.Email,
        })
    }
}

func LoginUser(db *sql.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        var user models.User
        err := c.ShouldBindJSON(&user)
        if err != nil {
            c.JSON(400, gin.H{
                "message": "Invalid JSON",
            })
            return
        }
        var hashedPassword string
        err = db.QueryRow(
            `SELECT id, name, password
             FROM users
             WHERE email = $1`,
            user.Email,
        ).Scan(
            &user.ID,
            &user.Name,
            &hashedPassword,
        )
        if err != nil {
            c.JSON(401, gin.H{
                "message": "Email atau password salah",
            })
            return
        }
        err = bcrypt.CompareHashAndPassword(
            []byte(hashedPassword),
            []byte(user.Password),
        )
        if err != nil {
            c.JSON(401, gin.H{
                "message": "Email atau password salah",
            })  
            return
        }
        token, err := GenerateToken(user.ID, user.Email)
        if err != nil {
            c.JSON(500, gin.H{
                "message": "Failed to generate token",
            })
            return
        }
        c.JSON(200, gin.H{
            "message": "Login berhasil",
            "id":      user.ID,
            "name":    user.Name,
            "email":   user.Email,
            "token":   token,
        })
    }
}

func GetProfile(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, exists := c.Get("user_id")
		if !exists {
			c.JSON(401, gin.H{
				"message": "User not found",
			})
			return
		}
		userIDInt := userID.(int)
		var user models.User
		err := db.QueryRow(
			`SELECT id, name, email
			 FROM users
			 WHERE id = $1`,
			userIDInt,
		).Scan(
			&user.ID,
			&user.Name,
			&user.Email,
		)
		if err != nil {
			c.JSON(404, gin.H{
				"message": "User not found",
			})
			return
		}
		c.JSON(200, gin.H{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
		})
	}
}