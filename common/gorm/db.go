package gorm

import (
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
