package inject

import "github.com/rinnguyen1614/rin-echo/internal/core/auth/jwt"

type Claims struct {
	jwt.ClaimsSession

	Email string `json:"email"`
}

func GetAuther() *jwt.JWT {
	if service.auther == nil {
		cfg := GetConfig()
		service.auther = jwt.NewJWTWithConfig(jwt.JWTConfig{
			SigningKey:  []byte(cfg.JWT.SigningKey),
			ExpiresTime: cfg.JWT.ExpiresTime,
			Issuer:      cfg.JWT.Issuer,
			Claims:      &Claims{},
		})
	}
	return service.auther
}
