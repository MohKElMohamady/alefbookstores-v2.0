package main

import (
	"alefbookstores-bibliotheca/builders"
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

func (a AlefBookStoresSearchService) SearchForBookByTitle(v *wrapperspb.StringValue, server bibliotheca.AlefBookstoresSearchService_SearchForBookByTitleServer) error {
	googleBooks, err := a.googleSearchProducer.DoBooksSearchByTitle(v.Value)
	if err != nil {
		return err
	}

	for _, book := range googleBooks {
		gRPCBook := builders.NewGrpcAlefBookstores(book)
		log.Printf("sending the author %s \n", gRPCBook)
		err := server.Send(gRPCBook)
		if err != nil {
			return fmt.Errorf("failed to send author %s to the client because %s", gRPCBook, err)
		}
	}
	return nil
}

func (a AlefBookStoresSearchService) SearchForBookByAuthorName(v *wrapperspb.StringValue, server bibliotheca.AlefBookstoresSearchService_SearchForBookByAuthorNameServer) error {
	googleBooks, err := a.googleSearchProducer.DoBooksSearchByAuthor(v.Value)
	if err != nil {
		return err
	}

	for _, book := range googleBooks {
		gRPCBook := builders.NewGrpcAlefBookstores(book)
		log.Printf("sending the author %s \n", gRPCBook)
		err := server.Send(gRPCBook)
		if err != nil {
			return fmt.Errorf("failed to send author %s to the client because %s", gRPCBook, err)
		}
	}
	return nil
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
