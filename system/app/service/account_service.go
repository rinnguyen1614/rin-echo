package service

import (
	"rin-echo/common/auth/jwt"
	echox "rin-echo/common/echo"
	"rin-echo/common/setting"
	"rin-echo/common/uow"
	iuow "rin-echo/common/uow/interfaces"
	"rin-echo/system/adapters/repository"
	"rin-echo/system/app/model/request"
	"rin-echo/system/app/model/response"
	"rin-echo/system/domain"
	"rin-echo/system/errors"
	"time"

	"github.com/jinzhu/copier"
	"go.uber.org/zap"
)

type (
	AccountService interface {
		Login(request.Login) (interface{}, error)

		ChangePassword(cmd request.ChangePassword) (interface{}, error)

		Token(id uint) (interface{}, error)

		Profile(id uint) (response.Profile, error)

		FindMenuTrees(userId uint) (response.UserMenus, error)

		WithContext(ctx echox.Context) AccountService
	}

	accountService struct {
		*echox.Service

		auther       *jwt.JWT
		repo         domain.UserRepository
		repoSecurity domain.SecurityLogRepository
	}
)

func NewAccountService(uow iuow.UnitOfWork, settingProvider setting.Provider, logger *zap.Logger, auther *jwt.JWT) AccountService {
	return &accountService{
		Service:      echox.NewService(uow, settingProvider, logger),
		auther:       auther,
		repo:         repository.NewUserRepository(uow.DB()),
		repoSecurity: repository.NewSecurityLogRepository(uow.DB()),
	}
}

func (s *accountService) WithContext(ctx echox.Context) AccountService {
	return &accountService{
		Service:      s.Service.WithContext(ctx),
		auther:       s.auther,
		repo:         s.repo.WithTransaction(s.Service.Uow.DB()),
		repoSecurity: s.repoSecurity.WithTransaction(s.Service.Uow.DB()),
	}
}

func (s accountService) Login(cmd request.Login) (interface{}, error) {

	defer func() {
		s.createSecurityLog(cmd.Username, "login")
	}()

	user, err := s.getUserByUserNameAndPassword(cmd.Username, cmd.Password)

	if err != nil {
		return nil, err
	}

	return s.token(user)
}

func (s accountService) createSecurityLog(username, message string) {
	var (
		ctx        = s.Context()
		location   string
		ipAddress  = ctx.RealIP()
		userAgent  = ctx.Request().UserAgent()
		deviceID   = ctx.Request().Header.Get(echox.HeaderDeviceID)
		deviceName = ctx.Request().Header.Get(echox.HeaderDeviceName)
		time       = time.Now()
		statusCode int
	)

	loginLog := domain.NewSecurityLog(
		username,
		location,
		ipAddress,
		userAgent,
		deviceID,
		deviceName,
		time,
		statusCode,
		message,
	)

	s.repoSecurity.Create(loginLog)
}

func (s accountService) ChangePassword(cmd request.ChangePassword) (interface{}, error) {
	defer func() {
		s.createSecurityLog(cmd.Username, "change_password")
	}()

	user, err := s.getUserByUserNameAndPassword(cmd.Username, cmd.CurrentPassword)

	if err != nil {
		return nil, err
	}

	if err = s.repo.UpdatePassword(user, cmd.NewPassword); err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	return s.token(user)
}

func (s accountService) Token(id uint) (interface{}, error) {
	var user domain.User
	if err := s.repo.GetID(&user, id, nil); err != nil {
		return nil, err
	}

	return s.token(&user)
}

func (s accountService) Profile(id uint) (response.Profile, error) {
	var user domain.User
	if err := s.repo.GetID(&user, id, nil); err != nil {
		return response.Profile{}, err
	}

	return response.NewProfile(user), nil
}

func (s accountService) UpdateProfile(id uint, cmd request.UpdateProfile) error {

	var user domain.User
	if err := s.repo.GetID(&user, id, nil); err != nil {
		return err
	}

	defer func() {
		s.createSecurityLog(user.Username, "update_profile")
	}()
	return nil
}

func (s accountService) FindMenuTrees(userId uint) (response.UserMenus, error) {
	var (
		repoMenu = repository.NewMenuRepository(s.Uow.DB())
		menus    domain.Menus
		result   response.UserMenus
	)
	if err := uow.Find(repoMenu.QueryByUsers([]uint{userId}, nil), &menus); err != nil {
		return nil, err
	}

	if err := copier.CopyWithOption(&result, menus.ToTree(), copier.Option{IgnoreEmpty: true, DeepCopy: true}); err != nil {
		return nil, err
	}

	return result, nil
}

func (s accountService) token(user *domain.User) (interface{}, error) {
	token, err := s.auther.Token(s.Context().RequestContext(), map[string]interface{}{
		"FullName": user.FullName,
		"Email":    user.Email,
		"ID":       user.ID,
		"UUID":     user.UUID,
		"Username": user.Username,
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

func (s accountService) getUserByUserNameAndPassword(username, password string) (*domain.User, error) {
	user, err := s.repo.FirstByUsernameOrEmail(username, nil)
	if err != nil {
		return nil, errors.ErrUserNotFound
	}

	if !user.CheckPassword(password) {
		return nil, errors.ErrUserNamePasswordNotMatch
	}
	return user, nil
}

func (s accountService) checkPassword(password string) {

}
