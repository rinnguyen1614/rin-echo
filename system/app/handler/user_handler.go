package handler

import (
	echox "rin-echo/common/echo"
	"rin-echo/common/model"
	"rin-echo/common/setting"
	iuow "rin-echo/common/uow/interfaces"
	"rin-echo/common/validation"

	"rin-echo/system/app/model/request"
	"rin-echo/system/app/service"
	"rin-echo/system/domain"

	"go.uber.org/zap"

	rquery "rin-echo/common/echo/models/query/rest_query"
)

type UserHandler struct {
	echox.Handler
	service service.UserService
}

func NewUserHandler(uow iuow.UnitOfWork,
	permissionManager domain.PermissionManager,
	logger *zap.Logger,
	restQuery rquery.RestQuery,
	settingProvider setting.Provider,
	validator *validation.Validator) UserHandler {
	return UserHandler{
		Handler: echox.NewHandler(logger, restQuery, settingProvider, validator),
		service: service.NewUserService(uow, permissionManager, settingProvider, logger),
	}
}

func (h UserHandler) Create(c echox.Context) error {
	var cmd request.CreateUser
	if err := c.Bind(&cmd); err != nil {
		return err
	}
	if err := c.Validate(cmd); err != nil {
		return err
	}
	id, err := h.service.WithContext(c).Create(cmd)
	if err != nil {
		return err
	}

	echox.OKWithData(c, model.NewModel(id))
	return nil
}

func (h UserHandler) Update(c echox.Context) error {
	var cmd request.UpdateUser
	if err := c.Bind(&cmd); err != nil {
		return err
	}
	if err := c.Validate(cmd); err != nil {
		return err
	}

	id, err := CheckRequestIDParam(c.Param("id"))
	if err != nil {
		return err
	}
	err = h.service.WithContext(c).Update(id, cmd)
	if err != nil {
		return err
	}

	echox.OKWithData(c, model.NewModel(id))
	return nil
}

func (h UserHandler) Query(c echox.Context) error {
	query, err := h.RestQuery.Query(c.Request())

	if err != nil {
		return err
	}
	result, err := h.service.WithContext(c).Query(query)
	if err != nil {
		return err
	}

	echox.OKWithData(c, result)

	return nil
}
