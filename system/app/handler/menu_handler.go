package handler

import (
	echox "rin-echo/common/echo"
	rquery "rin-echo/common/echo/models/query/rest_query"
	"rin-echo/common/model"
	"rin-echo/common/setting"
	iuow "rin-echo/common/uow/interfaces"
	"rin-echo/common/validation"

	"rin-echo/system/app/model/request"
	"rin-echo/system/app/service"
	"rin-echo/system/domain"

	"go.uber.org/zap"
)

type MenuHandler struct {
	echox.Handler
	service service.MenuService
}

func NewMenuHandler(uow iuow.UnitOfWork,
	permissionManager domain.PermissionManager,
	logger *zap.Logger,
	restQuery rquery.RestQuery,
	settingProvider setting.Provider,
	validator *validation.Validator) MenuHandler {

	return MenuHandler{
		Handler: echox.NewHandler(logger, restQuery, settingProvider, validator),
		service: service.NewMenuService(uow, permissionManager, settingProvider, logger),
	}
}

func (h MenuHandler) Create(c echox.Context) error {
	var cmd request.CreateMenu
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

func (h MenuHandler) Update(c echox.Context) error {
	var cmd request.UpdateMenu
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

func (h MenuHandler) Delete(c echox.Context) error {
	// id, err := CheckRequestIDParam(c.Param("id"))
	// err = h.service.WithContext(c).Delete(id)
	// if err = h.service.WithContext(c).Delete(id); err != nil {
	// 	return err
	// }

	// echox.OKWithData(c, nil)
	return nil
}

func (h MenuHandler) Get(c echox.Context) error {
	id, err := CheckRequestIDParam(c.Param("id"))
	result, err := h.service.WithContext(c).Get(id)
	if err != nil {
		return err
	}

	echox.OKWithData(c, result)
	return nil
}

func (h MenuHandler) TreeQuery(c echox.Context) error {
	query, err := h.RestQuery.Query(c.Request())

	if err != nil {
		return err
	}
	result, err := h.service.WithContext(c).FindTrees(query)
	if err != nil {
		return err
	}

	echox.OKWithData(c, result)

	return nil
}
