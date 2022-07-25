package handler

import (
	"github.com/rinnguyen1614/rin-echo/internal/system/app/service"

	echox "github.com/rinnguyen1614/rin-echo/internal/core/echo"
	rquery "github.com/rinnguyen1614/rin-echo/internal/core/echo/models/query/rest_query"
	"github.com/rinnguyen1614/rin-echo/internal/core/setting"
	iuow "github.com/rinnguyen1614/rin-echo/internal/core/uow/interfaces"
	"github.com/rinnguyen1614/rin-echo/internal/core/validation"

	"go.uber.org/zap"
)

type AuditLogHandler struct {
	echox.Handler
	service service.AuditLogService
}

func NewAuditLogHandler(uow iuow.UnitOfWork,
	logger *zap.Logger,
	restQuery rquery.RestQuery,
	settingProvider setting.Provider,
	validator *validation.Validator) AuditLogHandler {
	return AuditLogHandler{
		Handler: echox.NewHandler(logger, restQuery, settingProvider, validator),
		service: service.NewAuditLogService(uow, settingProvider, logger),
	}
}

// GetAuditLog godoc
// @Summary 	Get details for a given id
// @Description Get details of resource corresponding to the input id
// @Tags 		resources
// @Accept  	application/json
// @Produce  	application/json
// @Param 		id path int true "ID of the resource"
// @Success     200  {object} models.Response{data=response.AuditLog} "{"data": {}}"
// @Router 		/resources/{id} [get]
// @Security Bearer
func (h AuditLogHandler) Get(c echox.Context) error {
	id, err := CheckRequestIDParam(c.Param("id"))
	result, err := h.service.WithContext(c).Get(id)
	if err != nil {
		return err
	}

	echox.OKWithData(c, result)
	return nil
}

// GetAuditLogs godoc
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
// @Success     200  {object} models.Response{data=model.QueryResult{records=response.AuditLogs}} "{"data": {}}"
// @Router 		/users/trees [get]
// @Security Bearer
func (h AuditLogHandler) Query(c echox.Context) error {
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
