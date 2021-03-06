package feeder

type Channel struct {
	Title          string
	Links          []Link
	Description    string
	Language       string
	Copyright      string
	ManagingEditor string
	WebMaster      string
	PubDate        string
	LastBuildDate  string
	Docs           string
	Categories     []*Category
	Generator      Generator
	TTL            int
	Rating         string
	SkipHours      []int
	SkipDays       []int
	Image          Image
	Items          []*Item
	Cloud          Cloud
	TextInput      Input

	// Atom fields
	Id       string
	Rights   string
	Author   Author
	SubTitle SubTitle
}
