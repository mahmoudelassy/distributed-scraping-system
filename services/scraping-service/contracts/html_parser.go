package contracts

type HTMLParser interface {
	GetElement(selector string) HTMLElement
	GetAllElements(selector string) []HTMLElement
}
