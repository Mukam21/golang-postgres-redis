package handlers

import (
	"context"
	"fmt"
	"golang_postgresql_redis/pkg/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func GetUserHandler(db *gorm.DB, rdb *redis.Client, ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		val, err := rdb.Get(ctx, "user:"+id).Result()
		if err == nil {
			c.String(http.StatusOK, "From Cache: %s", val)
			return
		}

		var user models.User
		if err := db.First(&user, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
			return
		}

		rdb.Set(ctx, "user:"+id, fmt.Sprintf("%s <%s>", user.Name, user.Email), time.Minute*5)
		c.String(http.StatusOK, "From DB: %s <%s>", user.Name, user.Email)
	}
}
