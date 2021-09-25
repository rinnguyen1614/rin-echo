package jwt

import (
	"context"
	"fmt"
	"rin-echo/common/auth"
	"testing"

	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
)

func TestAuth_JWT_TokenWithMapClaims(t *testing.T) {

	j := NewJWT("secret")
	mapClaims := jwt.MapClaims{
		"Username": "anhnguyen",
		"FullName": "Anh Nguyen",
	}

	token, err := j.Token(context.Background(), mapClaims)

	if err != nil {
		t.Errorf("err: %v\n", err)
		return
	}

	claims, err := j.Parse(context.Background(), token.(*Token).AccessToken)
	if err != nil {
		t.Errorf("err: %v\n", err)
		return
	}

	if mapClaimsResult, ok := claims.(*jwt.Token).Claims.(jwt.MapClaims); ok {
		assert.Equal(t, mapClaims, mapClaimsResult)
	} else {
		t.Errorf("err: claims isn't a mapClaims\n")
		return
	}

}

func TestAuth_JWT_TokenWithCustomClaims(t *testing.T) {

	type jwtCustomClaims struct {
		jwt.StandardClaims
		auth.ClaimsSession
	}

	j := NewJWTWithConfig(JWTConfig{
		Claims:     &jwtCustomClaims{},
		SigningKey: []byte("secret"),
	})

	token, err := j.Token(context.Background(), map[string]interface{}{
		"Username": "anhnguyen",
		"FullName": "Anh Nguyen"})

	if err != nil {
		t.Errorf("err: %v\n", err)
		return
	}

	claimsParsed, err := j.Parse(context.Background(), token.(*Token).AccessToken)
	if err != nil {
		t.Errorf("err: %v\n", err)
		return
	}

	if claimsResult, ok := claimsParsed.(*jwt.Token).Claims.(*jwtCustomClaims); ok {
		assert.Equal(t,
			"anhnguyen",
			claimsResult.Username)

		fmt.Printf("%+v\n", claimsResult)
	} else {
		t.Errorf("err: claims isn't a jwtCustomClaims\n")
		return
	}

}
