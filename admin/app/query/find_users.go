package query

import (
	"rin-echo/admin/domain"
	querybuilder "rin-echo/admin/domain/query_builder"
	echox "rin-echo/common/echo"
	"rin-echo/common/echo/models"
	query "rin-echo/common/query"
	iuow "rin-echo/common/uow/interfaces"
	"rin-echo/common/utils"

	"github.com/jinzhu/copier"
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
		dbContext       = h.uow.DB().WithContext(c.RequestContext())
		queryBuilder    = querybuilder.NewUserQueryBuilder(dbContext)
		preloadBuilders = map[string]query.QueryBuilder{
			"UserRoles": querybuilder.NewUserRoleQueryBuilder(dbContext),
			"Role":      querybuilder.NewRoleQueryBuilder(dbContext),
		}
	)

	err := q.Bind(queryBuilder, preloadBuilders, &User{})

	if err != nil {
		return nil, err
	}

	err = queryBuilder.Find(&users)
	if err != nil {
		return nil, err
	}

	prune, err := utils.NewSliceOfStructsByTag(User{}, q.FlatSelect(), "json")
	if err != nil {
		return nil, err
	}

	err = copier.CopyWithOption(prune, users, copier.Option{IgnoreEmpty: true, DeepCopy: true})
	if err != nil {
		return nil, err
	}

	return models.NewQueryResult(prune, queryBuilder.Count(), q.Paging().Limit, q.Paging().Offset), nil
}
