package handler

import (
	"net/http"

	"github.com/Udehlee/Task-Management/pkg/service"
	"github.com/Udehlee/Task-Management/pkg/store"
)

type Handler struct {
	Service service.Service
}

func NewHandler(datab store.PgConn) Handler {

	return Handler{
		Service: service.Service{
			Store: datab,
		},
	}
}

func (h Handler) Index(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "applicatin/json")
	// w.WriteHeader(http.StatusOK)
	// w.Write([]byte("Hello!"))
	w.Write([]byte("Welcome to the Index page"))

}
