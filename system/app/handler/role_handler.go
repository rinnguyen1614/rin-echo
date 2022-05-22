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

type RoleHandler struct {
	echox.Handler
	service service.RoleService
}

func NewRoleHandler(uow iuow.UnitOfWork,
	permissionManager domain.PermissionManager,
	logger *zap.Logger,
	restQuery rquery.RestQuery,
	settingProvider setting.Provider,
	validator *validation.Validator) RoleHandler {

	return RoleHandler{
		Handler: echox.NewHandler(logger, restQuery, settingProvider, validator),
		service: service.NewRoleService(uow, permissionManager, settingProvider, logger),
	}
}

// CreateRole godoc
// @Summary      Create a new role
// @Description  Create a new role with the input payload
// @Tags         roles
// @Accept       application/json
// @Produce      application/json
// @Param 		 data body request.CreateRole true "Create role"
// @Success      200  {object}  models.Response{data=model.Model} "{"data": {}}"
// @Router       /roles [post]
// @Security Bearer
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

// UpdateRole godoc
// @Summary      Update role identified by the given id
// @Description  Update the role corresponding to the input id
// @Tags         roles
// @Accept       application/json
// @Produce      application/json
// @Param 		 id path int true "ID of the role to be updated"
// @Param 		 data body request.UpdateRole true "Update role"
// @Success      200  {object}  models.Response{data=model.Model} "{"data": {}}"
// @Router       /roles/{id} [put]
// @Security Bearer
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

// DeleteRole godoc
// @Summary      Delete role identified by the given id
// @Description  Delete the role corresponding to the input id
// @Tags         roles
// @Accept       application/json
// @Produce      application/json
// @Param 		 id path int true "ID of the role to be deleted"
// @Success      200  {object}  models.Response "{"data": {}}"
// @Router       /roles/{id} [delete]
// @Security Bearer
func (h RoleHandler) Delete(c echox.Context) error {
	id, err := CheckRequestIDParam(c.Param("id"))
	err = h.service.WithContext(c).Delete(id)
	if err = h.service.WithContext(c).Delete(id); err != nil {
		return err
	}

	echox.OKWithData(c, nil)
	return nil
}

// GetRole godoc
// @Summary 	Get details for a given id
// @Description Get details of role corresponding to the input id
// @Tags 		roles
// @Accept  	application/json
// @Produce  	application/json
// @Param 		id path int true "ID of the role"
// @Success     200  {object} models.Response{data=response.Role} "{"data": {}}"
// @Router 		/roles/{id} [get]
// @Security Bearer
func (h RoleHandler) Get(c echox.Context) error {
	id, err := CheckRequestIDParam(c.Param("id"))
	result, err := h.service.WithContext(c).Get(id)
	if err != nil {
		return err
	}

	echox.OKWithData(c, result)
	return nil
}

// GetRoles godoc
// @Summary 	Get details of all roles
// @Description Get details of all roles
// @Tags 		roles
// @Accept  	application/json
// @Produce  	application/json
// @Param 		page_size query int true "pageSize"
// @Param 		page query int true "page"
// @Param 		filters query string false "filters separated by ",""
// @Param 		selects query string false "selects separated by ",""
// @Param 		sorts query string false "sorts separated by ",""
// @Success     200  {object} models.Response{data=model.QueryResult{records=response.Roles}} "{"data": {}}"
// @Router 		/roles/trees [get]
// @Security Bearer
func (h RoleHandler) Query(c echox.Context) error {
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
