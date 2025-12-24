package domain

type ScraperError struct {
	Stage string
	URL   string
	Err   error
}

func (e *ScraperError) Error() string { return e.Stage + " error for " + e.URL + ": " + e.Err.Error() }
func (e *ScraperError) Unwrap() error { return e.Err }
