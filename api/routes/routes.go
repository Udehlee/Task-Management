package routes

import (
	"net/http"

	"github.com/Udehlee/Task-Management/api/handler"

	"github.com/Udehlee/Task-Management/middleware"
)

func SetupRoutes(h handler.Handler) *http.ServeMux {

	mux := http.NewServeMux()

	mux.HandleFunc("/", h.Index)
	mux.HandleFunc("POST /auth/signup", h.Signup)
	mux.HandleFunc("POST /auth/login", h.Login)

	mux.Handle("GET /api/users", middleware.AuthMiddleware(http.HandlerFunc(h.GetAllUser)))
	mux.Handle("GET /api/users/{id}", middleware.AuthMiddleware(http.HandlerFunc(h.GetUserById)))

	mux.Handle("POST /api/tasks/{id}", middleware.AuthMiddleware(http.HandlerFunc(h.AddUserTask)))
	mux.Handle("POST /api/tasks/update/{id}", middleware.AuthMiddleware(http.HandlerFunc(h.UpdateUserTask)))

	return mux
}
