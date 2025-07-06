package handlers

import (
	"golang_postgresql_redis/pkg/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateUserHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		name := c.Query("name")
		email := c.Query("email")
		if name == "" || email == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "name and email are required"})
			return
		}

		user := models.User{Name: name, Email: email}
		db.Create(&user)
		c.JSON(http.StatusCreated, user)
	}
}
