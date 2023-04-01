package kafka

import (
	"alefbookstores-search-service/google"
	"errors"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

const googleBooksSearchApiUrl = "https://www.googleapis.com/books/v1/volumes?q="
const googleBooksSearchByAuthorApiUrl = "https://www.googleapis.com/books/v1/volumes?q=inauthor:"

// GoogleApiSearchProducer Google Books API does not provide a search functionality for authors
type GoogleApiSearchProducer struct {
}

func (g *GoogleApiSearchProducer) ProduceBooksSearchResults(foundBooks []Book) error {
	//TODO the producer search result should post the results in Kafka
	panic("implement me")
}

func (g *GoogleApiSearchProducer) DoBooksSearchByTitle(searchText string) ([]Book, error) {
	httpClient := httpClients.Get().(*http.Client)
	defer httpClients.Put(httpClient)

	searchQuery := googleBooksSearchApiUrl + strings.ReplaceAll(searchText, " ", "+")

	searchRequest, _ := http.NewRequest("GET", searchQuery, nil)
	// searchRequest = newGoogleBooksSearchRequest(searchText)

	searchResponse, err := httpClient.Do(searchRequest)
	if err != nil {
		return nil, errors.New("failed to issue the request, reason " + err.Error())
	}

	defer searchResponse.Body.Close()

	searchResultInBytes, err := io.ReadAll(searchResponse.Body)
	if err != nil {
		return nil, errors.New("failed to parse the google search response into bytes, reason: " + err.Error())
	}
	// log.Print(string(searchResultInBytes))
	searchResultJsonBlob := string(searchResultInBytes)

	books := google.ParseGoogleBookSearchResult(searchResultJsonBlob)
	log.Println(len(books))
	for _, book := range books {
		log.Println(book.Id)
	}
	return nil, nil

}

func (g *GoogleApiSearchProducer) DoBooksSearchByAuthor(authorName string) ([]Book, error) {

	httpClient := httpClients.Get().(*http.Client)
	defer httpClients.Put(httpClient)

	searchQuery := googleBooksSearchByAuthorApiUrl + strings.ReplaceAll(authorName, " ", "+")

	searchRequest, _ := http.NewRequest("GET", searchQuery, nil)
	// searchRequest = newGoogleBooksSearchRequest(searchText)

	searchResponse, err := httpClient.Do(searchRequest)
	if err != nil {
		return nil, errors.New("failed to issue the request, reason " + err.Error())
	}

	defer searchResponse.Body.Close()

	searchResultInBytes, err := io.ReadAll(searchResponse.Body)
	if err != nil {
		return nil, errors.New("failed to parse the google search response into bytes, reason: " + err.Error())
	}
	searchResultJsonBlob := string(searchResultInBytes)

	books := google.ParseGoogleBookSearchResult(searchResultJsonBlob)

	for _, book := range books {
		log.Print(book.Isbn10)
	}

	return nil, nil
}

/**/
func newGoogleBooksSearchRequest(searchTerm string) *http.Request {

	searchUrlInString := googleBooksSearchApiUrl + strings.ReplaceAll(searchTerm, " ", "+")
	searchRequestHeader := http.Header{}
	searchRequestHeader.Add("Accept", "*/*")
	searchRequestHeader.Add("Accept-Encoding", "gzip,deflate,br")
	searchRequestHeader.Add("Connection", "keep-alive")

	url := &url.URL{
		Path: searchUrlInString,
	}

	return &http.Request{
		Method:     "Get",
		URL:        url,
		Proto:      "HTTP/1.1",
		ProtoMajor: 1,
		ProtoMinor: 1,
		Header:     searchRequestHeader,
		Host:       url.Host,
	}
}
