package domain

type Query struct {
	Selector string
	Label    string
	All      bool
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
		return &result
	}

	result.Elements = append(result.Elements, el)
	return &result
}
