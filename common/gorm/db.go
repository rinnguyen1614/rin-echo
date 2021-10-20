package gorm

import (
	"errors"
	"reflect"
	"rin-echo/common"

	"gorm.io/gorm"
)

type (
	funcSession func() common.Session
)

func AuthSession(db *gorm.DB) (common.Session, error) {
	sctx := db.Statement.Context
	ctx, ok := sctx.(common.Context)
	if !ok || ctx.Session == nil {
		return nil, common.NewRinError("not_found_current_session", "Not found current session")
	}

	return ctx.Session, nil
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
