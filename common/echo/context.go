package echo

import (
	"rin-echo/common"

	"github.com/labstack/echo/v4"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type (
	Context interface {
		echo.Context

		Localizer() (*i18n.Localizer, error)
		MustLocalizer() *i18n.Localizer
		SetLocalizer(*i18n.Localizer)

		// UserID() uint
		// Username() string
		SetSession(common.Session)
		MustSession() common.Session
		Session() (common.Session, error)

		Operation() (name, method string)
		SetOperation(name, method string)

		RequestContext() common.Context
	}

	contextx struct {
		echo.Context
	}

	SessionInstance func() common.Session
)

var (
	SessionKey         = "rin-echo-session"
	LocalizerKey       = "rin-echo-localizer"
	OperationNameKey   = "rin-echo-operation-name"
	OperationMethodKey = "rin-echo-operation-method"
)

func NewContextx(c echo.Context) Context {
	return &contextx{
		Context: c,
	}
}

func MustContext(c echo.Context) Context {
	cc, err := Contextx(c)
	if err != nil {
		panic(err)
	}

	return cc
}

// Cast to Contextx
func Contextx(c echo.Context) (Context, error) {
	cc, ok := c.(Context)
	if !ok {
		return nil, ERR_MISSING_CONTEXTX
	}

	return cc, nil
}

func (c contextx) Localizer() (*i18n.Localizer, error) {
	localizer := c.Get(LocalizerKey)
	if localizer == nil {
		return nil, ERR_LOCALIZER_NOT_FOUND
	}
	if _, ok := localizer.(*i18n.Localizer); !ok {
		return nil, ERR_LOCALIZER_NOT_FOUND
	}
	return localizer.(*i18n.Localizer), nil
}

func (c *contextx) MustLocalizer() *i18n.Localizer {
	localizer, err := c.Localizer()
	if err != nil {
		panic(err)
	}

	return localizer
}

func (c *contextx) SetLocalizer(localizer *i18n.Localizer) {
	c.Set(LocalizerKey, localizer)
}

func (c *contextx) SetSession(session common.Session) {
	// using to get concrete type of session.
	var f SessionInstance = func() common.Session {
		return session
	}

	c.Set(SessionKey, f)
}

func (c contextx) MustSession() common.Session {
	session, err := c.Session()
	if err != nil {
		panic(err)
	}
	return session
}

func (c contextx) Session() (common.Session, error) {
	session := c.Get(SessionKey)
	if session == nil {
		return nil, ERR_SESSION_NOT_FOUND
	}
	f := session.(SessionInstance)
	return f(), nil
}

func (c *contextx) Operation() (name, method string) {
	if v := c.Get(OperationNameKey); v != nil {
		name = v.(string)
	}
	if v := c.Get(OperationMethodKey); v != nil {
		method = v.(string)
	}

	return name, method
}

func (c *contextx) SetOperation(name, method string) {
	c.Set(OperationNameKey, name)
	c.Set(OperationMethodKey, method)
}

func (c *contextx) RequestContext() common.Context {
	cc, _ := c.Session()
	return common.NewContext(c.Request().Context(), cc)
}
