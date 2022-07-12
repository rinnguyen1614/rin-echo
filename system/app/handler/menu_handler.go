package handler

import (
	echox "github.com/rinnguyen1614/rin-echo-core/echo"
	rquery "github.com/rinnguyen1614/rin-echo-core/echo/models/query/rest_query"
	"github.com/rinnguyen1614/rin-echo-core/model"
	"github.com/rinnguyen1614/rin-echo-core/setting"
	iuow "github.com/rinnguyen1614/rin-echo-core/uow/interfaces"
	"github.com/rinnguyen1614/rin-echo-core/validation"

	"rin-echo/system/app/model/request"
	_ "rin-echo/system/app/model/response"
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

// CreateMenu godoc
// @Summary      Create a new menu
// @Description  Create a new menu with the input payload
// @Tags         menus
// @Accept       application/json
// @Produce      application/json
// @Param 		 data body request.CreateMenu true "Create menu"
// @Success      200  {object}  models.Response{data=model.Model} "{"data": {}}"
// @Router       /menus [post]
// @Security Bearer
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

// UpdateMenu godoc
// @Summary      Update menu identified by the given id
// @Description  Update the menu corresponding to the input id
// @Tags         menus
// @Accept       application/json
// @Produce      application/json
// @Param 		 id path int true "ID of the menu to be updated"
// @Param 		 data body request.UpdateMenu true "Update menu"
// @Success      200  {object}  models.Response{data=model.Model} "{"data": {}}"
// @Router       /menus/{id} [put]
// @Security Bearer
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

// DeleteMenu godoc
// @Summary      Delete menu identified by the given id
// @Description  Delete the menu corresponding to the input id
// @Tags         menus
// @Accept       application/json
// @Produce      application/json
// @Param 		 id path int true "ID of the menu to be deleted"
// @Success      200  {object}  models.Response "{"data": {}}"
// @Router       /menus/{id} [delete]
// @Security Bearer
func (h MenuHandler) Delete(c echox.Context) error {
	id, err := CheckRequestIDParam(c.Param("id"))
	err = h.service.WithContext(c).Delete(id)
	if err = h.service.WithContext(c).Delete(id); err != nil {
		return err
	}

	echox.OKWithData(c, nil)
	return nil
}

// GetMenu godoc
// @Summary 	Get details for a given id
// @Description Get details of menu corresponding to the input id
// @Tags 		menus
// @Accept  	application/json
// @Produce  	application/json
// @Param 		id path int true "ID of the menu"
// @Success     200  {object} models.Response{data=response.Menu} "{"data": {}}"
// @Router 		/menus/{id} [get]
// @Security Bearer
func (h MenuHandler) Get(c echox.Context) error {
	id, err := CheckRequestIDParam(c.Param("id"))
	result, err := h.service.WithContext(c).Get(id)
	if err != nil {
		return err
	}

	echox.OKWithData(c, result)
	return nil
}

// GetMenus godoc
// @Summary 	Get details of all menus
// @Description Get details of all menus
// @Tags 		menus
// @Accept  	application/json
// @Produce  	application/json
// @Param 		page_size query int true "pageSize"
// @Param 		page query int true "page"
// @Param 		filters query string false "filters separated by ",""
// @Param 		selects query string false "selects separated by ",""
// @Param 		sorts query string false "sorts separated by ",""
// @Success     200  {object} models.Response{data=model.QueryResult{records=response.MenuTrees}} "{"data": {}}"
// @Router 		/menus/trees [get]
// @Security Bearer
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
