package serp

type Spider interface {
	Search(urls ...string) (resp []SpiderResult)
}

type SpiderResult struct {
	Title       string
	Url         string
	Description string
	Content     string
	Error       error
}
