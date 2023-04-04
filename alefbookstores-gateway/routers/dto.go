package routers

type requestError struct {
	message string `json:"message"`
}

type book struct {
	BookId           string   `json:"bookId"`
	Etag             string   `json:"etag"`
	SelfLink         string   `json:"selfLink"`
	Title            string   `json:"title"`
	Subtitle         string   `json:"subtitle"`
	Authors          []author `json:"authors"`
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

type author struct {
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

type quote struct {
	QuoteId string `json:"quoteId"`
	Text    string `json:"text"`
	Book    book   `json:"book"`
	Author  author `json:"author"`
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
