package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

func InitializeRoutes() {
	//user
	http.HandleFunc("/user/signUp", handleSignUp)
	http.HandleFunc("/user/logIn", handleLogIn)
	http.HandleFunc("/user", handleDeleteUser)

	//task
	http.HandleFunc("/task", handleTask)
	http.HandleFunc("/task/text", handlePatchTask)
	http.HandleFunc("/task/status", handlePatchStatus)
}

func CorsHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Content-Type", "application/json")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func decoder(w http.ResponseWriter, r *http.Request, s interface{}) error {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&s)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return err
	}
	return nil
}
