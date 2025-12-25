package domain

type Query struct {
	Selector string `json:"selector" validate:"required"`
	Label    string `json:"label" validate:"required"`
	All      bool   `json:"all" validate:"required"`
}

func (q *Query) Select(document Document) *Result {
	result := Result{
		Query:    q,
		Elements: nil,
	}

	if q.All {
		result.Elements = document.GetAll(q.Selector)
		return &result
	}

	el, err := document.Get(q.Selector)
	if err != nil {
		result.status = ""
		return &result
	}

	result.Elements = append(result.Elements, el)
	return &result
}
