package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/Moha192/ToDo-App/database"
	"github.com/Moha192/ToDo-App/models"
)

var task models.Task

func handleTask(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == http.MethodGet:
		getTasks(w, r)

	case r.Method == http.MethodPost:
		postTask(w, r)

	case r.Method == http.MethodDelete:
		deleteTask(w, r)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func getTasks(w http.ResponseWriter, r *http.Request) {
	var err error
	task.UserID, err = strconv.Atoi(r.URL.Query().Get("userID"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return

	} else if task.UserID <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	resp, err := database.GetTasks(&task)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(resp)
}

func postTask(w http.ResponseWriter, r *http.Request) {
	if err := decoder(w, r, &task); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return

	} else if task.UserID <= 0 || task.Task == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := database.PostTask(&task)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(task)
}

func handlePatchTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	if err := decoder(w, r, &task); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return

	} else if task.TaskID <= 0 || task.Task == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := database.PatchTask(&task)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(task)
}

func handlePatchStatus(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	if err := decoder(w, r, &task); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return

	} else if task.TaskID <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := database.PatchStatus(&task)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(task)
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	if err := decoder(w, r, &task); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return

	} else if task.TaskID <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := database.DeleteTask(&task)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
