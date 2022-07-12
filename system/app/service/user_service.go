package service

import (
	"rin-echo/system/adapters/repository"
	"rin-echo/system/app/model/request"
	"rin-echo/system/app/model/response"
	"rin-echo/system/domain"
	querybuilder "rin-echo/system/domain/query_builder"
	"rin-echo/system/errors"

	echox "github.com/rinnguyen1614/rin-echo-core/echo"
	"github.com/rinnguyen1614/rin-echo-core/model"
	"github.com/rinnguyen1614/rin-echo-core/query"
	"github.com/rinnguyen1614/rin-echo-core/setting"
	"github.com/rinnguyen1614/rin-echo-core/uow"
	iuow "github.com/rinnguyen1614/rin-echo-core/uow/interfaces"
	"github.com/rinnguyen1614/rin-echo-core/utils"

	"github.com/jinzhu/copier"
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
	if err := s.CheckExistByUsername(cmd.Username); err != nil {
		return 0, err
	}

	if err := s.CheckExistByEmail(cmd.Email); err != nil {
		return 0, err
	}

	user, err := domain.NewUser(cmd.Username, cmd.Password, cmd.FullName, cmd.Email, cmd.RoleIDs)
	if err != nil {
		return 0, err
	}

	if err := s.repo.Create(user); err != nil {
		return 0, err
	}

	if err := s.ResetPassword(user); err != nil {
		return 0, err
	}

	for _, uR := range user.UserRoles {
		if _, err := s.permissionManager.AddRoleForUser(user.ID, uR.RoleID); err != nil {
			return 0, err
		}

	}
	// send mail
	return user.ID, nil
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

		if cmd.Email != user.Email {
			if err = s.CheckExistByEmail(cmd.Email); err != nil {
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
		return errors.ErrUserNameExists
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

	newPassword := utils.RandomSymbol(16)
	if err := s.repo.UpdatePassword(user, newPassword); err != nil {
		return err
	}

	// send mail

	return nil
}

func (s userService) Delete(id uint) (err error) {
	return s.Uow.TransactionUnitOfWork(func(ux iuow.UnitOfWork) error {
		var (
			repo       = s.repo.WithTransaction(ux.DB())
			hasRole, _ = uow.Contains(ux.DB().Table("user_roles").Where("user_id", id))
		)

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
