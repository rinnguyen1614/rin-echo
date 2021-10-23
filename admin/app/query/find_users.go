package query

import (
	"rin-echo/admin/adapters/repository"
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

func (h FindUsersHandler) Handle(ctx echox.Context, q *query.Query) (*models.QueryResult, error) {

	var (
		uow   = h.uow.WithContext(ctx.RequestContext())
		repo  = repository.NewUserRepository(uow.DB())
		users []*domain.User
	)
	var (
		queryBuilder    = querybuilder.NewUserQueryBuilder()
		preloadBuilders = map[string]iuow.QueryBuilder{
			"UserRoles": querybuilder.NewUserRoleQueryBuilder(),
			"Role":      querybuilder.NewRoleQueryBuilder(),
		}
	)

	err := q.Bind(uow.DB(), queryBuilder, preloadBuilders, &User{})

	if err != nil {
		return nil, err
	}

	err = repo.QueryBuilderFind(&users, queryBuilder)
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

	return models.NewQueryResult(prune, repo.QueryBuilderCount(queryBuilder), q.Paging().Limit, q.Paging().Offset), nil
}
