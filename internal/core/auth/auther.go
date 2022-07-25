package auth

import "context"

type Auther interface {
	Token(ctx context.Context, params map[string]interface{}) (interface{}, error)
	Parse(context.Context, string) (interface{}, error)
}
