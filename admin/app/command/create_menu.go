package command

import (
	"rin-echo/admin/adapters/repository"
	"rin-echo/admin/domain"
	"rin-echo/admin/errors"
	"rin-echo/admin/inject"
	"rin-echo/common/cqrs"
	echox "rin-echo/common/echo"
	iuow "rin-echo/common/uow/interfaces"
	"rin-echo/common/utils"
)

type CreateMenu struct {
	cqrs.CreateCommand

	Name        string `validate:"required,min=5"`
	Slug        string `validate:"min=6"`
	ParentID    uint   `json:"parent_id"`
	Path        string `validate:"required,min=6"`
	Hidden      bool
	Component   string
	Sort        int
	Type        string
	Title       string
	Icon        string
	ResourceIds []uint `json:"resource_ids"`
	Children    []*CreateMenu
}

func (m CreateMenu) NewMenu() (domain.Menu, error) {
	menu, err := domain.NewMenu(m.Name, m.Slug, m.Path, m.Hidden, m.Component, m.Sort)
	if m.Type != "" && err != nil {
		menu.SetType(m.Type)
	}
	return menu, err
}

func (r CreateMenu) CheckSlug(repo *repository.MenuRepository) error {
	if ok := repo.Contains(map[string][]interface{}{"slug": {r.Slug}}); ok {
		return errors.ERR_RESOURCE_SLUG_EXISTS
	}

	return nil
}

func (m CreateMenu) CheckParent(repo *repository.MenuRepository) (*domain.Menu, error) {
	var parent *domain.Menu
	if m.ParentID != 0 {
		if err := repo.FirstID(&parent, m.ParentID, nil); err != nil {
			return nil, err
		}
	}
	return parent, nil
}

func (m CreateMenu) CheckResources(repo *repository.ResourceRepository) (domain.Resources, error) {
	var resources domain.Resources
	if len(m.ResourceIds) != 0 {
		if err := repo.GetID(&resources, m.ResourceIds, nil); err != nil {
			return nil, err
		}
	}
	return resources, nil
}

func (m CreateMenu) CreateRecursive(uow iuow.UnitOfWork, parent *domain.Menu, resources domain.Resources) error {
	var (
		repoMenu = repository.NewMenuRepository(uow.DB())
	)

	var (
		menu, err = m.NewMenu()
	)

	if err != nil {
		return err
	}

	menu.SetMeta(m.Title, m.Icon)
	menu.SetResources(resources)
	menu.SetParent(parent)

	if err = repoMenu.Create(&menu); err != nil {
		return err
	}

	m.ID = menu.ID

	for _, m := range m.Children {
		if err := m.CreateRecursive(uow, parent, resources); err != nil {
			return err
		}
	}

	return nil
}

type CreateMenuHandler struct {
	uow iuow.UnitOfWork
}

func NewCreateMenuHandler(uow iuow.UnitOfWork) CreateMenuHandler {
	if uow == nil {
		panic("NewCreateMenuHandler requires uow")
	}

	return CreateMenuHandler{uow}
}

func (h CreateMenuHandler) Handle(ctx echox.Context, cmd *CreateMenu) (err error) {
	defer func() {
		cqrs.LogCommandExecution(inject.GetLogger(), utils.GetTypeName(h), cmd, err)
	}()

	var (
		uow          = h.uow.WithContext(ctx.RequestContext())
		repoMenu     = repository.NewMenuRepository(uow.DB())
		repoResource = repository.NewResourceRepository(uow.DB())
		parent       *domain.Menu
		resources    domain.Resources
	)

	if err := cmd.CheckSlug(repoMenu); err != nil {
		return err
	}

	if parent, err = cmd.CheckParent(repoMenu); err != nil {
		return err
	}

	if resources, err = cmd.CheckResources(repoResource); err != nil {
		return err
	}

	if err = uow.TransactionUnitOfWork(func(ux iuow.UnitOfWork) error {
		return cmd.CreateRecursive(ux, parent, resources)
	}); err != nil {
		return err
	}

	return nil
}
