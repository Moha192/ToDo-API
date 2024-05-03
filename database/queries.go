package database

import (
	"database/sql"
	"errors"
	"log"

	"github.com/Moha192/ToDo-App/models"
)

// TASK
func GetTasks(task *models.Task) ([]models.Task, error) {
	rows, err := DB.Query("SELECT taskID, userID, task, status FROM tasks WHERE userID = $1 ORDER BY taskID ASC", task.UserID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []models.Task

	for rows.Next() {
		var task models.Task

		err := rows.Scan(&task.TaskID, &task.UserID, &task.Task, &task.Status)
		if err != nil {
			log.Fatal(err)
		}

		tasks = append(tasks, task)
	}
	return tasks, nil
}

func PostTask(task *models.Task) error {
	err := DB.QueryRow("INSERT INTO tasks (userID, task) VALUES ($1, $2) RETURNING taskID, userID, task, status", task.UserID, task.Task).Scan(&task.TaskID, &task.UserID, &task.Task, &task.Status)
	if err != nil {
		return err
	}
	return nil
}

func PatchTask(task *models.Task) error {
	err := DB.QueryRow("UPDATE tasks SET task = $1 WHERE taskID = $2 RETURNING taskID, userID, task, status", task.Task, task.TaskID).Scan(&task.TaskID, &task.UserID, &task.Task, &task.Status)
	if err != nil {
		return err
	}
	return nil
}

func PatchStatus(task *models.Task) error {
	err := DB.QueryRow("UPDATE tasks SET status = NOT status WHERE taskID = $1 RETURNING taskID, userID, task, status", task.TaskID).Scan(&task.TaskID, &task.UserID, &task.Task, &task.Status)
	if err != nil {
		return err
	}
	return nil
}

func DeleteTask(task *models.Task) error {
	result, err := DB.Exec("DELETE FROM tasks WHERE taskID = $1", task.TaskID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err

	} else if rowsAffected == 0 {
		return errors.New("taskID does not exist")
	}
	return nil
}

// USER
func SignUp(user *models.User) error {
	emailTaken, err := isEmailTaken(user.Email)
	if err != nil {
		return err
	} else if emailTaken {
		return errors.New("email is taken")
	}

	err = DB.QueryRow("INSERT INTO users(email, password) VALUES($1, $2) RETURNING userid, email, password", user.Email, user.Password).Scan(&user.UserID, &user.Email, &user.Password)
	if err != nil {
		return err
	}
	return nil
}

func isEmailTaken(email string) (bool, error) {
	var scan int
	err := DB.QueryRow("SELECT userid FROM users WHERE email = $1", email).Scan(&scan)
	if err == sql.ErrNoRows {
		return false, nil
	} else if err != nil {
		return true, err
	}
	return true, nil
}

func LogIn(user *models.User) error {
	err := DB.QueryRow("SELECT userID FROM users WHERE email = $1 AND password = $2", user.Email, user.Password).Scan(&user.UserID)
	switch {
	case err == sql.ErrNoRows:
		return err
	case err != nil:
		return err
	}
	return nil
}

func DeleteUser(user *models.User) error {
	_, err := DB.Exec("DELETE FROM tasks WHERE userID = $1", user.UserID)
	if err != nil {
		return err
	}

	result, err := DB.Exec("DELETE FROM users WHERE userID = $1", user.UserID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err

	} else if rowsAffected == 0 {
		return errors.New("userID does not exist")
	}
	return nil
}
