package querytypes

type Pagination struct {
	Limit  int
	Offset int
}

func (p Pagination) Unwrap() (offset, limit int) {
	offset = p.Offset
	limit = p.Limit
	return
}
