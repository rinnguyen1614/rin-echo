package command

import (
	"rin-echo/admin/adapters/repository"
	"rin-echo/admin/domain"
	"rin-echo/admin/inject"
	"rin-echo/common"
	"rin-echo/common/cqrs"
	echox "rin-echo/common/echo"
	iuow "rin-echo/common/uow/interfaces"
	"rin-echo/common/utils"
)

type KeyValue struct {
	Key   interface{}
	Value interface{}
}

type CreateMenus []*CreateMenu

func (ms CreateMenus) CreateRecursive(uow iuow.UnitOfWork, parent *domain.Menu, resources domain.Resources) error {
	for _, m := range ms {
		if err := m.CreateRecursive(uow, parent, resources); err != nil {
			return err
		}
	}

	return nil
}

type CreateMenusHandler struct {
	uow iuow.UnitOfWork
}

func NewCreateMenusHandler(uow iuow.UnitOfWork) CreateMenusHandler {
	if uow == nil {
		panic("NewCreateMenusHandler requires uow")
	}

	return CreateMenusHandler{uow}
}

func (h CreateMenusHandler) Handle(ctx echox.Context, cmds CreateMenus) (err error) {
	defer func() {
		cqrs.LogCommandExecution(inject.GetLogger(), utils.GetTypeName(h), cmds, err)
	}()

	var (
		uow              = h.uow.WithContext(ctx.RequestContext())
		repoMenu         = repository.NewMenuRepository(uow.DB())
		repoResource     = repository.NewResourceRepository(uow.DB())
		errors           map[int][]error
		parentsByIndex   map[int]*domain.Menu
		resourcesByIndex map[int]domain.Resources
	)

	for i, cmd := range cmds {
		var errs []error
		if err := cmd.CheckSlug(repoMenu); err != nil {
			errs = append(errs, err)
		}

		if parent, err := cmd.CheckParent(repoMenu); err != nil {
			errs = append(errs, err)
		} else {
			parentsByIndex[i] = parent
		}

		if resources, err := cmd.CheckResources(repoResource); err != nil {
			errs = append(errs, err)
		} else {
			resourcesByIndex[i] = resources
		}

		if len(errs) != 0 {
			errors[i] = errs
		}
	}

	if len(errors) != 0 {
		return common.NewRinErrors(errors, "create_menu", "You have some errors for create menu")
	}

	if err = uow.TransactionUnitOfWork(func(ux iuow.UnitOfWork) error {
		for i, cmd := range cmds {
			if err = cmd.CreateRecursive(ux, parentsByIndex[i], resourcesByIndex[i]); err != nil {
				return err
			}
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}
