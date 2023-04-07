package builders

import (
	"alefbookstores-bibliotheca/pb"
	"alefbookstores-search-service/google"
)

func NewGrpcAlefBookstores(b google.GoogleBook) *pb.AlefBookStoresBook {
	return &pb.AlefBookStoresBook{
		Id:               b.Id,
		Etag:             b.Etag,
		SelfLink:         b.SelfLink,
		Title:            b.Title,
		Subtitle:         b.Subtitle,
		Authors:          b.Authors,
		Publisher:        b.Publisher,
		PublishedDate:    b.PublishedDate,
		Isbn10:           b.Isbn10,
		Isbn13:           b.Isbn13,
		PageCount:        int64(b.PageCount),
		Price:            b.Price,
		Categories:       b.Categories,
		AverageRating:    b.AverageRating,
		RatingsCount:     b.RatingsCount,
		Language:         b.Language,
		PreviewLink:      b.PreviewLink,
		InfoLink:         b.InfoLink,
		IsEbookAvailable: b.IsEBook,
	}
}
