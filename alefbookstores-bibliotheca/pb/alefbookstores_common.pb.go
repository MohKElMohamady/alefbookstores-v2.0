// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: alefbookstores_common.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AlefBookStoresBook struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Etag             string   `protobuf:"bytes,2,opt,name=etag,proto3" json:"etag,omitempty"`
	SelfLink         string   `protobuf:"bytes,3,opt,name=self_link,json=selfLink,proto3" json:"self_link,omitempty"`
	Title            string   `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`
	Subtitle         string   `protobuf:"bytes,5,opt,name=subtitle,proto3" json:"subtitle,omitempty"`
	Authors          []string `protobuf:"bytes,6,rep,name=authors,proto3" json:"authors,omitempty"`
	Publisher        string   `protobuf:"bytes,7,opt,name=publisher,proto3" json:"publisher,omitempty"`
	PublishedDate    string   `protobuf:"bytes,8,opt,name=published_date,json=publishedDate,proto3" json:"published_date,omitempty"`
	Isbn10           string   `protobuf:"bytes,9,opt,name=isbn10,proto3" json:"isbn10,omitempty"`
	Isbn13           string   `protobuf:"bytes,10,opt,name=isbn13,proto3" json:"isbn13,omitempty"`
	PageCount        int64    `protobuf:"varint,11,opt,name=page_count,json=pageCount,proto3" json:"page_count,omitempty"`
	Price            float64  `protobuf:"fixed64,12,opt,name=price,proto3" json:"price,omitempty"`
	Categories       []string `protobuf:"bytes,13,rep,name=categories,proto3" json:"categories,omitempty"`
	AverageRating    float64  `protobuf:"fixed64,15,opt,name=average_rating,json=averageRating,proto3" json:"average_rating,omitempty"`
	RatingsCount     int64    `protobuf:"varint,16,opt,name=ratings_count,json=ratingsCount,proto3" json:"ratings_count,omitempty"`
	Language         string   `protobuf:"bytes,17,opt,name=language,proto3" json:"language,omitempty"`
	PreviewLink      string   `protobuf:"bytes,18,opt,name=preview_link,json=previewLink,proto3" json:"preview_link,omitempty"`
	InfoLink         string   `protobuf:"bytes,19,opt,name=info_link,json=infoLink,proto3" json:"info_link,omitempty"`
	IsEbookAvailable bool     `protobuf:"varint,20,opt,name=is_ebook_available,json=isEbookAvailable,proto3" json:"is_ebook_available,omitempty"`
}

func (x *AlefBookStoresBook) Reset() {
	*x = AlefBookStoresBook{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alefbookstores_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlefBookStoresBook) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlefBookStoresBook) ProtoMessage() {}

func (x *AlefBookStoresBook) ProtoReflect() protoreflect.Message {
	mi := &file_alefbookstores_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlefBookStoresBook.ProtoReflect.Descriptor instead.
func (*AlefBookStoresBook) Descriptor() ([]byte, []int) {
	return file_alefbookstores_common_proto_rawDescGZIP(), []int{0}
}

func (x *AlefBookStoresBook) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AlefBookStoresBook) GetEtag() string {
	if x != nil {
		return x.Etag
	}
	return ""
}

func (x *AlefBookStoresBook) GetSelfLink() string {
	if x != nil {
		return x.SelfLink
	}
	return ""
}

func (x *AlefBookStoresBook) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *AlefBookStoresBook) GetSubtitle() string {
	if x != nil {
		return x.Subtitle
	}
	return ""
}

func (x *AlefBookStoresBook) GetAuthors() []string {
	if x != nil {
		return x.Authors
	}
	return nil
}

func (x *AlefBookStoresBook) GetPublisher() string {
	if x != nil {
		return x.Publisher
	}
	return ""
}

func (x *AlefBookStoresBook) GetPublishedDate() string {
	if x != nil {
		return x.PublishedDate
	}
	return ""
}

func (x *AlefBookStoresBook) GetIsbn10() string {
	if x != nil {
		return x.Isbn10
	}
	return ""
}

func (x *AlefBookStoresBook) GetIsbn13() string {
	if x != nil {
		return x.Isbn13
	}
	return ""
}

func (x *AlefBookStoresBook) GetPageCount() int64 {
	if x != nil {
		return x.PageCount
	}
	return 0
}

func (x *AlefBookStoresBook) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *AlefBookStoresBook) GetCategories() []string {
	if x != nil {
		return x.Categories
	}
	return nil
}

func (x *AlefBookStoresBook) GetAverageRating() float64 {
	if x != nil {
		return x.AverageRating
	}
	return 0
}

func (x *AlefBookStoresBook) GetRatingsCount() int64 {
	if x != nil {
		return x.RatingsCount
	}
	return 0
}

func (x *AlefBookStoresBook) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

func (x *AlefBookStoresBook) GetPreviewLink() string {
	if x != nil {
		return x.PreviewLink
	}
	return ""
}

func (x *AlefBookStoresBook) GetInfoLink() string {
	if x != nil {
		return x.InfoLink
	}
	return ""
}

func (x *AlefBookStoresBook) GetIsEbookAvailable() bool {
	if x != nil {
		return x.IsEbookAvailable
	}
	return false
}

type Author struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key           string  `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Name          string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	PersonalName  string  `protobuf:"bytes,3,opt,name=personal_name,json=personalName,proto3" json:"personal_name,omitempty"`
	BirthDate     string  `protobuf:"bytes,4,opt,name=birth_date,json=birthDate,proto3" json:"birth_date,omitempty"`
	DeathDate     string  `protobuf:"bytes,5,opt,name=death_date,json=deathDate,proto3" json:"death_date,omitempty"`
	Bio           string  `protobuf:"bytes,6,opt,name=bio,proto3" json:"bio,omitempty"`
	Photos        []int64 `protobuf:"varint,7,rep,packed,name=photos,proto3" json:"photos,omitempty"`
	WebsiteLink   string  `protobuf:"bytes,8,opt,name=website_link,json=websiteLink,proto3" json:"website_link,omitempty"`
	WikipediaLink string  `protobuf:"bytes,9,opt,name=wikipedia_link,json=wikipediaLink,proto3" json:"wikipedia_link,omitempty"`
}

func (x *Author) Reset() {
	*x = Author{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alefbookstores_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Author) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Author) ProtoMessage() {}

func (x *Author) ProtoReflect() protoreflect.Message {
	mi := &file_alefbookstores_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Author.ProtoReflect.Descriptor instead.
func (*Author) Descriptor() ([]byte, []int) {
	return file_alefbookstores_common_proto_rawDescGZIP(), []int{1}
}

func (x *Author) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Author) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Author) GetPersonalName() string {
	if x != nil {
		return x.PersonalName
	}
	return ""
}

func (x *Author) GetBirthDate() string {
	if x != nil {
		return x.BirthDate
	}
	return ""
}

func (x *Author) GetDeathDate() string {
	if x != nil {
		return x.DeathDate
	}
	return ""
}

func (x *Author) GetBio() string {
	if x != nil {
		return x.Bio
	}
	return ""
}

func (x *Author) GetPhotos() []int64 {
	if x != nil {
		return x.Photos
	}
	return nil
}

func (x *Author) GetWebsiteLink() string {
	if x != nil {
		return x.WebsiteLink
	}
	return ""
}

func (x *Author) GetWikipediaLink() string {
	if x != nil {
		return x.WikipediaLink
	}
	return ""
}

var File_alefbookstores_common_proto protoreflect.FileDescriptor

var file_alefbookstores_common_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x61, 0x6c, 0x65, 0x66, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x73,
	0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x61,
	0x6c, 0x65, 0x66, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x73, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x22, 0xc1, 0x04, 0x0a, 0x12, 0x41, 0x6c, 0x65, 0x66, 0x42, 0x6f, 0x6f,
	0x6b, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x73, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x65,
	0x74, 0x61, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x65, 0x74, 0x61, 0x67, 0x12,
	0x1b, 0x0a, 0x09, 0x73, 0x65, 0x6c, 0x66, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x6c, 0x66, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x75, 0x62, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x75, 0x62, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x07, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x73, 0x68, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73,
	0x68, 0x65, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x69, 0x73, 0x62, 0x6e, 0x31, 0x30, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69,
	0x73, 0x62, 0x6e, 0x31, 0x30, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x73, 0x62, 0x6e, 0x31, 0x33, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x73, 0x62, 0x6e, 0x31, 0x33, 0x12, 0x1d, 0x0a,
	0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73,
	0x18, 0x0d, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69,
	0x65, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x61, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x72, 0x61,
	0x74, 0x69, 0x6e, 0x67, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0d, 0x61, 0x76, 0x65, 0x72,
	0x61, 0x67, 0x65, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x61, 0x74,
	0x69, 0x6e, 0x67, 0x73, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x10, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0c, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72,
	0x65, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x70, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x1b, 0x0a,
	0x09, 0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x69, 0x6e, 0x66, 0x6f, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x2c, 0x0a, 0x12, 0x69, 0x73,
	0x5f, 0x65, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65,
	0x18, 0x14, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x69, 0x73, 0x45, 0x62, 0x6f, 0x6f, 0x6b, 0x41,
	0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x22, 0x85, 0x02, 0x0a, 0x06, 0x41, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x62, 0x69, 0x72, 0x74, 0x68, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x62, 0x69, 0x72, 0x74, 0x68, 0x44, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x64, 0x65, 0x61, 0x74, 0x68, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x64, 0x65, 0x61, 0x74, 0x68, 0x44, 0x61, 0x74, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x62, 0x69, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x62, 0x69, 0x6f, 0x12, 0x16,
	0x0a, 0x06, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x03, 0x52, 0x06,
	0x70, 0x68, 0x6f, 0x74, 0x6f, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74,
	0x65, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x77, 0x65,
	0x62, 0x73, 0x69, 0x74, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x25, 0x0a, 0x0e, 0x77, 0x69, 0x6b,
	0x69, 0x70, 0x65, 0x64, 0x69, 0x61, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x77, 0x69, 0x6b, 0x69, 0x70, 0x65, 0x64, 0x69, 0x61, 0x4c, 0x69, 0x6e, 0x6b,
	0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_alefbookstores_common_proto_rawDescOnce sync.Once
	file_alefbookstores_common_proto_rawDescData = file_alefbookstores_common_proto_rawDesc
)

func file_alefbookstores_common_proto_rawDescGZIP() []byte {
	file_alefbookstores_common_proto_rawDescOnce.Do(func() {
		file_alefbookstores_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_alefbookstores_common_proto_rawDescData)
	})
	return file_alefbookstores_common_proto_rawDescData
}

var file_alefbookstores_common_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_alefbookstores_common_proto_goTypes = []interface{}{
	(*AlefBookStoresBook)(nil), // 0: alefbookstores.common.AlefBookStoresBook
	(*Author)(nil),             // 1: alefbookstores.common.Author
}
var file_alefbookstores_common_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_alefbookstores_common_proto_init() }
func file_alefbookstores_common_proto_init() {
	if File_alefbookstores_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_alefbookstores_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlefBookStoresBook); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_alefbookstores_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Author); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_alefbookstores_common_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_alefbookstores_common_proto_goTypes,
		DependencyIndexes: file_alefbookstores_common_proto_depIdxs,
		MessageInfos:      file_alefbookstores_common_proto_msgTypes,
	}.Build()
	File_alefbookstores_common_proto = out.File
	file_alefbookstores_common_proto_rawDesc = nil
	file_alefbookstores_common_proto_goTypes = nil
	file_alefbookstores_common_proto_depIdxs = nil
}
