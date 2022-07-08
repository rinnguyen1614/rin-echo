package rest_query

import (
	"net/http"
	query "rin-echo/common/query"
	"strconv"
)

type (
	RestConfig struct {
		PagingRequestHandle    PagingRequestHandle
		SortingRequestHandle   SortingRequestHandle
		SelectingRequestHandle SelectingRequestHandle
		FilteringRequestHandle FilteringRequestHandle
		query.Config
	}
	PagingRequestHandle    func(req *http.Request) (page int, pageSize int)
	SortingRequestHandle   func(req *http.Request) string
	SelectingRequestHandle func(req *http.Request) string
	FilteringRequestHandle func(req *http.Request) string
)

var (
	DefaultConfig = RestConfig{
		PagingRequestHandle: func(req *http.Request) (int, int) {
			page, _ := strconv.Atoi(req.URL.Query().Get("page"))
			if page == 0 {
				page = 1
			}

			pageSize, _ := strconv.Atoi(req.URL.Query().Get("page_size"))
			return page, pageSize
		},
		SortingRequestHandle: func(req *http.Request) string {
			return req.URL.Query().Get("sort")
		},
		SelectingRequestHandle: func(req *http.Request) string {
			return req.URL.Query().Get("select")
		},
		FilteringRequestHandle: func(req *http.Request) string {
			return req.URL.Query().Get("filter")
		},
		Config: query.Config{
			MaxPageSize:       100,
			MinPageSize:       10,
			SeparateFields:    ",",
			SeparateSortField: ":",
		},
	}
)

type RestQuery interface {
	Query(req *http.Request) (*query.Query, error)
}

func NewRestQuery(config RestConfig) RestQuery {
	if config.MaxPageSize <= 0 {
		config.MaxPageSize = DefaultConfig.MaxPageSize
	}
	if config.MinPageSize <= 0 {
		config.MaxPageSize = DefaultConfig.MinPageSize
	}
	if config.SeparateFields == "" {
		config.SeparateFields = DefaultConfig.SeparateFields
	}
	if config.SeparateSortField == "" {
		config.SeparateSortField = DefaultConfig.SeparateSortField
	}
	if config.PagingRequestHandle == nil {
		config.PagingRequestHandle = DefaultConfig.PagingRequestHandle
	}
	if config.SortingRequestHandle == nil {
		config.SortingRequestHandle = DefaultConfig.SortingRequestHandle
	}
	if config.SelectingRequestHandle == nil {
		config.SelectingRequestHandle = DefaultConfig.SelectingRequestHandle
	}
	if config.FilteringRequestHandle == nil {
		config.FilteringRequestHandle = DefaultConfig.FilteringRequestHandle
	}
	return &restQuery{
		config,
	}
}

type restQuery struct {
	RestConfig
}

func (r *restQuery) Query(req *http.Request) (*query.Query, error) {
	page, pageSize := r.PagingRequestHandle(req)
	sorts := r.SortingRequestHandle(req)
	selects := r.SelectingRequestHandle(req)
	filter := r.FilteringRequestHandle(req)

	return query.Parse(sorts, selects, filter, page, pageSize, r.Config)
}
