package clients

import (
	"alefbookstores-bibliotheca/pb"
	"alefbookstores-gateway/routers"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"io"
	"log"
	"sync"
)

const (
	searchServiceAddress = "localhost:50051"
)

var searchServiceConnectionPool sync.Pool

func init() {

	searchServiceConnectionPool = sync.Pool{New: func() any {

		conn, err := grpc.Dial(searchServiceAddress, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("failed to connect to remote search service host, %s", err)
		}

		return conn
	}}

}

type SearchClient struct {
}

func (c SearchClient) SearchForBookByTitle(ctx context.Context, title string) ([]routers.Book, error) {
	books := []routers.Book{}

	conn := searchServiceConnectionPool.Get().(*grpc.ClientConn)

	client := pb.NewAlefBookstoresSearchServiceClient(conn)

	foundBooks, err := client.SearchForBookByTitle(ctx, &wrapperspb.StringValue{Value: title})
	if err != nil {
		return nil, fmt.Errorf("failed to invoke SearchForBookByTitle, reason %s", err)
	}

	for {
		book, err := foundBooks.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("failed to receive one book, %s", err)
		}
		books = append(books, routers.BookFromGrpc(book))
	}

	return books, nil
}

func (c SearchClient) SearchForBookByAuthor(ctx context.Context, title string) ([]routers.Book, error) {
	books := []routers.Book{}

	conn := searchServiceConnectionPool.Get().(*grpc.ClientConn)

	client := pb.NewAlefBookstoresSearchServiceClient(conn)

	foundBooks, err := client.SearchForBookByAuthorName(ctx, &wrapperspb.StringValue{Value: title})
	if err != nil {
		return nil, fmt.Errorf("failed to invoke SearchForBookByTitle, reason %s", err)
	}

	for {
		book, err := foundBooks.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("failed to receive one book, %s", err)
		}
		books = append(books, routers.BookFromGrpc(book))
	}

	return books, nil
}

func (c *SearchClient) SearchForAuthor(context context.Context, name string) ([]routers.Author, error) {
	authors := []routers.Author{}

	conn := searchServiceConnectionPool.Get().(*grpc.ClientConn)

	client := pb.NewAlefBookstoresSearchServiceClient(conn)

	foundAuthors, err := client.SearchForAuthorByName(context, &wrapperspb.StringValue{Value: name})
	if err != nil {
		return nil, fmt.Errorf("failed to invoke SearchForAuthorByName, reason %s", err)
	}

	for {
		author, err := foundAuthors.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("failed to receive one author %s \n", err)
		}

		authors = append(authors, routers.AuthorFromGrpc(author))
	}

	return authors, nil
}
