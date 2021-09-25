package jwt

import (
	"context"
	"reflect"
	"rin-echo/common/auth"
	"rin-echo/common/utils"
	"time"

	"github.com/golang-jwt/jwt"
)

type (
	JWTConfig struct {
		SigningKey  []byte
		ExpiresTime int64
		Issuer      string
		Claims      jwt.Claims
		// It is a param's req for VerifyIssuer(issuer, req).
		// If required is false, this method will return true if the value matches or is unset
		RequiredVerifyIssuer bool

		jwt.Keyfunc
	}

	JWT struct {
		JWTConfig
	}

	Token struct {
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"`
		Expiry      int64  `json:"expiry"`
	}

	ClaimsSession struct {
		auth.ClaimsSession
		jwt.StandardClaims
	}
)

var DefaultJWTConfig = JWTConfig{
	ExpiresTime:          7200, // 2hours
	Issuer:               "rin-echo",
	Claims:               &ClaimsSession{},
	RequiredVerifyIssuer: true,
	Keyfunc:              nil,
}

func NewJWT(key string) *JWT {
	c := DefaultJWTConfig
	c.SigningKey = []byte(key)
	return NewJWTWithConfig(c)
}

func NewJWTWithConfig(config JWTConfig) *JWT {

	if config.SigningKey == nil && config.Keyfunc == nil {
		panic("jwt requires signing key")
	}

	if config.Issuer == "" {
		config.Issuer = DefaultJWTConfig.Issuer
	}

	if config.Claims == nil {
		config.Claims = DefaultJWTConfig.Claims
	}

	if config.Keyfunc == nil {
		config.Keyfunc = config.defaultKeyfunc
	}

	return &JWT{config}
}

func (j *JWT) Token(ctx context.Context, params map[string]interface{}) (interface{}, error) {
	now := time.Now()
	expiresAt := now.Add(time.Duration(j.ExpiresTime) * time.Second).Unix()

	var claims jwt.Claims
	if _, ok := j.Claims.(jwt.MapClaims); ok {
		claims = jwt.MapClaims(params)
	} else {
		// Initial StandardClaims first, avoid override field of params
		src := map[string]interface{}{
			"StandardClaims": jwt.StandardClaims{
				Issuer:    j.Issuer,
				IssuedAt:  now.Unix(),
				ExpiresAt: expiresAt,
			},
		}
		// clone params to src
		for k, v := range params {
			src[k] = v
		}

		c := reflect.New(reflect.ValueOf(j.Claims).Type().Elem()).Interface()
		if err := utils.ConvertMapToStruct(c, src); err != nil {
			return nil, err
		}
		claims = c.(jwt.Claims)
	}

	// Create token with claims
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Sign and get the complete encoded token as a string
	token, err := tokenClaims.SignedString(j.SigningKey)
	if err != nil {
		return nil, err
	}

	return &Token{
		AccessToken: token,
		TokenType:   "Bearer",
		Expiry:      expiresAt,
	}, nil
}

func (j *JWT) Parse(ctx context.Context, tokenString string) (interface{}, error) {
	var token *jwt.Token
	var err error
	if _, ok := j.Claims.(jwt.MapClaims); ok {
		token, err = jwt.Parse(tokenString, j.Keyfunc)
	} else {
		t := reflect.ValueOf(j.Claims).Type().Elem()
		claims := reflect.New(t).Interface().(jwt.Claims)
		token, err = jwt.ParseWithClaims(tokenString, claims, j.Keyfunc)
	}

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, auth.ERR_TOKEN_MALFORMED
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, auth.ERR_TOKEN_EXPIRED
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, auth.ERR_TOKEN_INVALID
			} else {
				return nil, auth.ERR_TOKEN_INVALID
			}
		}
		return nil, auth.ERR_TOKEN_INVALID
	}

	if !token.Valid {
		return nil, auth.ERR_TOKEN_INVALID
	}

	if stdClaims, ok := token.Claims.(*jwt.StandardClaims); ok && !stdClaims.VerifyIssuer(j.Issuer, j.RequiredVerifyIssuer) {
		return nil, auth.ERR_TOKEN_ISSUER
	}

	return token, nil
}

func (config *JWTConfig) defaultKeyfunc(t *jwt.Token) (interface{}, error) {
	if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, auth.ERR_TOKEN_INVALID
	}
	return []byte(config.SigningKey), nil
}
