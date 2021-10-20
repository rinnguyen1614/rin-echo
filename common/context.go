package common

import (
	"context"
	"time"
)

type (
	Context struct {
		context.Context
		Session Session
	}
)

func (c Context) WithDeadline(d time.Time) (Context, context.CancelFunc) {
	ctx, f := context.WithDeadline(c.Context, d)
	return Context{
		Context: ctx,
		Session: c.Session,
	}, f
}

func (c Context) WithTimeout(timeout time.Duration) (Context, context.CancelFunc) {
	ctx, f := context.WithTimeout(c.Context, timeout)
	return Context{
		Context: ctx,
		Session: c.Session,
	}, f
}

func (c Context) WithValue(key, value interface{}) Context {
	return Context{
		Context: context.WithValue(c.Context, key, value),
		Session: c.Session,
	}
}
