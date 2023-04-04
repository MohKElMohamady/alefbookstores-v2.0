package routers

import "fmt"

type BookBuilder struct {
	bookId           string
	etag             string
	selfLink         string
	title            string
	subtitle         string
	authors          []author
	publisher        string
	publisherDate    string
	isbn10           string
	isbn13           string
	pageCount        int64
	price            float64
	categories       []string
	averageRating    float64
	ratingCounts     int64
	language         string
	previewLink      string
	infoLink         string
	isEbookAvailable bool
}

func NewBookBuilder() *BookBuilder {
	return &BookBuilder{}
}

func (b BookBuilder) Build() (book, error) {

	if b.bookId == "" {
		return book{}, fmt.Errorf("a book must have an id")
	}

	if b.title == "" {
		return book{}, fmt.Errorf("a book must have a title")
	}

	if len(b.authors) == 0 {
		return book{}, fmt.Errorf("a book must have at least one author")
	}

	return book{
		BookId:           b.bookId,
		Etag:             b.etag,
		SelfLink:         b.selfLink,
		Title:            b.title,
		Subtitle:         b.subtitle,
		Authors:          b.authors,
		Publisher:        b.publisher,
		PublisherDate:    b.publisherDate,
		Isbn10:           b.isbn10,
		Isbn13:           b.isbn13,
		PageCount:        b.pageCount,
		Price:            b.price,
		Categories:       b.categories,
		AverageRating:    b.averageRating,
		RatingCounts:     b.ratingCounts,
		Language:         b.language,
		PreviewLink:      b.previewLink,
		InfoLink:         b.infoLink,
		IsEbookAvailable: b.isEbookAvailable,
	}, nil
}

func (b *BookBuilder) BookId(bookId string) *BookBuilder {
	b.bookId = bookId
	return b
}

func (b *BookBuilder) Etag(etag string) *BookBuilder {
	b.etag = etag
	return b
}

func (b *BookBuilder) SelfLink(selfLink string) *BookBuilder {
	b.selfLink = selfLink
	return b
}

func (b *BookBuilder) Title(title string) *BookBuilder {
	b.title = title
	return b
}

func (b *BookBuilder) SubTitle(subtitle string) *BookBuilder {
	b.subtitle = subtitle
	return b
}

func (b *BookBuilder) Authors(authors ...author) *BookBuilder {
	b.authors = authors
	return b
}

func (b *BookBuilder) Publisher(publisher string) *BookBuilder {
	b.publisher = publisher
	return b
}

func (b *BookBuilder) PublisherDate(publisherDate string) *BookBuilder {
	b.publisherDate = publisherDate
	return b
}

func (b *BookBuilder) Isbn10(isbn10 string) *BookBuilder {
	b.isbn10 = isbn10
	return b
}

func (b *BookBuilder) Isbn13(isbn13 string) *BookBuilder {
	b.isbn13 = isbn13
	return b
}

func (b *BookBuilder) PageCount(pageCount int64) *BookBuilder {
	b.pageCount = pageCount
	return b
}

func (b *BookBuilder) Price(price float64) *BookBuilder {
	b.price = price
	return b
}

func (b *BookBuilder) Categories(categories ...string) *BookBuilder {
	b.categories = categories
	return b
}

func (b *BookBuilder) AverageRating(averageRating float64) *BookBuilder {
	b.averageRating = averageRating
	return b
}

func (b *BookBuilder) RatingCounts(ratingCounts int64) *BookBuilder {
	b.ratingCounts = ratingCounts
	return b
}

func (b *BookBuilder) Language(language string) *BookBuilder {
	b.language = language
	return b
}

func (b *BookBuilder) PreviewLink(previewLink string) *BookBuilder {
	b.previewLink = previewLink
	return b
}

func (b *BookBuilder) InfoLink(infoLink string) *BookBuilder {
	b.infoLink = infoLink
	return b
}

func (b *BookBuilder) IsEbookAvailable(isEbookAvailable bool) *BookBuilder {
	b.isEbookAvailable = isEbookAvailable
	return b
}

type AuthorBuilder struct {
	authorId      string
	name          string
	personalName  string
	birthDate     string
	deathDate     string
	bio           string
	photos        []int64
	websiteLink   string
	wikipediaLink string
}

func (a AuthorBuilder) Build() (author, error) {

	if a.authorId == "" {
		return author{}, fmt.Errorf("cannot create an author without the id")
	}
	if a.name == "" {
		return author{}, fmt.Errorf("cannot create an author without his/her name")
	}

	return author{
		AuthorId:      a.authorId,
		Name:          a.name,
		PersonalName:  a.personalName,
		BirthDate:     a.birthDate,
		DeathDate:     a.deathDate,
		Bio:           a.bio,
		Photos:        a.photos,
		WebsiteLink:   a.websiteLink,
		WikipediaLink: a.wikipediaLink,
	}, nil
}

func (a *AuthorBuilder) AuthorId(authorId string) *AuthorBuilder {
	a.authorId = authorId
	return a
}

func (a *AuthorBuilder) Name(name string) *AuthorBuilder {
	a.name = name
	return a
}

func (a *AuthorBuilder) PersonalName(personalName string) *AuthorBuilder {
	a.personalName = personalName
	return a
}

func (a *AuthorBuilder) BirthDate(birthDate string) *AuthorBuilder {
	a.birthDate = birthDate
	return a
}

func (a *AuthorBuilder) DeathDate(deathDate string) *AuthorBuilder {
	a.deathDate = deathDate
	return a
}

func (a *AuthorBuilder) Bio(bio string) *AuthorBuilder {
	a.bio = bio
	return a
}

func (a *AuthorBuilder) Photos(photos ...int64) *AuthorBuilder {
	a.photos = photos
	return a
}

func (a *AuthorBuilder) WebsiteLink(websiteLink string) *AuthorBuilder {
	a.websiteLink = websiteLink
	return a
}

func (a *AuthorBuilder) WikipediaLink(wikipediaLink string) *AuthorBuilder {
	a.wikipediaLink = wikipediaLink
	return a
}
