package database

import (
	"context"
	"fmt"
	"golang_postgresql_redis/pkg/handlers"
	"golang_postgresql_redis/pkg/models"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func StartServer() {
	ctx := context.Background()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Не удалось загрузить .env файл")
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("PASSWORD"),
		os.Getenv("DBNAME"),
		os.Getenv("PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}

	db.AutoMigrate(&models.User{})

	rdb := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_ADDR"),
	})

	router := gin.Default()

	router.POST("/users", handlers.CreateUserHandler(db))
	router.GET("/users/:id", handlers.GetUserHandler(db, rdb, ctx))

	fmt.Println("Server is running on :8080")
	http.ListenAndServe(":8080", router)
}
