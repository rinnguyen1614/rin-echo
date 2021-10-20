package command

import (
	"rin-echo/admin/adapters"
	"rin-echo/admin/adapters/repository"
	"rin-echo/admin/domain"
	"rin-echo/admin/errors"
	"rin-echo/admin/inject"
	"rin-echo/common/cqrs"
	echox "rin-echo/common/echo"
	iuow "rin-echo/common/uow/interfaces"
	"rin-echo/common/utils"
)

type CreateRole struct {
	cqrs.CreateCommand

	Name      string `validate:"required,min=5"`
	Slug      string `validate:"min=6"`
	IsStatic  bool   `json:"is_static"`
	IsDefault bool   `json:"is_default"`
	MenuIDs   []uint `json:"menu_ids"`
}

func (r CreateRole) NewRole() (domain.Role, error) {
	return domain.NewRole(r.Name, r.Slug, r.IsStatic, r.IsDefault)
}

type CreateRoleHandler struct {
	uow  iuow.UnitOfWork
	rbac *adapters.RBACCasbin
}

func NewCreateRoleHandler(uow iuow.UnitOfWork, rbac *adapters.RBACCasbin) CreateRoleHandler {
	if uow == nil {
		panic("NewCreateRoleHandler requires uow")
	}

	return CreateRoleHandler{uow, rbac}
}

func (h CreateRoleHandler) Handle(ctx echox.Context, cmd CreateRole) (err error) {
	defer func() {
		cqrs.LogCommandExecution(inject.GetLogger(), utils.GetTypeName(h), cmd, err)
	}()

	var (
		uow      = h.uow.WithContext(ctx.RequestContext())
		repoRole = repository.NewRoleRepository(uow.DB())
		repoMenu = repository.NewMenuRepository(uow.DB())
		menus    []*domain.Menu
	)

	if _, err = repoRole.FirstBySlug(cmd.Slug, nil); err == nil {
		return errors.ERR_ROLE_SLUG_EXISTS
	}

	if len(cmd.MenuIDs) != 0 {
		if err = repoMenu.GetID(&menus, cmd.MenuIDs, map[string][]interface{}{"Resources": nil}); err != nil {
			return err
		}
	}

	if err = uow.TransactionUnitOfWork(func(ux iuow.UnitOfWork) error {
		var (
			r              domain.Role
			uow            = ux
			repoRole       = repository.NewRoleRepository(uow.DB())
			repoPermission = repository.NewPermissionRepository(uow.DB())
		)

		r, err = cmd.NewRole()
		if err != nil {
			return err
		}

		if err = repoRole.Create(&r); err != nil {
			return err
		}

		if _, err = h.rbac.AddRole(utils.ToString(r.ID)); err != nil {
			return err
		}

		if len(menus) != 0 {
			var (
				permissions domain.Permissions
				resources   domain.Resources
			)

			for _, m := range menus {
				p, _ := domain.NewPermissionForRole(m.ID, r.ID)
				permissions = append(permissions, &p)
				resources = append(resources, m.Resources...)
			}
			if err = repoPermission.Create(&permissions); err != nil {
				return err
			}

			// casbin

			for _, re := range resources {
				if _, err = h.rbac.AddPermissionForRole(utils.ToString(r.ID), re.Path, re.Method); err != nil {
					return err
				}
			}
		}

		return nil
	}); err != nil {
		return err
	}
	return nil
}
