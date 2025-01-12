package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/ShekleinAleksey/GoTasks/models"
)

type TaskHandlers struct {
	db *sql.DB
}

func NewTaskHandkers(db *sql.DB) *TaskHandlers {
	return &TaskHandlers{db: db}
}

func (h *TaskHandlers) CreateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query := `INSERT INTO tasks (title, description, status, due_date, creator_id) VALUES ($1, $2, $3, $4, $5)
	RETURNING id`
	row := h.db.QueryRow(query, task.Title, task.Description, task.Status, task.DueDate, task.CreatorId)
	err = row.Scan(&task.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func (h *TaskHandlers) GetTask(w http.ResponseWriter, r *http.Request) {
	var tasks []models.Task

	query := `SELECT * FROM tasks`
	rows, err := h.db.Query(query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var task models.Task
		err = rows.Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.DueDate, &task.CreatorId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tasks = append(tasks, task)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func (h *TaskHandlers) GetTaskById(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	// id := vars["id"]

	// query := `SELECT * FROM tasks WHERE id=$1`
}

func (h *TaskHandlers) UpdateTask(w http.ResponseWriter, r *http.Request) {

}

func (h *TaskHandlers) DeleteTask(w http.ResponseWriter, r *http.Request) {

}
