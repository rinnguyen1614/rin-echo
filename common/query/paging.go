package query

type Paging struct {
	Limit  int
	Offset int
}

func ParsePaging(pageSize, page, maxPageSize, minPageSize int) Paging {
	switch {
	case pageSize > maxPageSize:
		pageSize = maxPageSize
	case pageSize <= 0:
		pageSize = minPageSize
	}
	return Paging{
		Limit:  pageSize,
		Offset: (page - 1) * pageSize,
	}
}
