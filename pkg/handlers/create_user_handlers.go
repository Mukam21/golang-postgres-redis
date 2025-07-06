package handlers

import (
	"fmt"
	"golang_postgresql_redis/pkg/models"
	"net/http"

	"gorm.io/gorm"
)

func CreateUserHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		email := r.URL.Query().Get("email")
		if name == "" || email == "" {
			http.Error(w, "name and email are required", http.StatusBadRequest)
			return
		}

		user := models.User{Name: name, Email: email}
		db.Create(&user)
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, "User created: %+v\n", user)
	}
}
