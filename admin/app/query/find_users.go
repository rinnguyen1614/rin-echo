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
	uow          iuow.UnitOfWork
	queryBuilder querybuilder.UserQueryBuilder
}

func NewFindUsersHandler(uow iuow.UnitOfWork) FindUsersHandler {
	if uow == nil {
		panic("NewTokenUserHandler requires uow")
	}

	return FindUsersHandler{uow, querybuilder.NewUserQueryBuilder()}
}

func (h FindUsersHandler) Handle(c echox.Context, q *query.Query) (*models.QueryResult, error) {
	var (
		preloadBuilders = map[string]query.QueryBuilder{
			"UserRoles": querybuilder.NewUserRoleQueryBuilder(),
			"Role":      querybuilder.NewRoleQueryBuilder(),
		}
		users []domain.User
	)

	err := q.Bind(h.uow.DB(), &h.queryBuilder, preloadBuilders, &User{})

	if err != nil {
		return nil, err
	}

	err = h.queryBuilder.Find(h.uow.WithContext(c.RequestContext()), &users)
	if err != nil {
		return nil, err
	}

	return models.NewQueryResult(users, 0, q.Paging().Limit, q.Paging().Offset), nil
}
