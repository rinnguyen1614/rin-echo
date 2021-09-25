package common

import (
	"context"
)

type (
	Context struct {
		context.Context
		Session Session
	}
)
