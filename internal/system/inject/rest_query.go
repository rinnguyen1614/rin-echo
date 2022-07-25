package inject

import (
	rquery "github.com/rinnguyen1614/rin-echo/internal/core/echo/models/query/rest_query"
	query "github.com/rinnguyen1614/rin-echo/internal/core/query"
)

func GetRestQuery() rquery.RestQuery {
	if service.query == nil {
		cfg := GetConfig()
		service.query = rquery.NewRestQuery(rquery.RestConfig{
			Config: query.Config{
				MaxPageSize: cfg.App.MaxPageSize,
				MinPageSize: cfg.App.MinPageSize,
			},
		})
	}
	return service.query
}
