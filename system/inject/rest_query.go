package inject

import (
	rquery "rin-echo/common/echo/models/query/rest_query"
	query "rin-echo/common/query"
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
