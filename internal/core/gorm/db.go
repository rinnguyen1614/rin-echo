package gorm

import (
	"errors"
	"reflect"

	core "github.com/rinnguyen1614/rin-echo/internal/core"
	"gorm.io/gorm"
)

type (
	funcSession func() core.Session
)

func AuthSession(db *gorm.DB) (core.Session, error) {
	sctx := db.Statement.Context
	ctx, ok := sctx.(core.Context)
	if !ok {
		return nil, core.NewRinError("not_found_current_session", "Not found current session")
	}
	session := ctx.MustSession()
	if session == nil {
		return nil, core.NewRinError("not_found_current_session", "Not found current session")
	}
	return session, nil
}

func FindWrapError(db *gorm.DB, dest interface{}) error {
	if err := db.Find(dest).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		dest = reflect.Zero(reflect.TypeOf(dest))
	}

	return nil
}
