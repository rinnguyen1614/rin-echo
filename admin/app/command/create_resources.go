package command

import (
	"fmt"
	"rin-echo/admin/adapters/repository"
	"rin-echo/admin/domain"
	"rin-echo/admin/errors"
	"rin-echo/admin/inject"
	"rin-echo/common/cqrs"
	echox "rin-echo/common/echo"
	gormx "rin-echo/common/gorm"
	iuow "rin-echo/common/uow/interfaces"
	"rin-echo/common/utils"
	"strings"
)

type CreateResources []*CreateResource

func (cmds CreateResources) CreateRecursive(uow iuow.UnitOfWork, parent *domain.Resource) error {
	for _, cmd := range cmds {
		if err := cmd.CreateRecursive(uow, parent); err != nil {
			return err
		}
	}

	return nil
}

type CreateResourcesHandler struct {
	uow iuow.UnitOfWork
}

func NewCreateResourcesHandler(uow iuow.UnitOfWork) CreateResourceHandler {
	if uow == nil {
		panic("NewCreateResourcesHandler requires uow")
	}

	return CreateResourceHandler{uow}
}

func (h CreateResourcesHandler) Handle(ctx echox.Context, cmds CreateResources) (err error) {
	defer func() {
		cqrs.LogCommandExecution(inject.GetLogger(), utils.GetTypeName(h), cmds, err)
	}()

	var (
		uow          = h.uow.WithContext(ctx.RequestContext())
		repoResource = repository.NewResourceRepository(uow.DB())
		parentsMap   map[uint]*domain.Resource
		slugs        []string
		parentIDs    []uint
		parents      domain.Resources
	)

	for _, cmd := range cmds {
		slugs = append(slugs, cmd.Slug)
		if cmd.ParentID > 0 {
			parentIDs = append(parentIDs, cmd.ParentID)
		}
	}

	err = gormx.FindWrapError(repoResource.Query(map[string][]interface{}{"slug": {slugs}}, nil).Select("slug"), &slugs)
	if err != nil {
		return err
	}

	if len(slugs) != 0 {
		return fmt.Errorf("%v: %s", errors.ERR_RESOURCE_SLUG_EXISTS, strings.Join(slugs, ","))
	}

	if len(parentIDs) != 0 {
		err = repoResource.FindID(&parents, parentIDs, nil)
		if err != nil {
			return err
		}

		if len(parents) != 0 {
			parentsMap = parents.ToMap()
		}
	}

	if err = uow.TransactionUnitOfWork(func(ux iuow.UnitOfWork) error {
		for _, cmd := range cmds {
			if err = CreateResources([]*CreateResource{cmd}).CreateRecursive(ux, parentsMap[cmd.ParentID]); err != nil {
				return err
			}
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}
