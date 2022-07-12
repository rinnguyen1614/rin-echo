package query_builder

import (
	"fmt"
	"rin-echo/system/domain"

	uow "github.com/rinnguyen1614/rin-echo-core/uow"

	iuow "github.com/rinnguyen1614/rin-echo-core/uow/interfaces"
)

type ResourceQueryBuilder struct {
	iuow.QueryBuilder
}

func NewResourceQueryBuilder() ResourceQueryBuilder {
	qbuilder, err := uow.NewQueryBuilder(&domain.Resource{})
	if err != nil {
		panic(err)
	}

	qbuilder.SetOrder("parent_id", "asc NULLS FIRST")
	qbuilder.SetOrder("id", "asc")

	return ResourceQueryBuilder{
		QueryBuilder: qbuilder,
	}
}

func (q *ResourceQueryBuilder) WhereID(id uint) {
	q.SetCondition("id", id)
}

func (q *ResourceQueryBuilder) WhereIDsIn(ids []uint) {
	q.SetCondition(fmt.Sprintf("%v %v ?", "id", "IN"), ids)
}

func (q *ResourceQueryBuilder) WhereIDsNotIn(ids []uint) {
	q.SetCondition(fmt.Sprintf("%v %v ?", "id", "NOT IN"), ids)
}

func (q *ResourceQueryBuilder) WhereSlug(slug string) {
	q.SetCondition("slug", slug)
}

func (q *ResourceQueryBuilder) WhereSlugsIn(slugs []string) {
	q.SetCondition(fmt.Sprintf("%v %v ?", "slug", "IN"), slugs)
}

func (q *ResourceQueryBuilder) WhereSlugsNotIn(slugs []string) {
	q.SetCondition(fmt.Sprintf("%v %v ?", "slug", "NOT IN"), slugs)
}
