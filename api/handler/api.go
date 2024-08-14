package handler

import (
	"net/http"

	"github.com/Udehlee/Task-Management/pkg/service"
	"github.com/Udehlee/Task-Management/pkg/store"
)

package handler

import (
	"net/http"

	"github.com/Udehlee/Task-Management/pkg/service"
	"github.com/Udehlee/Task-Management/pkg/store"
)

type Handler struct {
	Service *service.Service
}

func NewHandler(db store.PgConn) *Handler {
	svc := service.NewService(db)
	return &Handler{
		Service: svc,
	}
}

func (h Handler) Index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome Home"))

}


func (h Handler) Index(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Welcome Home"))
}
