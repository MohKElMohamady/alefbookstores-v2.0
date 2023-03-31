package google

import (
	"log"
	"github.com/tidwall/gjson"
)

// https://developers.google.com/books/docs/v1/reference/volumes#resource-volumes
type GoogleBook struct {
	Kind          string
	Id            string
	Etag          string
	SelfLink      string
	Title         string
	Subtitle      string
	Authors       []string
	Publisher     string
	PublishedDate string
	Description   string
	Isbn10        string
	Isbn13        string
	PageCount     uint64
	Price         float64
	/* Print type is not hardcover or softcover but rather if it is a book or a magazine */
	PrintType string
	/* MainCategory is optional as we have categories, just comment it out for future use */
	// MainCategory   string
	Categories     []string
	AverageRating  float64
	RatingsCount   int64
	ContentVersion string
	ImageLinks     ImageLink
	Language       string
	PreviewLink    string
	InfoLink       string
	IsEBook        bool
}

//TODO Format the book to printed in a nice
func (gb GoogleBook) String() string {
	return ""
}

type IndustryIdentifier struct {
	T           string
	Indentifier string
}

type ImageLink struct {
	SmallThumbnail string
	Thumbnail      string
	Small          string
	Medium         string
	Large          string
	ExtraLarge     string
}

func ParseGoogleBookSearchResult(searchResultJsonBlob string) []GoogleBook {
	booksInJson := gjson.Parse(searchResultJsonBlob).Get("items").Array()

	var books []GoogleBook
	for _, book := range booksInJson {
		volumeInfo := book.Get("volumeInfo")

		/* Parsing all of the authors and adding them to an array */
		authors := []string{}
		for _, author := range volumeInfo.Get("authors").Array() {
			//log.Println(author)
			authors = append(authors, author.Str)
		}

		/* Parsing all the categories/genres and adding them to an array */
		categories := []string{}
		for _, category := range volumeInfo.Get("categories").Array() {
			categories = append(categories, category.Str)
		}

		/* Industry identifiers are ISBN10 and ISBN13 */
		industryIdentifiersMap := make(map[string]string)
		for _, industryIdentifier := range volumeInfo.Get("industryIdentifiers").Array() {

			industryIdentifiersMap[industryIdentifier.Get("type").Str] = industryIdentifier.Get("identifier").Str
		}

		/* Sometimes the price is available and specified by salesInfo, it might not be present and if it is not, generate a random price */
		var designatedPriceForBook float64 = 0
		if googleBookPrice := book.Get("saleInfo").Get("listPrice").Get("amount").Num; googleBookPrice != 0 {
			designatedPriceForBook = googleBookPrice
		} else {
			//TODO replace this fixed price with a random generated price between 10 and 60
			designatedPriceForBook = 15
		}
		searchedBook := GoogleBook{
			Kind:          book.Get("kind").Str,
			Id:            book.Get("id").Str,
			Etag:          book.Get("etag").Str,
			SelfLink:      book.Get("selfLink").Str,
			Title:         volumeInfo.Get("title").Str,
			Subtitle:      volumeInfo.Get("subtitle").Str,
			Authors:       authors,
			Publisher:     volumeInfo.Get("publisher").Str,
			PublishedDate: volumeInfo.Get("publishedDate").Str,
			Isbn10:        industryIdentifiersMap["ISBN_10"],
			Isbn13:        industryIdentifiersMap["ISBN_13"],
			Price:         designatedPriceForBook,
			Description:   volumeInfo.Get("description").Str,
			Categories:    categories,
			PageCount:     uint64(volumeInfo.Get("pageCount").Int()),
			Language:      volumeInfo.Get("language").Str,
			PreviewLink:   volumeInfo.Get("previewLink").Str,
			AverageRating: volumeInfo.Get("averageRating").Float(),
			RatingsCount:  volumeInfo.Get("ratingsCount").Int(),
			InfoLink:      volumeInfo.Get("infoLink").Str,
			IsEBook:       volumeInfo.Get("isEbook").Bool(),
		}
		log.Println(searchedBook.Title)
		books = append(books, searchedBook)
	}
	return books
}
