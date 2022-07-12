package query_builder

import (
	"rin-echo/system/domain"

	uow "github.com/rinnguyen1614/rin-echo-core/uow"

	iuow "github.com/rinnguyen1614/rin-echo-core/uow/interfaces"
)

type PermissionQueryBuilder struct {
	iuow.QueryBuilder
}

func NewPermissionQueryBuilder() PermissionQueryBuilder {
	qbuilder, err := uow.NewQueryBuilder(&domain.Permission{})
	if err != nil {
		panic(err)
	}

	return PermissionQueryBuilder{
		QueryBuilder: qbuilder,
	}
}
