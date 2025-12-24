package domain

type Attributes = map[string]string

type HTMLElement interface {
	GetText() string
	GetAttributes() Attributes
}
