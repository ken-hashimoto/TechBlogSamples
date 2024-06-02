package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", s.HelloWorldHandler)

	r.Get("/tasks", s.GetTaskHandler)
	r.Get("/health", s.healthHandler)

	return r
}

func (s *Server) HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}

type Task struct {
	TaskID      int    `json:"task_id"`
	TaskName    string `json:"task_name"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	IsCompleted bool   `json:"is_completed"`
}

func (s *Server) GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := s.db.Query("SELECT task_id, task_name, description, created_at, is_completed FROM tasks")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var tasks []Task

	for rows.Next() {
		var task Task
		err := rows.Scan(&task.TaskID, &task.TaskName, &task.Description, &task.CreatedAt, &task.IsCompleted)
		if err != nil {
			log.Fatal(err)
		}
		tasks = append(tasks, task)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	jsonResp, _ := json.Marshal(s.db.Health())
	_, _ = w.Write(jsonResp)
}
