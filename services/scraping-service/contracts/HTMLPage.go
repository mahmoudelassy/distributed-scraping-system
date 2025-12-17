package contracts

type HTMLPage interface {
	GetURL() string
	GetSource() *string
}
