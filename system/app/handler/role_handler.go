package handler

import (
	echox "rin-echo/common/echo"
	rquery "rin-echo/common/echo/models/query/rest_query"
	"rin-echo/common/model"
	iuow "rin-echo/common/uow/interfaces"

	"rin-echo/system/app/model/request"
	"rin-echo/system/app/service"
	"rin-echo/system/domain"

	"go.uber.org/zap"
)

type RoleHandler struct {
	echox.Handler
	service service.RoleService
}

func NewRoleHandler(uow iuow.UnitOfWork, permissionManager domain.PermissionManager, logger *zap.Logger, restQuery rquery.RestQuery) RoleHandler {
	return RoleHandler{
		Handler: echox.NewHandler(logger, restQuery),
		service: service.NewRoleService(uow, permissionManager, logger),
	}
}

func (h RoleHandler) Create(c echox.Context) error {
	var cmd request.CreateRole
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

func (h RoleHandler) Update(c echox.Context) error {
	var cmd request.UpdateRole
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

func (h RoleHandler) Delete(c echox.Context) error {
	// id, err := CheckRequestIDParam(c.Param("id"))
	// err = h.service.WithContext(c).Delete(id)
	// if err = h.service.WithContext(c).Delete(id); err != nil {
	// 	return err
	// }

	// echox.OKWithData(c, nil)
	return nil
}

func (h RoleHandler) Get(c echox.Context) error {
	// id, err := CheckRequestIDParam(c.Param("id"))
	// result, err := h.service.WithContext(c).Get(id)
	// if err != nil {
	// 	return err
	// }

	// echox.OKWithData(c, result)
	return nil
}