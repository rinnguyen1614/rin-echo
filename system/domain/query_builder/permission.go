package query_builder

import (
	"rin-echo/system/domain"

	uow "rin-echo/common/uow"

	iuow "rin-echo/common/uow/interfaces"
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
