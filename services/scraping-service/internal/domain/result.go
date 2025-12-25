package domain

type Result struct {
	Query    *Query
	Elements []HTMLElement
	status   string
}
