package openlibrary

type Author struct {
	Key           string
	Name          string
	PersonalName  string
	BirthDate     string
	DeathDate     string
	Bio           string
	Photos        []int64
	WebsiteLink   string
	WikipediaLink string
}
