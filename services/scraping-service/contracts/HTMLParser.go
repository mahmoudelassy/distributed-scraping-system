package contracts

type HTMLParser interface {
	Parse(page HTMLPage)
	GetElement() HTMLElement
	GetAllElements() []HTMLElement
}
