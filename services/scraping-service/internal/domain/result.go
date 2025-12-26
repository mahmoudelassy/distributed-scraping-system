package domain

type QueryStatus string

const (
	QuerySuccess QueryStatus = "SUCCESS"
	QueryFailed  QueryStatus = "FAILED"
)

type Result struct {
	Query    *Query
	Elements []HTMLElement
<<<<<<< Updated upstream
=======
	Status   QueryStatus
>>>>>>> Stashed changes
}
