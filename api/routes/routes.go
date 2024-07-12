package routes

import (
	"net/http"

	"github.com/Udehlee/Task-Management/api/handler"
)

// func SetupRoutes(h api.Handler) *http.ServeMux {

// 	mux := http.NewServeMux()

// 	//User
// 	mux.HandleFunc("/index", h.Index)
// 	mux.HandleFunc("POST /auth/signup", h.Signup)
// 	mux.HandleFunc("POST /auth/login", h.Login)

// 	//validate
// 	// mux.HandleFunc()

// 	//user api

// 	mux.HandleFunc("GET /api/users", h.GetAllUser)
// 	mux.HandleFunc("GET /api/users/{id}", h.GetUserById)

// 	//task api

// 	// mux.HandleFunc("POST /api/tasks/:id", h.UpdateUserTask)
// 	mux.HandleFunc("POST /api/tasks/{id}", h.AddUserTask)
// 	mux.HandleFunc("POST /api/tasks/update/{id}", h.UpdateUserTask)
// 	//
// 	return mux
// }

func SetupRoutes(h handler.Handler) *http.ServeMux {

	mux := http.NewServeMux()

	mux.HandleFunc("/", h.Index)
	mux.HandleFunc("POST /auth/signup", h.Signup)
	mux.HandleFunc("POST /auth/login", h.Login)

	mux.HandleFunc("GET /api/users", h.GetAllUser)
	mux.HandleFunc("GET /api/users/{id}", h.GetUserById)

	mux.HandleFunc("POST /api/tasks/{id}", h.AddUserTask)
	mux.HandleFunc("POST /api/tasks/update/{id}", h.UpdateUserTask)

	return mux
}
