package models

import (
	"alefbookstores-bibliotheca/pb"
	"encoding/json"
	"fmt"
	"log"
)

type requestError struct {
	message string `json:"message"`
}

type Book struct {
	BookId           string   `json:"bookId"`
	Etag             string   `json:"etag"`
	SelfLink         string   `json:"selfLink"`
	Title            string   `json:"title"`
	Subtitle         string   `json:"subtitle"`
	Authors          []Author `json:"authors"`
	Publisher        string   `json:"publisher"`
	PublisherDate    string   `json:"publisherDate"`
	Isbn10           string   `json:"isbn10"`
	Isbn13           string   `json:"isbn13"`
	PageCount        int64    `json:"pageCount"`
	Price            float64  `json:"price"`
	Categories       []string `json:"categories"`
	AverageRating    float64  `json:"averageRating"`
	RatingCounts     int64    `json:"ratingCounts"`
	Language         string   `json:"language"`
	PreviewLink      string   `json:"previewLink"`
	InfoLink         string   `json:"infoLink"`
	IsEbookAvailable bool     `json:"isEbookAvailable"`
}

func BookFromGrpc(b *pb.AlefBookStoresBook) Book {
	authors := []Author{}
	for _, author := range b.Authors {
		authors = append(authors, Author{Name: author})
	}
	return Book{
		BookId:           b.Id,
		Etag:             b.Etag,
		SelfLink:         b.SelfLink,
		Title:            b.Title,
		Subtitle:         b.Subtitle,
		Authors:          authors,
		Publisher:        b.Publisher,
		PublisherDate:    b.PublishedDate,
		Isbn10:           b.Isbn10,
		Isbn13:           b.Isbn13,
		PageCount:        b.PageCount,
		Price:            b.Price,
		Categories:       b.Categories,
		AverageRating:    b.AverageRating,
		RatingCounts:     b.RatingsCount,
		Language:         b.Language,
		PreviewLink:      b.PreviewLink,
		InfoLink:         b.InfoLink,
		IsEbookAvailable: b.IsEbookAvailable,
	}
}

type Author struct {
	AuthorId      string  `json:"authorId"`
	Name          string  `json:"name"`
	PersonalName  string  `json:"personalName"`
	BirthDate     string  `json:"birthDate"`
	DeathDate     string  `json:"deathDate"`
	Bio           string  `json:"bio"`
	Photos        []int64 `json:"photos"`
	WebsiteLink   string  `json:"websiteLink"`
	WikipediaLink string  `json:"wikipediaLink"`
}

func AuthorFromGrpc(a *pb.Author) Author {
	return Author{
		AuthorId:      a.Key,
		Name:          a.Name,
		PersonalName:  a.PersonalName,
		BirthDate:     a.BirthDate,
		DeathDate:     a.DeathDate,
		Bio:           a.Bio,
		Photos:        a.Photos,
		WebsiteLink:   a.WebsiteLink,
		WikipediaLink: a.WikipediaLink,
	}
}

type Quote struct {
	QuoteId string `json:"quoteId"`
	Text    string `json:"text"`
	Book    Book   `json:"Book"`
	Author  Author `json:"Author"`
}

type QuoteCreationDto struct {
	Text     string `json:"text"`
	AuthorId string `json:"authorId"`
	BookId   string `json:"bookId"`
}

type QuoteUpdateDto struct {
	QuoteId     string `json:"quoteId"`
	UpdatedText string `json:"text"`
	AuthorId    string `json:"authorId"`
	BookId      string `json:"bookId"`
}

type QuoteDeleteDto struct {
	QuoteId  string `json:"quoteId"`
	AuthorId string `json:"authorId"`
	BookId   string `json:"bookId"`
}

func RequestError(reason string) requestError {
	return requestError{
		message: reason,
	}
}

func JsonParsingError(failedToParsedObj interface{}, reason error) []byte {
	errMsg := requestError{
		message: fmt.Sprintf("failed to encode the object from json to bytes %s because %s\n", failedToParsedObj, reason),
	}

	errorMsgInBytes, err := json.Marshal(errMsg)
	if err != nil {
		log.Fatalln("even parsing this didn't succeed, terminating...")
	}

	return errorMsgInBytes
}
