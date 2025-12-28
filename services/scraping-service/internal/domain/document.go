package domain

type Document interface {
	Get(selector string) (HTMLElement, error)
	GetAll(selector string) ([]HTMLElement, error)
}
