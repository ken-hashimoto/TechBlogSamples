package server

import (
	"encoding/json"
	"log"
	"net/http"
	"sample-project/types"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", s.HelloWorldHandler)
	r.Get("/health", s.healthHandler)

	// my original handler
	r.Get("/tasks", s.GetTaskHandler)
	r.Post("/tasks", s.CreateTaskHandler)
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

func (s *Server) GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := s.db.Query("SELECT id, name, description, created_at, is_completed FROM tasks")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	var tasks []types.Task

	for rows.Next() {
		var task types.Task
		err := rows.Scan(&task.ID, &task.Name, &task.Description, &task.CreatedAt, &task.IsCompleted)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tasks = append(tasks, task)
	}

	if err = rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func (s *Server) CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var req types.CreateTaskRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query := "INSERT INTO tasks (task_name, description, is_completed) VALUES ($1, $2, $3)"
	_, err := s.db.Exec(query, req.TaskName, req.Description, req.IsCompleted)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	jsonResp, _ := json.Marshal(s.db.Health())
	_, _ = w.Write(jsonResp)
}
