package routers

import (
	"alefbookstores-gateway/clients"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type readersRouter struct {
	chi.Router
	client clients.ReadersClient
}

func NewReadersRouter() chi.Router {

	embeddableRouter := chi.NewRouter()

	r := readersRouter{Router: embeddableRouter}

	embeddableRouter.Post("/", r.Register())
	embeddableRouter.Post("/login", r.Login())
	embeddableRouter.Get("/", r.GetAllUsers())
	embeddableRouter.Get("/{email}", r.GetUserByEmail())

	return r
}

func (r readersRouter) Register() http.HandlerFunc {
	return func(writer http.ResponseWriter, req *http.Request) {

	}
}

func (r readersRouter) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

	}
}

func (r readersRouter) GetAllUsers() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

	}
}

func (r readersRouter) GetUserByEmail() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

	}
}
