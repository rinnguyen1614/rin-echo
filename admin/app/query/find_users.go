package query

import (
	"rin-echo/admin/domain"
	querybuilder "rin-echo/admin/domain/query_builder"
	echox "rin-echo/common/echo"
	"rin-echo/common/echo/models"
	query "rin-echo/common/query"
	iuow "rin-echo/common/uow/interfaces"
)

type FindUsersHandler struct {
	uow iuow.UnitOfWork
}

func NewFindUsersHandler(uow iuow.UnitOfWork) FindUsersHandler {
	if uow == nil {
		panic("NewTokenUserHandler requires uow")
	}

	return FindUsersHandler{uow}
}

func (h FindUsersHandler) Handle(c echox.Context, q *query.Query) (*models.QueryResult, error) {
	var (
		users []*domain.User
	)

	var (
		queryBuilder    = querybuilder.NewUserQueryBuilder(h.uow.DB())
		preloadBuilders = map[string]query.QueryBuilder{
			"UserRoles": querybuilder.NewUserRoleQueryBuilder(h.uow.DB()),
			"Role":      querybuilder.NewRoleQueryBuilder(h.uow.DB()),
		}
	)

	err := q.Bind(queryBuilder, preloadBuilders, &User{})

	if err != nil {
		return nil, err
	}

	err = queryBuilder.WithContext(c.RequestContext()).Find(&users)
	if err != nil {
		return nil, err
	}

	return models.NewQueryResult(newUsers(users), 0, q.Paging().Limit, q.Paging().Offset), nil
}
