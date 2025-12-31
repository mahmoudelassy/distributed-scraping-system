package domain

type QueryStatus string

const (
	QueryMatched QueryStatus = "MATCHED"
	QueryEmpty   QueryStatus = "EMPTY"
)

type Result struct {
	Query    *Query
	Elements []HTMLElement
	Status   QueryStatus
}
