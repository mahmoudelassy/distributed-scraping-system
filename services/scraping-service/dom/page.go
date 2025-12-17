package dom

type DOMPage struct {
	URL    string
	Source *string
}

func (p *DOMPage) GetURL() string {
	return p.URL
}
func (p *DOMPage) GetSource() *string {
	return p.Source
}
