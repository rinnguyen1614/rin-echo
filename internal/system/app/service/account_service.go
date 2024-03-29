package service

import (
	"mime/multipart"
	"strings"
	"time"

	"github.com/rinnguyen1614/rin-echo/internal/system/adapters/repository"
	"github.com/rinnguyen1614/rin-echo/internal/system/app/model/request"
	"github.com/rinnguyen1614/rin-echo/internal/system/app/model/response"

	"github.com/rinnguyen1614/rin-echo/internal/core/auth"
	echox "github.com/rinnguyen1614/rin-echo/internal/core/echo"
	"github.com/rinnguyen1614/rin-echo/internal/core/setting"
	iuow "github.com/rinnguyen1614/rin-echo/internal/core/uow/interfaces"
	"github.com/rinnguyen1614/rin-echo/internal/core/utils"
	"github.com/rinnguyen1614/rin-echo/internal/core/utils/upload"

	"github.com/jinzhu/copier"
	"go.uber.org/zap"

	"github.com/rinnguyen1614/rin-echo/internal/system/domain"
	"github.com/rinnguyen1614/rin-echo/internal/system/errors"
)

type (
	AccountService interface {
		Login(request.Login) (interface{}, error)

		ChangePassword(cmd request.ChangePassword) (interface{}, error)

		Token(id uint) (interface{}, error)

		Profile(id uint) (response.Profile, error)

		UpdateProfile(id uint, cmd request.UpdateProfile) error

		ChangeAvatar(id uint, file *multipart.FileHeader) (interface{}, error)

		ChangeEmail(id uint, cmd request.ChangeEmail) error

		VerifyEmail(id uint, cmd request.VerifyEmail) error

		ChangePhone(id uint, cmd request.ChangePhone) error

		VerifyPhone(id uint, cmd request.VerifyPhone) error

		FindMenuTrees(userID uint) (response.UserMenus, error)

		FindPermissions(userID uint) (response.UserPermissions, error)

		WithContext(ctx echox.Context) AccountService
	}

	accountService struct {
		*echox.Service

		auther       auth.Auther
		repo         domain.UserRepository
		repoSecurity domain.SecurityLogRepository
		upload       upload.Upload
	}
)

func NewAccountService(uow iuow.UnitOfWork, settingProvider setting.Provider, logger *zap.Logger, auther auth.Auther) AccountService {
	uploadMaxSize := setting.MustGet[int64](settingProvider, "files.upload.max_size")
	return &accountService{
		Service:      echox.NewService(uow, settingProvider, logger),
		auther:       auther,
		repo:         repository.NewUserRepository(uow.DB()),
		repoSecurity: repository.NewSecurityLogRepository(uow.DB()),
		upload:       upload.NewLocal(uploadMaxSize),
	}
}

func (s *accountService) WithContext(ctx echox.Context) AccountService {
	return &accountService{
		Service:      s.Service.WithContext(ctx),
		auther:       s.auther,
		repo:         s.repo.WithTransaction(s.Service.Uow.DB()),
		repoSecurity: s.repoSecurity.WithTransaction(s.Service.Uow.DB()),
		upload:       s.upload,
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

	return s.repo.UpdateProfile(id, cmd.FullName, cmd.DateOfBirth, cmd.Gender)
}

func (s accountService) ChangeAvatar(id uint, file *multipart.FileHeader) (interface{}, error) {

	var (
		user          domain.User
		basePath, err = s.SettingProvider.Get("files.avatar_path")
		path          string
	)
	if err != nil {
		return nil, err
	}

	if err = s.repo.GetID(&user, id, nil); err != nil {
		return nil, err
	}
	path = basePath + "/" + utils.Encrypt(user.Username, time.Now().Format("2006010250101")+utils.RandomLetter(2))
	fileUploaded, err := s.upload.Save(file, path)
	if err != nil {
		return nil, err
	}

	if err := s.repo.UpdateAvatar(user.ID, strings.TrimPrefix(fileUploaded.Path(), "/static")); err != nil {
		return nil, err
	}

	defer func() {
		s.createSecurityLog(user.Username, "change_avatar")
	}()

	return response.NewFile(*fileUploaded), nil
}

func (s accountService) ChangeEmail(id uint, cmd request.ChangeEmail) error {

	var (
		err  error
		user domain.User
		code = utils.RandomLetter(20)
	)

	if err = s.repo.GetID(&user, id, nil); err != nil {
		return err
	}

	if err := s.repo.ChangeEmail(user.ID, cmd.Email, utils.Encrypt(user.Username, code)); err != nil {
		return err
	}

	// send verification email via email_service
	defer func() {
		s.createSecurityLog(user.Username, "change_email")
	}()

	return nil
}

func (s accountService) VerifyEmail(id uint, cmd request.VerifyEmail) error {

	var (
		err  error
		user domain.User
	)

	if err = s.repo.GetID(&user, id, nil); err != nil {
		return err
	}

	if len(user.EmailVerificationCodeHashed) == 0 {
		return errors.ErrVericationEmail
	}

	if utils.Decrypt(user.Username, user.EmailVerificationCodeHashed) != cmd.Code {
		return errors.ErrVericationEmail
	}

	if err := s.repo.VerifyEmail(user.ID); err != nil {
		return err
	}

	defer func() {
		s.createSecurityLog(user.Username, "verify_email")
	}()

	return nil
}

func (s accountService) ChangePhone(id uint, cmd request.ChangePhone) error {

	var (
		err  error
		user domain.User
		code = utils.RandomNumeric(5)
	)

	if err = s.repo.GetID(&user, id, nil); err != nil {
		return err
	}

	if err := s.repo.ChangePhone(user.ID, cmd.Phone, utils.Encrypt(user.Username, code)); err != nil {
		return err
	}

	// send verification phone via phone_service
	defer func() {
		s.createSecurityLog(user.Username, "change_phone")
	}()

	return nil
}

func (s accountService) VerifyPhone(id uint, cmd request.VerifyPhone) error {

	var (
		err  error
		user domain.User
	)

	if err = s.repo.GetID(&user, id, nil); err != nil {
		return err
	}

	if len(user.PhoneVerificationCodeHashed) == 0 {
		return errors.ErrVericationPhone
	}

	if utils.Decrypt(user.Username, user.PhoneVerificationCodeHashed) != cmd.Code {
		return errors.ErrVericationPhone
	}

	if err := s.repo.VerifyPhone(user.ID); err != nil {
		return err
	}

	defer func() {
		s.createSecurityLog(user.Username, "verify_phone")
	}()

	return nil
}

func (s accountService) FindMenuTrees(userID uint) (response.UserMenus, error) {
	var (
		repoMenu   = repository.NewMenuRepository(s.Uow.DB())
		menus, err = repoMenu.FindByUser(userID, map[string][]interface{}{"hidden": {false}})
		result     response.UserMenus
	)
	if err != nil {
		return nil, err
	}

	if err := copier.CopyWithOption(&result, menus.ToTree(), copier.Option{IgnoreEmpty: true, DeepCopy: true}); err != nil {
		return nil, err
	}

	return result, nil
}

func (s accountService) FindPermissions(userID uint) (response.UserPermissions, error) {
	var (
		repoPermission = repository.NewPermissionRepository(s.Uow.DB())
		fields, err    = repoPermission.FindByUser(userID)
	)

	if err != nil {
		return nil, err
	}

	return response.NewUserPermissions(fields), nil
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
