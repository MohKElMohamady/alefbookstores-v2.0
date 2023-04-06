package main

import (
	bibliotheca "alefbookstores-bibliotheca/pb"
	"alefbookstores-search-service/kafka"
	"fmt"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"log"
)

type AlefBookStoresSearchService struct {
	bibliotheca.UnimplementedAlefBookstoresSearchServiceServer
	googleSearchProducer      kafka.GoogleApiSearchProducer
	openLibrarySearchProducer kafka.OpenLibraryApiSearchProducer
}

func NewAlefBookStoresSearchService() AlefBookStoresSearchService {
	return AlefBookStoresSearchService{
		googleSearchProducer:      kafka.GoogleApiSearchProducer{},
		openLibrarySearchProducer: kafka.OpenLibraryApiSearchProducer{},
	}
}

func (a AlefBookStoresSearchService) SearchForBookByTitle(value *wrapperspb.StringValue, server bibliotheca.AlefBookstoresSearchService_SearchForBookByTitleServer) error {
	//TODO implement me
	panic("implement me")
}

func (a AlefBookStoresSearchService) SearchForBookByAuthorName(value *wrapperspb.StringValue, server bibliotheca.AlefBookstoresSearchService_SearchForBookByAuthorNameServer) error {
	//TODO implement me
	panic("implement me")
}

func (a *AlefBookStoresSearchService) SearchForAuthorByName(v *wrapperspb.StringValue, server bibliotheca.AlefBookstoresSearchService_SearchForAuthorByNameServer) error {

	log.Printf("gRPC search service called by client searching for an author by name %s\n", v.Value)

	openLibraryAuthors, _ := a.openLibrarySearchProducer.DoAuthorsSearch(v.Value)

	for _, author := range openLibraryAuthors {
		grpcAuthor := &bibliotheca.Author{
			Key:           author.Key,
			Name:          author.Name,
			PersonalName:  author.PersonalName,
			BirthDate:     author.BirthDate,
			DeathDate:     author.DeathDate,
			Bio:           author.Bio,
			Photos:        author.Photos,
			WebsiteLink:   author.WebsiteLink,
			WikipediaLink: author.WikipediaLink,
		}
		log.Printf("%s", grpcAuthor)
		err := server.Send(grpcAuthor)
		if err != nil {
			return fmt.Errorf("failed to send author %s to the client because %s", grpcAuthor, err)
		}
	}

	return nil
}
