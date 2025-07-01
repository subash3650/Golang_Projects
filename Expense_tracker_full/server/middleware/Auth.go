package middleware

import (
	"context"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

func AuthMiddleware(Collection *mongo.Collection) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct {
			Email           string `json:"email"`
			Password        string `json:"password"`
			ConfirmPassword string `json:"confirm_password"`
		}
		if err := c.ShouldBindJSON(&req); err != nil || req.Email == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid User"})
			return
		}
		if !emailRegex.MatchString(req.Email) {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid email format"})
			return
		}
		count, err := Collection.CountDocuments(context.Background(), bson.M{"email": req.Email})
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
			return
		}
		if count > 0 {
			c.AbortWithStatusJSON(http.StatusConflict, gin.H{"message": "User already exists"})
			return
		}
		if req.Password != req.ConfirmPassword {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Passwords do not match"})
			return
		}
		if len(req.Password) < 8 {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Password must be at least 8 characters long"})
			return
		}
		if len(req.Password) > 20 {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Password must not exceed 20 characters"})
			return
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Error hashing password"})
			return
		}
		c.Set("email", req.Email)
		c.Set("password", string(hashedPassword))
		c.Set("collection", Collection)
		c.Next()
	}
}
