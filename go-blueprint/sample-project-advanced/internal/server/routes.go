package server

import (
	"encoding/json"
	"log"
	"net/http"

	"sample-project-advanced/cmd/web"
	"sample-project-advanced/types"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/google/uuid"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", s.HelloWorldHandler)

	r.Get("/health", s.healthHandler)

	fileServer := http.FileServer(http.FS(web.Files))
	r.Handle("/assets/*", fileServer)
	r.Get("/web", templ.Handler(web.HelloForm()).ServeHTTP)
	r.Post("/hello", web.HelloWebHandler)

	// my original Handler
	r.Get("/web/tasks", s.HandleTasks)
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

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	jsonResp, _ := json.Marshal(s.db.Health())
	_, _ = w.Write(jsonResp)
}

func (s *Server) CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var req = types.CreateTaskRequest{
		TaskName:    r.FormValue("name"),
		Description: r.FormValue("description"),
		IsCompleted: r.FormValue("is_completed") == "on",
	}

	id := uuid.New()

	query := "INSERT INTO tasks (id, name, description, is_completed) VALUES ($1, $2, $3, $4)"
	_, err = s.db.Exec(query, id, req.TaskName, req.Description, req.IsCompleted)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	s.HandleTasks(w, r)
}

func (s *Server) HandleTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := s.GetTasks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	web.CreateTaskForm(tasks).Render(r.Context(), w)
}

func (s *Server) GetTasks() ([]types.Task, error) {
	rows, err := s.db.Query("SELECT id, name, description, created_at, is_completed FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []types.Task

	for rows.Next() {
		var task types.Task
		err := rows.Scan(&task.ID, &task.Name, &task.Description, &task.CreatedAt, &task.IsCompleted)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil

}
