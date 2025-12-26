package domain

type Query struct {
	Selector string
	Label    string
	All      bool
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
<<<<<<< Updated upstream
		return &result
=======
		result.Status = QueryFailed
		return result
>>>>>>> Stashed changes
	}

	result.Elements = []HTMLElement{el}
	return result
}
