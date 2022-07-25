package middleware

import (
	"github.com/rinnguyen1614/rin-echo/internal/core/auth"
	jwtx "github.com/rinnguyen1614/rin-echo/internal/core/auth/jwt"
	echox "github.com/rinnguyen1614/rin-echo/internal/core/echo"

	"github.com/golang-jwt/jwt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type (
	JWTConfig struct {
		Header  string
		Skipper middleware.Skipper
		Schema  string
		// you can set claims from token
		WrapSessionContext func(c echo.Context, claims jwt.Claims)
		Auther             auth.Auther
	}

	jwtExtractor func(echo.Context) (string, error)
)

var (
	// DefaultJWTConfig is the default JWT middleware config.
	defaultAutherConfig = jwtx.DefaultJWTConfig
	DefaultJWTConfig    = JWTConfig{
		Schema:             "Bearer",
		Header:             echo.HeaderAuthorization,
		Skipper:            middleware.DefaultSkipper,
		WrapSessionContext: jwtDefaultWrapSessionContext,
	}
)

func JWT(auther auth.Auther) echo.MiddlewareFunc {
	c := DefaultJWTConfig
	c.Auther = auther
	return JWTWithConfig(c)
}

func JWTWithConfig(config JWTConfig) echo.MiddlewareFunc {
	if config.Auther == nil {
		panic("jwt requires Auther")
	}
	if config.Skipper == nil {
		config.Skipper = DefaultJWTConfig.Skipper
	}
	if config.Header == "" {
		config.Header = DefaultJWTConfig.Header
	}
	if config.Schema == "" {
		config.Schema = DefaultJWTConfig.Schema
	}
	if config.WrapSessionContext == nil {
		config.WrapSessionContext = DefaultJWTConfig.WrapSessionContext
	}

	extractor := jwtFromHeader(config.Header, config.Schema)

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if config.Skipper != nil && config.Skipper(c) {
				return next(c)
			}

			token, err := extractor(c)
			if err != nil {
				return err
			}

			tokenParsed, err := config.Auther.Parse(c.Request().Context(), token)
			if err != nil {
				return err
			}

			config.WrapSessionContext(c, tokenParsed.(*jwt.Token).Claims)

			return next(c)
		}
	}
}

func jwtFromHeader(header string, authScheme string) jwtExtractor {
	return func(c echo.Context) (string, error) {
		authHeader := c.Request().Header.Get(header)
		l := len(authScheme)
		if len(authHeader) > l+1 && authHeader[:l] == authScheme {
			return authHeader[l+1:], nil
		}
		return "", auth.ERR_TOKEN_MISSING
	}
}

func jwtDefaultWrapSessionContext(c echo.Context, claims jwt.Claims) {
	cc := echox.MustContext(c)
	cc.SetSession(claims.(*jwtx.ClaimsSession))
}
