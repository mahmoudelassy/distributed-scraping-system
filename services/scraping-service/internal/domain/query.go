package domain

type Query struct {
	Selector string `json:"selector"`
	Label    string `json:"label"`
	All      bool   `json:"all"`
}

func (q *Query) Select(document Document) *Result {
	result := &Result{
		Query:  q,
		Status: QuerySuccess,
	}

	if q.All {
		elements, err := document.GetAll(q.Selector)
		if err != nil {
			result.Status = QueryFailed
			return result
		}

		result.Elements = elements
		return result
	}

	el, err := document.Get(q.Selector)
	if err != nil {
		result.Status = QueryFailed
		return result
	}

	result.Elements = []HTMLElement{el}
	return result
}
