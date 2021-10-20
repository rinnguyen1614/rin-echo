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

type CreateResource struct {
	cqrs.CreateCommand

	Name        string `validate:"required,min=5"`
	Slug        string `validate:"required,min=6"`
	ParentID    uint   `json:"parent_id"`
	Path        string
	Method      string
	Description string
	Children    []*CreateResource
}

func (r CreateResource) NewResource() (domain.Resource, error) {
	return domain.NewResource(r.Name, r.Slug, r.Path, r.Method, r.Description)
}

func (r CreateResource) CheckSlug(repo *repository.ResourceRepository) error {
	if ok := repo.Contains(map[string][]interface{}{"slug": {r.Slug}}); ok {
		return errors.ERR_RESOURCE_SLUG_EXISTS
	}

	return nil
}

func (m CreateResource) CheckParent(repo *repository.ResourceRepository) (*domain.Resource, error) {
	var parent *domain.Resource
	if m.ParentID != 0 {
		if err := repo.FirstID(&parent, m.ParentID, nil); err != nil {
			return nil, err
		}
	}
	return parent, nil
}

func (cmd CreateResource) CreateRecursive(uow iuow.UnitOfWork, parent *domain.Resource) error {
	var repoResource = repository.NewResourceRepository(uow.DB())
	resource, err := cmd.NewResource()
	if err != nil {
		return err
	}

	resource.SetParent(parent)
	if err = repoResource.Create(&resource); err != nil {
		return err
	}

	cmd.ID = resource.ID

	if len(cmd.Children) != 0 {
		if err = CreateResources(cmd.Children).CreateRecursive(uow, &resource); err != nil {
			return err
		}
	}
	return nil
}

type CreateResourceHandler struct {
	uow iuow.UnitOfWork
}

func NewCreateResourceHandler(uow iuow.UnitOfWork) CreateResourceHandler {
	if uow == nil {
		panic("NewCreateResourceHandler requires uow")
	}

	return CreateResourceHandler{uow}
}

func (h CreateResourceHandler) Handle(ctx echox.Context, cmd *CreateResource) (err error) {
	defer func() {
		cqrs.LogCommandExecution(inject.GetLogger(), utils.GetTypeName(h), cmd, err)
	}()

	var (
		uow          = h.uow.WithContext(ctx.RequestContext())
		repoResource = repository.NewResourceRepository(uow.DB())
		parent       *domain.Resource
	)

	if err = cmd.CheckSlug(repoResource); err != nil {
		return err
	}

	if parent, err = cmd.CheckParent(repoResource); err != nil {
		return err
	}

	if err = uow.TransactionUnitOfWork(func(ux iuow.UnitOfWork) error {
		return cmd.CreateRecursive(ux, parent)
	}); err != nil {
		return err
	}

	return nil
}
