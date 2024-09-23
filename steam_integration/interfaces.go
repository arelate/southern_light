package steam_integration

type DisplayTypesGetter interface {
	GetDisplayTypes() []string
}

type ResultsGetter interface {
	GetResults() []string
}

type BlogUrlGetter interface {
	GetBlogUrl() string
}
