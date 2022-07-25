package echo

import (
	"net/http"
	"sort"

	"github.com/rinnguyen1614/rin-echo/internal/core/auth"
	"github.com/rinnguyen1614/rin-echo/internal/core/casbin"
	"github.com/rinnguyen1614/rin-echo/internal/core/domain"
	"github.com/rinnguyen1614/rin-echo/internal/core/echo/models"
	"github.com/rinnguyen1614/rin-echo/internal/core/utils"
	"github.com/rinnguyen1614/rin-echo/internal/core/validation"

	"github.com/labstack/echo/v4"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	core "github.com/rinnguyen1614/rin-echo/internal/core"
	"github.com/thoas/go-funk"
)

type (
	wrapError struct {
		ID         string      `json:"id,omitempty"`
		Message    string      `json:"message,omitempty"`
		InnerError interface{} `json:".inner_error,omitempty"`
		Index      int         `json:"index,omitempty"`
		Errors     []wrapError `json:"errors,omitempty"`
	}
)

func HTTPErrorHandlerWrapOnError(isWrapOnError bool) echo.HTTPErrorHandler {
	return errorHandler(isWrapOnError)
}
func HTTPErrorHandler(err error, c echo.Context) {
	errorHandler(false)(err, c)
}

func errorHandler(isWrapOnError bool) echo.HTTPErrorHandler {
	return func(err error, c echo.Context) {
		if c.Response().Committed {
			return
		}

		he, ok := err.(*echo.HTTPError)
		if ok {
			if he.Internal != nil {
				if herr, ok := he.Internal.(*echo.HTTPError); ok {
					he = herr
				}
			}
		} else {
			he = &echo.HTTPError{
				Code:    getStatusCode(err),
				Message: err.Error(),
			}
		}

		// Send response
		if c.Request().Method == http.MethodHead {
			err = c.NoContent(he.Code)
		} else {
			err = c.JSON(he.Code, models.NewResponseWithError(getWrapError(isWrapOnError, c, err)))
		}

		if err != nil {
			c.Logger().Error(err)
		}
	}
}

func getStatusCode(err error) int {
	switch err.(type) {
	case *auth.AuthenticationError:
		return http.StatusUnauthorized
	case *casbin.AuthorizationError:
		return http.StatusForbidden
	case *domain.EntityNotFoundError:
		return http.StatusNotFound
	case *validation.ValidationError:
		return http.StatusBadRequest
	case *core.RinError:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}

func getWrapError(isWrapOnError bool, c echo.Context, err error) interface{} {
	cc := NewContextx(c)
	localizer, _ := cc.Localizer()

	var wrap *wrapError

	if rinerr, ok := err.(core.Error); ok {
		wrap = &wrapError{
			InnerError: rinerr.Cause().Error(),
			ID:         rinerr.ID(),
			Message:    utils.Translate(localizer, rinerr.ID(), rinerr.Message()),
			Errors:     make([]wrapError, 0),
		}

		if vaerr, ok := err.(*validation.ValidationError); ok {
			for _, fe := range vaerr.FieldErrors() {
				wrerr := wrapError{
					InnerError: fe.Cause().Error(),
					ID:         fe.ID(),
					Message:    translateFieldError(localizer, fe),
				}

				wrap.Errors = append(wrap.Errors, wrerr)
			}
		}

		if err, ok := err.(*core.RinErrors); ok {
			var (
				errs = err.Errors()
				keys = funk.Keys(errs).([]int)
			)
			if !sort.IntsAreSorted(keys) {
				sort.Ints(keys)
			}
			for i, k := range keys {
				wrerrs := wrapError{
					Index: i,
				}
				for _, e := range errs[k] {
					wrerr := getWrapError(isWrapOnError, c, e).(*wrapError)
					wrerrs.Errors = append(wrerrs.Errors, *wrerr)
				}
				wrap.Errors = append(wrap.Errors, wrerrs)
			}
		}

	} else if he, ok := err.(*echo.HTTPError); ok {
		id := utils.ToString(he.Code)
		wrap = &wrapError{
			InnerError: err,
			ID:         id,
			Message:    utils.Translate(localizer, id, err.Error()),
		}
	} else {
		wrap = &wrapError{
			InnerError: err,
			ID:         err.Error(),
			Message:    err.Error(),
		}
	}

	if !isWrapOnError {
		wrap.InnerError = nil
	}

	return wrap
}

func translateFieldError(localizer *i18n.Localizer, fe validation.FieldError) string {
	if localizer == nil {
		return fe.Message()
	}

	return fe.TranslateWithI18n(localizer)
}
