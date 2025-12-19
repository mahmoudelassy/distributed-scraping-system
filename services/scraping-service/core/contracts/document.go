package contracts

type Document interface {
	Get(selector string) HTMLElement
	GetAll(selector string) []HTMLElement
}
