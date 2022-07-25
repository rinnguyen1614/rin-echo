package core

import (
	ctx "context"

	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type (
	Context interface {
		ctx.Context
		SetSession(Session)
		MustSession() Session
		Session() (Session, error)

		Localizer() (*i18n.Localizer, error)
		MustLocalizer() *i18n.Localizer
		SetLocalizer(*i18n.Localizer)
	}
	context struct {
		ctx.Context
		session   Session
		localizer *i18n.Localizer
	}
)

func NewContext(ctx ctx.Context, session Session) Context {
	return &context{
		Context: ctx,
		session: session,
	}
}

func (c *context) WithContext(ctx ctx.Context) Context {
	return &context{
		Context: ctx,
		session: c.session,
	}
}

func (c context) Session() (Session, error) {
	return c.session, nil
}

func (c context) MustSession() Session {
	return c.session
}

func (c *context) SetSession(session Session) {
	c.session = session
}

func (c context) Localizer() (*i18n.Localizer, error) {
	return c.localizer, nil
}

func (c context) MustLocalizer() *i18n.Localizer {
	return c.localizer
}

func (c *context) SetLocalizer(localizer *i18n.Localizer) {
	c.localizer = localizer
}
