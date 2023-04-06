package main

import (
	"alefbookstores-gateway/routers"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func main() {

	r := chi.NewRouter()
	r.Mount("/search", routers.NewSearchRouter())
	r.Mount("/quotes", routers.NewQuotesRouter())

	http.ListenAndServe("localhost:8080", r)
}
