package kafka

import (
	"alefbookstores-search-service/openlibrary"
	"net/http"
	"sync"
)

var httpClients sync.Pool

func init() {
	httpClients = sync.Pool{
		New: func() any {
			return &http.Client{}
		},
	}
}

type BooksSearchResultsProducer interface {
	ProduceBooksSearchResults(foundBooks []Book) error
	DoBooksSearchByTitle(searchText string) ([]Book, error)
	DoBooksSearchByAuthor(authorName string) ([]Book, error)
}

type AuthorsSearchResultsProducer interface {
	ProduceAuthorsSearchResults(foundAuthors []openlibrary.Author) error
	DoAuthorsSearch(searchText string) ([]openlibrary.Author, error)
}

type Book struct {
}
