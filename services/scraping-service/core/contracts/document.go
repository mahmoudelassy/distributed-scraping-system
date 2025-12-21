package contracts

type Document interface {
	Get(selector string) (HTMLElement, error)
	GetAll(selector string) []HTMLElement
}
