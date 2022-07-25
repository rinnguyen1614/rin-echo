package query_builder

import (
	uow "github.com/rinnguyen1614/rin-echo/internal/core/uow"

	iuow "github.com/rinnguyen1614/rin-echo/internal/core/uow/interfaces"

	"github.com/rinnguyen1614/rin-echo/internal/system/domain"
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
