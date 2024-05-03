package models

type User struct {
	UserID   int    `json:"userID"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Task struct {
	TaskID int    `json:"taskID"`
	UserID int    `json:"userID"`
	Task   string `json:"task"`
	Status bool   `json:"status"`
}
