fmt:
	go fmt ./..
proto common:
	protoc -I=./alefbookstores-bibliotheca/proto --go_out=./alefbookstores-bibliotheca/pb --go-grpc_out=./alefbookstores-bibliotheca/pb ./alefbookstores-bibliotheca/proto/alefbookstores_common.proto

proto search service:
	protoc -I=./alefbookstores-bibliotheca/proto --go_out=./alefbookstores-bibliotheca/pb --go-grpc_out=./alefbookstores-bibliotheca/pb ./alefbookstores-bibliotheca/proto/alefbookstores_search_service.proto
proto quotes service:
	protoc -I=./alefbookstores-bibliotheca/proto --go_out=./alefbookstores-bibliotheca/pb --go-grpc_out=./alefbookstores-bibliotheca/pb ./alefbookstores-bibliotheca/proto/alefbookstores_quotes_service.proto
