package contracts

type HTMLFetcher interface {
	Fetch(url string, args ...any) (HTMLPage, error)
}
