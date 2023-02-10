package paging

type Pagination struct {
}

type PageRequest struct {
	Page  int `query:"page"`
	Limit int `query:"limit"`
}

func ConvertToOffsetLimit(pageRequest *PageRequest) (int, int) {
	offset := pageRequest.Page*pageRequest.Limit + 1
	limit := 10
	if pageRequest.Limit != 0 {
		limit = pageRequest.Limit
	}
	return offset, limit
}
