package service

import (
	"github.com/rinnguyen1614/rin-echo/internal/system/adapters/repository"
	"github.com/rinnguyen1614/rin-echo/internal/system/app/model/request"
	"github.com/rinnguyen1614/rin-echo/internal/system/app/model/response"
	querybuilder "github.com/rinnguyen1614/rin-echo/internal/system/domain/query_builder"

	echox "github.com/rinnguyen1614/rin-echo/internal/core/echo"
	"github.com/rinnguyen1614/rin-echo/internal/core/model"
	"github.com/rinnguyen1614/rin-echo/internal/core/query"
	"github.com/rinnguyen1614/rin-echo/internal/core/setting"
	"github.com/rinnguyen1614/rin-echo/internal/core/uow"
	iuow "github.com/rinnguyen1614/rin-echo/internal/core/uow/interfaces"
	"github.com/rinnguyen1614/rin-echo/internal/core/utils"

	"github.com/rinnguyen1614/rin-echo/internal/system/domain"

	"github.com/jinzhu/copier"
	"github.com/rinnguyen1614/rin-echo/internal/system/errors"
	"go.uber.org/zap"
)

type (
	UserService interface {
		WithContext(echox.Context) UserService

		Create(request.CreateUser) (uint, error)

		// with roles are is_default
		CreateDefault(cmd request.CreateUser) (uint, error)

		Update(id uint, cmd request.UpdateUser) (err error)

		Delete(id uint) (err error)

		Get(id uint) (response.User, error)

		Query(q *query.Query) (*model.QueryResult, error)

		// 	// Update()

		// 	// Disable()

		// 	// Enable()

		// 	// ChangePassword()

		// 	// AssignToRoles(id uint, roleIDs []uint)

		// 	// AssignToRole(id, roleID uint)

		// 	// FindByUsernameOrEmail(usernameOrEmail string) error

	}

	userService struct {
		*echox.Service

		permissionManager domain.PermissionManager
		repo              domain.UserRepository
		repoRole          domain.RoleRepository
		repoUserRole      domain.UserRoleRepository
	}
)

func NewUserService(uow iuow.UnitOfWork, permissionManager domain.PermissionManager, settingProvider setting.Provider, logger *zap.Logger) UserService {
	return &userService{
		Service: echox.NewService(uow, settingProvider, logger),

		permissionManager: permissionManager,
		repo:              repository.NewUserRepository(uow.DB()),
		repoRole:          repository.NewRoleRepository(uow.DB()),
		repoUserRole:      repository.NewUserRoleRepository(uow.DB()),
	}
}

func (s *userService) WithContext(ctx echox.Context) UserService {
	return &userService{
		Service: s.Service.WithContext(ctx),

		permissionManager: s.permissionManager,
		repo:              s.repo.WithTransaction(s.Service.Uow.DB()),
		repoRole:          s.repoRole.WithTransaction(s.Service.Uow.DB()),
		repoUserRole:      s.repoUserRole.WithTransaction(s.Service.Uow.DB()),
	}
}

func (s userService) Create(cmd request.CreateUser) (uint, error) {
	var id uint
	if err := s.Uow.TransactionUnitOfWork(func(ux iuow.UnitOfWork) error {
		var repo = s.repo.WithTransaction(ux.DB())

		if err := s.CheckExistByUsername(cmd.Username); err != nil {
			return err
		}

		if err := s.CheckExistByEmail(cmd.Email); err != nil {
			return err
		}

		if cmd.IsRandomPassword() {
			cmd.Password = generatePassword()
		}

		user, err := domain.NewUser(cmd.Username, cmd.Password, cmd.FullName, cmd.Email, cmd.RoleIDs)
		if err != nil {
			return err
		}

		if err := repo.Create(user); err != nil {
			return err
		}

		for _, uR := range user.UserRoles {
			if _, err := s.permissionManager.AddRoleForUser(user.ID, uR.RoleID); err != nil {
				return err
			}
		}

		id = user.ID
		return nil
	}); err != nil {
		return 0, err
	}

	return id, nil
}

func (s userService) CreateDefault(cmd request.CreateUser) (uint, error) {
	var roleIDs []uint
	if err := uow.Find(s.repoRole.Query(map[string][]interface{}{"is_default": {true}}, nil).Select("id"), &roleIDs); err != nil {
		return 0, err
	}
	cmd.RoleIDs = roleIDs
	return s.Create(cmd)
}

func (s userService) Update(id uint, cmd request.UpdateUser) (err error) {
	if err = s.Uow.TransactionUnitOfWork(func(ux iuow.UnitOfWork) error {

		var (
			repo = s.repo.WithTransaction(ux.DB())
			user domain.User
		)
		err := repo.GetID(&user, id, nil)
		if err != nil {
			return err
		}

		// if cmd.Username != user.Username {
		// 	if user.IsGlobalAdmin {
		// 		return errors.ErrCannotChangeUsernameOfAdmin
		// 	}
		// 	if err := s.CheckExistByUsername(cmd.Username); err != nil {
		// 		return err
		// 	}
		// }

		if cmd.Email != user.Email {
			if err = s.CheckExistByEmail(cmd.Email); err != nil {
				return err
			}
		}

		if cmd.RandomPassword {
			cmd.Password = generatePassword()
		}

		if cmd.Password != "" {
			if err := repo.UpdatePassword(&user, cmd.Password); err != nil {
				return err
			}
		}

		if err = repo.UpdateWithPrimaryKey(id, map[string]interface{}{
			"full_name": cmd.FullName,
			"email":     cmd.Email,
		}); err != nil {
			return err
		}

		return s.SetRoles(&user, cmd.RoleIDs)
	}); err != nil {
		return err
	}

	return nil
}

func (s userService) SetRoles(user *domain.User, roleIDs []uint) (err error) {

	if user == nil {
		panic("requires user")
	}

	if err = s.Uow.TransactionUnitOfWork(func(ux iuow.UnitOfWork) error {
		var (
			repoUserRole = s.repoUserRole.WithTransaction(ux.DB())
			repo         = s.repo.WithTransaction(ux.DB())
			newUserRoles domain.UserRoles
		)

		err := repo.Find(user, nil, map[string][]interface{}{"UserRoles": nil})
		if err != nil {
			return err
		}

		for _, roleID := range roleIDs {
			uR, _ := domain.NewUserRole(user.ID, roleID)
			newUserRoles = append(newUserRoles, uR)
		}

		userRoleNews, userRoleDels := user.CompareUserRoles(newUserRoles)

		// remove from removed roles
		if len(userRoleDels) != 0 {
			if err = repoUserRole.DeleteMany(userRoleDels.IDs()); err != nil {
				return err
			}

			for _, uR := range userRoleDels {
				if _, err := s.permissionManager.DeleteRoleForUser(user.ID, uR.RoleID); err != nil {
					return err
				}

			}
		}

		// add to added roles
		if len(userRoleNews) != 0 {
			if err = repoUserRole.Create(userRoleNews); err != nil {
				return err
			}

			for _, uR := range userRoleNews {
				if _, err := s.permissionManager.AddRoleForUser(user.ID, uR.RoleID); err != nil {
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

func (s userService) CheckExistByUsername(username string) error {
	if ok, _ := s.repo.Contains(map[string][]interface{}{"username": {username}}); ok {
		return errors.ErrUsernameExists
	}

	return nil
}

func (s userService) CheckExistByEmail(email string) error {
	if ok, _ := s.repo.Contains(map[string][]interface{}{"email": {email}}); ok {
		return errors.ErrEmailExists
	}

	return nil
}

func (s userService) ResetPassword(user *domain.User) error {
	if user == nil {
		panic("requires user")
	}

	newPassword := generatePassword()
	if err := s.repo.UpdatePassword(user, newPassword); err != nil {
		return err
	}

	// send mail

	return nil
}

func (s userService) Delete(id uint) (err error) {
	return s.Uow.TransactionUnitOfWork(func(ux iuow.UnitOfWork) error {
		var (
			repo             = s.repo.WithTransaction(ux.DB())
			hasRole, _       = uow.Contains(ux.DB().Table("user_roles").Where("user_id", id))
			isGlobalAdmin, _ = repo.Contains(map[string][]interface{}{"id": {id}, "is_global_admin": {true}})
		)

		if isGlobalAdmin {
			return errors.ErrCannotDeleteAdmin
		}

		if hasRole {
			return errors.ErrResourceReferencedRole
		}

		if err := repo.Delete(id); err != nil {
			return err
		}

		return nil
	})
}

func (s userService) Get(id uint) (response.User, error) {
	var (
		user domain.User
		res  response.User
	)
	if err := s.repo.GetID(&user, id, map[string][]interface{}{"UserRoles": nil, "UserRoles.Role": nil}); err != nil {
		return response.User{}, err
	}

	if err := copier.CopyWithOption(&res, user, copier.Option{IgnoreEmpty: true, DeepCopy: true}); err != nil {
		return response.User{}, err
	}
	return res, nil
}

func (s userService) Query(q *query.Query) (*model.QueryResult, error) {
	var (
		queryBuilder    = querybuilder.NewUserQueryBuilder()
		preloadBuilders = map[string]iuow.QueryBuilder{
			"UserRoles": querybuilder.NewUserRoleQueryBuilder(),
			"Role":      querybuilder.NewRoleQueryBuilder(),
		}
	)

	return q.QueryResult(s.repo, queryBuilder, preloadBuilders, domain.User{}, response.User{})
}

func generatePassword() string {
	return utils.RandomSymbol(16)
}
