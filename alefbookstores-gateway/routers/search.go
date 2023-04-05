package routers

import (
	"alefbookstores-gateway/clients"
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

// Router is equivalent to a controller is Spring Boot world
type searchRouter struct {
	chi.Router
	clients.SearchClient
}

func NewSearchRouter() searchRouter {
	return searchRouter{
		Router:       SearchRouter(),
		SearchClient: clients.SearchClient{},
	}
}

// We will be registering the http url paths with a corresponding handle function that is implemented by the router
func SearchRouter() searchRouter {

	searchChiRouter := chi.NewRouter()
	myRouter := NewSearchRouter()

	searchChiRouter.Get("/Book/{title}", myRouter.SearchForBookByTitleHandle)

	searchChiRouter.Get("/Book/{author_name}", myRouter.SearchForBookByAuthorHandle)

	searchChiRouter.Get("/Author/{author_name}", myRouter.SearchForAuthorHandle)

	return myRouter
}

func (s searchRouter) SearchForBookByTitleHandle(w http.ResponseWriter, r *http.Request) {
	bookTitle := chi.URLParam(r, "title")
	log.Println(bookTitle)

}

func (s searchRouter) SearchForBookByAuthorHandle(w http.ResponseWriter, r *http.Request) {
	authorName := chi.URLParam(r, "author_name")
	log.Println("Searching the books written by " + authorName)

	ctx := context.Background()

	books, err := s.SearchForBookByAuthor(ctx, authorName)
	if err != nil {

	}
}

func (s searchRouter) SearchForAuthorHandle(w http.ResponseWriter, r *http.Request) {

	authorName := chi.URLParam(r, "author_name")
	log.Println("Searching for authors with title " + authorName)

	ctx := context.Background()

	authors, err := s.SearchForAuthor(ctx, authorName)
	if err != nil {

	}

	authorsInBytes, err := json.Marshal(authors)
	if err != nil {

	}

	w.WriteHeader(http.StatusOK)
	w.Write(authorsInBytes)

}
