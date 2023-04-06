package routers

import (
	"alefbookstores-gateway/clients"
	"alefbookstores-gateway/models"
	"context"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

// Router is equivalent to a controller is Spring Boot world
// Important: Make sure that the dependencies and imports throughout the project are chi's version 5
type searchRouter struct {
	chi.Router
	client clients.SearchClient
}

// NewSearchRouter
// / We will be registering the http url paths with a corresponding handle function that is implemented by the router
func NewSearchRouter() chi.Router {

	/*
	 * The previous problem was that the embeddable router was not directly linked to the searchRouter.
	 * The original issue was attaching extra logic to the router i.e. the logic the service, similar to what controllers
	 * are in Spring Boot, where they have dependency of the services.
	 * So the most appropriate solution that I have found was to create a chi.Router as an embeddable interface to my
	 * custom router that will have the search and quotes server (and later on other services) as dependencies.
	 */
	embeddableRouter := chi.NewRouter()
	searchRouter := searchRouter{embeddableRouter, clients.SearchClient{}}

	embeddableRouter.Get("/books/title/{book_title}", searchRouter.SearchForBookByTitleHandler)
	embeddableRouter.Get("/books/author/{author_name}", searchRouter.SearchForBookByAuthorHandler)
	embeddableRouter.Get("/authors/{author_name}", searchRouter.SearchForAuthorHandler())

	return searchRouter
}

func (s searchRouter) SearchForBookByTitleHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	bookTitle := chi.URLParam(r, "book_title")
	log.Printf("searching for books with title %s\n", bookTitle)

	ctx := context.TODO()
	books, err := s.client.SearchForBookByTitle(ctx, bookTitle)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		reqErr, _ := json.Marshal(models.RequestError("failed to fetch books by author from search service " + err.Error()))
		w.Write(reqErr)
		return
	}

	booksInBytes, err := json.Marshal(books)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(models.JsonParsingError(books, err))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(booksInBytes)
}

func (s searchRouter) SearchForBookByAuthorHandler(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()

	authorName := chi.URLParam(r, "author_name")
	log.Println("searching for books written by " + authorName)

	ctx := context.TODO()
	books, err := s.client.SearchForBookByAuthor(ctx, authorName)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		reqErr, _ := json.Marshal(models.RequestError("failed to fetch books by author from search service " + err.Error()))
		w.Write(reqErr)
	}

	booksInBytes, err := json.Marshal(books)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(models.JsonParsingError(books, err))
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(booksInBytes)

}

func (s searchRouter) SearchForAuthorHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		authorName := chi.URLParam(r, "author_name")

		ctx := context.TODO()
		authors, err := s.client.SearchForAuthor(ctx, authorName)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			reqErr, _ := json.Marshal(models.RequestError("failed to fetch authors from search service " + err.Error()))
			w.Write(reqErr)
		}

		authorsInBytes, err := json.Marshal(authors)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(models.JsonParsingError(authors, err))
		}

		w.WriteHeader(http.StatusOK)
		w.Write(authorsInBytes)
	}
}
