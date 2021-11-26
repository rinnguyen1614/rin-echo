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

type ResourceHandler struct {
	echox.Handler
	service service.ResourceService
}

func NewResourceHandler(uow iuow.UnitOfWork, permissionManager domain.PermissionManager, logger *zap.Logger, restQuery rquery.RestQuery) ResourceHandler {
	return ResourceHandler{
		Handler: echox.NewHandler(logger, restQuery),
		service: service.NewResourceService(uow, permissionManager, logger),
	}
}

func (h ResourceHandler) Create(c echox.Context) error {
	var cmd request.CreateResource
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

func (h ResourceHandler) Update(c echox.Context) error {
	var cmd request.UpdateResource
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

func (h ResourceHandler) Delete(c echox.Context) error {
	id, err := CheckRequestIDParam(c.Param("id"))
	err = h.service.WithContext(c).Delete(id)
	if err = h.service.WithContext(c).Delete(id); err != nil {
		return err
	}

	echox.OKWithData(c, nil)
	return nil
}

func (h ResourceHandler) Get(c echox.Context) error {
	id, err := CheckRequestIDParam(c.Param("id"))
	result, err := h.service.WithContext(c).Get(id)
	if err != nil {
		return err
	}

	echox.OKWithData(c, result)
	return nil
}

func (h ResourceHandler) TreeQuery(c echox.Context) error {
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