package kafka

import (
	"alefbookstores-search-service/openlibrary"
	"fmt"
	"github.com/tidwall/gjson"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

const (
	openLibraryBooksSearchAPI   string = "https://openlibrary.org/search.json?title="
	openLibraryBooksAPI         string = "https://openlibrary.org/authors/"
	openLibraryAuthorsSearchAPI string = "https://openlibrary.org/search/authors.json?q="
)

type OpenLibraryApiSearchProducer struct {
}

func (o *OpenLibraryApiSearchProducer) DoBooksSearchByTitle(searchText string) ([]Book, error) {
	//TODO implement me
	panic("implement me")
}
func (o *OpenLibraryApiSearchProducer) ProduceBooksSearchResults(foundBooks []Book) error {
	//TODO implement me
	panic("implement me")
}

/* Search for books by author */
func (o *OpenLibraryApiSearchProducer) DoBooksSearchByAuthor(authorName string) ([]Book, error) {
	//TODO implement me
	panic("implement me")
}
func (o *OpenLibraryApiSearchProducer) ProduceAuthorsSearchResults(foundAuthors []openlibrary.Author) error {
	panic("not implemented") // TODO: Implement
}

// DoAuthorsSearch
// / First we will be executing a search that returns all the authors, but the problem the returned search results do not
// / provide full information about the authors, but what is provided from search results are the OpenLibrary author ids.
// / Using these author ids, we will be able to fetch an author one by one using goroutines and buffered channels.
// / Why am I using a buffered channel? Because we know the number of authors returned from the initial search so
// / a definite number of goroutines will be created, and it will write into the channel.
func (o *OpenLibraryApiSearchProducer) DoAuthorsSearch(searchText string) ([]openlibrary.Author, error) {
	authors := []openlibrary.Author{}

	httpClient := httpClients.Get().(*http.Client)
	defer httpClients.Put(httpClient)

	searchQuery := openLibraryAuthorsSearchAPI + strings.ReplaceAll(searchText, " ", "+")
	initialSearchRequest, err := http.NewRequest("GET", searchQuery, nil)
	if err != nil {
		log.Fatalln()
		return nil, fmt.Errorf("failed to search for authors in OpenLibrary, reason %s\n", err.Error())
	}
	res, err := httpClient.Do(initialSearchRequest)
	if err != nil {
		log.Fatalln()
		return nil, err
	}

	defer res.Body.Close()
	searchResultsInBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to parse the response body from author search results, reason :%s\n", err.Error())
	}
	authorSearchResultsJsonBlob := string(searchResultsInBytes)
	authorsDocs := gjson.Parse(authorSearchResultsJsonBlob).Get("docs").Array()

	numOfAuthors := len(authorsDocs)
	authorsChan := make(chan openlibrary.Author, numOfAuthors)
	authorFetchingCounter := sync.WaitGroup{}
	authorFetchingCounter.Add(numOfAuthors)

	for i, authorDoc := range authorsDocs {
		authorId := authorDoc.Get("key").Str
		log.Print(strconv.Itoa(i) + " " + string(authorId))
		go o.fetchAuthorWithId(authorId, &authorFetchingCounter, authorsChan)
	}

	go func() {
		authorFetchingCounter.Wait()
		close(authorsChan)

		for author := range authorsChan {
			log.Println(author)
			authors = append(authors, author)
		}
	}()

	return authors, nil
}

func (o OpenLibraryApiSearchProducer) fetchAuthorWithId(authorId string, authorFetchingCounter *sync.WaitGroup, authors chan<- openlibrary.Author) {

	// To speed things up, for each fetch, we will use a http client from the pool if it is available
	httpClient := httpClients.Get().(*http.Client)
	defer httpClients.Put(httpClient)

	// Let's build the url for fetching the author with the id
	requestUrl := openLibraryBooksAPI + authorId + ".json"
	authorFetchRequest, err := http.NewRequest("GET", requestUrl, nil)
	if err != nil {
		log.Printf("failed the create a request to fetch the author with id %s because %s", authorId, err.Error())
		return
	}
	res, err := httpClient.Do(authorFetchRequest)
	if err != nil {
		log.Printf("failed to fetch the author with id %s because %s ", authorId, err.Error())
		return
	}

	authorInBytes, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("failed to parsed the author from bytes because %s", err.Error())
		return
	}

	authorJsonBlob := string(authorInBytes)
	parsedJsonAuthor := gjson.Parse(authorJsonBlob)
	log.Printf("in the fetching go routine %s", parsedJsonAuthor)
	// Some pictures ids will be -1, they need to be filtered out because they are useless
	authorPhotosIds := []int64{}
	for _, photoId := range parsedJsonAuthor.Get("photos").Array() {
		if photoId.Int() == -1 {
			continue
		}
		authorPhotosIds = append(authorPhotosIds, photoId.Int())
	}

	// Because OpenLibrary is not uniform when it comes to author's bio, we will be checking various fields and take the string
	// that will be having length more than one
	authorBio := ""
	if parsedBio := parsedJsonAuthor.Get("bio").Str; parsedBio != "" {
		authorBio = parsedBio
	} else if parsedBio := parsedJsonAuthor.Get("bio").Get("value").Str; parsedBio != "" {
		authorBio = parsedBio
	}

	author := openlibrary.Author{
		Key:           parsedJsonAuthor.Get("key").Str,
		Name:          parsedJsonAuthor.Get("name").Str,
		PersonalName:  parsedJsonAuthor.Get("personal_name").Str,
		BirthDate:     parsedJsonAuthor.Get("birth_date").Str,
		DeathDate:     parsedJsonAuthor.Get("death_date").Str,
		Bio:           authorBio,
		Photos:        authorPhotosIds,
		WebsiteLink:   parsedJsonAuthor.Get("website").Str,
		WikipediaLink: parsedJsonAuthor.Get("wikipedia").Str,
	}

	authors <- author

	authorFetchingCounter.Done()
}
