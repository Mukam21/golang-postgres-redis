package database

import (
	"context"
	"fmt"
	"golang_postgresql_redis/pkg/handlers"
	"golang_postgresql_redis/pkg/models"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
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

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("failed to migrate database:", err)
	}

	redisAddr := os.Getenv("REDIS_ADDR")
	if redisAddr == "" {
		redisAddr = "localhost:6379"
	}

	rdb := redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})

	r := chi.NewRouter()
	r.Post("/users", handlers.CreateUserHandler(db))
	r.Get("/users/{id}", handlers.GetUserHandler(db, rdb, ctx))

	fmt.Println("Server is running on :8080")
	err = http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}
}
