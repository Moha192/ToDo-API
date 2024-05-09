package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/Moha192/ToDo-App/database"
	"github.com/Moha192/ToDo-App/models"
)

var user models.User

func handleSignUp(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	if err := decoder(w, r, &user); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return

	} else if user.Email == "" || user.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := database.SignUp(&user)
	if err != nil {
		if err.Error() == "email is taken" {
			w.WriteHeader(http.StatusConflict)
			return
		}
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func handleLogIn(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	if err := decoder(w, r, &user); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return

	} else if user.Email == "" || user.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := database.LogIn(&user)
	if err == sql.ErrNoRows {
		w.WriteHeader(http.StatusUnauthorized)
		return
	} else if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func handleDeleteUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	if err := decoder(w, r, &user); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return

	} else if user.UserID <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := database.DeleteUser(&user)
	if err == sql.ErrNoRows {
		w.WriteHeader(http.StatusUnauthorized)
		return
	} else if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
