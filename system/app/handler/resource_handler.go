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

type ResourceHandler struct {
	echox.Handler
	service service.ResourceService
}

func NewResourceHandler(uow iuow.UnitOfWork,
	permissionManager domain.PermissionManager,
	logger *zap.Logger,
	restQuery rquery.RestQuery,
	settingProvider setting.Provider,
	validator *validation.Validator) ResourceHandler {

	return ResourceHandler{
		Handler: echox.NewHandler(logger, restQuery, settingProvider, validator),
		service: service.NewResourceService(uow, permissionManager, settingProvider, logger),
	}
}

// CreateResource godoc
// @Summary      Create a new resource
// @Description  Create a new resource with the input payload
// @Tags         resources
// @Accept       application/json
// @Produce      application/json
// @Param 		 data body request.CreateResource true "Create resource"
// @Success      200  {object}  models.Response{data=model.Model} "{"data": {}}"
// @Router       /resources [post]
// @Security Bearer
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

// UpdateResource godoc
// @Summary      Update resource identified by the given id
// @Description  Update the resource corresponding to the input id
// @Tags         resources
// @Accept       application/json
// @Produce      application/json
// @Param 		 id path int true "ID of the resource to be updated"
// @Param 		 data body request.UpdateResource true "Update resource"
// @Success      200  {object}  models.Response{data=model.Model} "{"data": {}}"
// @Router       /resources/{id} [put]
// @Security Bearer
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

// DeleteResource godoc
// @Summary      Delete resource identified by the given id
// @Description  Delete the resource corresponding to the input id
// @Tags         resources
// @Accept       application/json
// @Produce      application/json
// @Param 		 id path int true "ID of the resource to be deleted"
// @Success      200  {object}  models.Response "{"data": {}}"
// @Router       /resources/{id} [delete]
// @Security Bearer
func (h ResourceHandler) Delete(c echox.Context) error {
	id, err := CheckRequestIDParam(c.Param("id"))
	err = h.service.WithContext(c).Delete(id)
	if err = h.service.WithContext(c).Delete(id); err != nil {
		return err
	}

	echox.OKWithData(c, nil)
	return nil
}

// GetResource godoc
// @Summary 	Get details for a given id
// @Description Get details of resource corresponding to the input id
// @Tags 		resources
// @Accept  	application/json
// @Produce  	application/json
// @Param 		id path int true "ID of the resource"
// @Success     200  {object} models.Response{data=response.Resource} "{"data": {}}"
// @Router 		/resources/{id} [get]
// @Security Bearer
func (h ResourceHandler) Get(c echox.Context) error {
	id, err := CheckRequestIDParam(c.Param("id"))
	result, err := h.service.WithContext(c).Get(id)
	if err != nil {
		return err
	}

	echox.OKWithData(c, result)
	return nil
}

// GetResources godoc
// @Summary 	Get details of all resources
// @Description Get details of all resources
// @Tags 		resources
// @Accept  	application/json
// @Produce  	application/json
// @Param 		page_size query int true "pageSize"
// @Param 		page query int true "page"
// @Param 		filters query string false "filters separated by ",""
// @Param 		selects query string false "selects separated by ",""
// @Param 		sorts query string false "sorts separated by ",""
// @Success     200  {object} models.Response{data=model.QueryResult{records=response.ResourceTrees}} "{"data": {}}"
// @Router 		/resources/trees [get]
// @Security Bearer
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
