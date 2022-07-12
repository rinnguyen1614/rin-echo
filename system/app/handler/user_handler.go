package handler

import (
	echox "github.com/rinnguyen1614/rin-echo-core/echo"
	"github.com/rinnguyen1614/rin-echo-core/model"
	"github.com/rinnguyen1614/rin-echo-core/setting"
	iuow "github.com/rinnguyen1614/rin-echo-core/uow/interfaces"
	"github.com/rinnguyen1614/rin-echo-core/validation"

	"rin-echo/system/app/model/request"
	"rin-echo/system/app/service"
	"rin-echo/system/domain"

	"go.uber.org/zap"

	rquery "github.com/rinnguyen1614/rin-echo-core/echo/models/query/rest_query"
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

// CreateUser godoc
// @Summary      Create a new user
// @Description  Create a new user with the input payload
// @Tags         users
// @Accept       application/json
// @Produce      application/json
// @Param 		 data body request.CreateUser true "Create user"
// @Success      200  {object}  models.Response{data=model.Model} "{"data": {}}"
// @Router       /users [post]
// @Security Bearer
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

// UpdateUser godoc
// @Summary      Update user identified by the given id
// @Description  Update the user corresponding to the input id
// @Tags         users
// @Accept       application/json
// @Produce      application/json
// @Param 		 id path int true "ID of the user to be updated"
// @Param 		 data body request.UpdateUser true "Update user"
// @Success      200  {object}  models.Response{data=model.Model} "{"data": {}}"
// @Router       /users/{id} [put]
// @Security Bearer
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

// DeleteUser godoc
// @Summary      Delete user identified by the given id
// @Description  Delete the user corresponding to the input id
// @Tags         users
// @Accept       application/json
// @Produce      application/json
// @Param 		 id path int true "ID of the user to be deleted"
// @Success      200  {object}  models.Response "{"data": {}}"
// @Router       /users/{id} [delete]
// @Security Bearer
func (h UserHandler) Delete(c echox.Context) error {
	id, err := CheckRequestIDParam(c.Param("id"))
	err = h.service.WithContext(c).Delete(id)
	if err = h.service.WithContext(c).Delete(id); err != nil {
		return err
	}

	echox.OKWithData(c, nil)
	return nil
}

// GetUser godoc
// @Summary 	Get details for a given id
// @Description Get details of user corresponding to the input id
// @Tags 		users
// @Accept  	application/json
// @Produce  	application/json
// @Param 		id path int true "ID of the user"
// @Success     200  {object} models.Response{data=response.User} "{"data": {}}"
// @Router 		/users/{id} [get]
// @Security Bearer
func (h UserHandler) Get(c echox.Context) error {
	id, err := CheckRequestIDParam(c.Param("id"))
	result, err := h.service.WithContext(c).Get(id)
	if err != nil {
		return err
	}

	echox.OKWithData(c, result)
	return nil
}

// GetUsers godoc
// @Summary 	Get details of all users
// @Description Get details of all users
// @Tags 		users
// @Accept  	application/json
// @Produce  	application/json
// @Param 		page_size query int true "pageSize"
// @Param 		page query int true "page"
// @Param 		filters query string false "filters separated by ",""
// @Param 		selects query string false "selects separated by ",""
// @Param 		sorts query string false "sorts separated by ",""
// @Success     200  {object} models.Response{data=model.QueryResult{records=response.Users}} "{"data": {}}"
// @Router 		/users/trees [get]
// @Security Bearer
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
