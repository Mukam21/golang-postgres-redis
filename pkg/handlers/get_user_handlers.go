package handlers

import (
	"context"
	"fmt"
	"golang_postgresql_redis/pkg/models"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func GetUserHandler(db *gorm.DB, rdb *redis.Client, ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		var user models.User
		val, err := rdb.Get(ctx, "user:"+id).Result()
		if err == nil {
			fmt.Fprintf(w, "From Cache: %s\n", val)
			return
		}

		if err := db.First(&user, id).Error; err != nil {
			http.Error(w, "user not found", http.StatusNotFound)
			return
		}

		rdb.Set(ctx, "user:"+id, fmt.Sprintf("%s <%s>", user.Name, user.Email), time.Minute*5)
		fmt.Fprintf(w, "From DB: %s <%s>\n", user.Name, user.Email)
	}
}
