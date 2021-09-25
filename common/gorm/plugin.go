package gorm

import (
	"reflect"
	"rin-echo/common/utils"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

const (
	callBackBeforeName       = "rin_echo:before"
	callBackBeforeCreateName = "rin_echo:before_create"
	callBackBeforeUpdateName = "rin_echo:before_update"
	callBackAfterName        = "rin_echo:after"
	appSession               = "_app_session"
)

type RinPlugin struct{}

func (p *RinPlugin) Name() string {
	return "rinPlugin"
}

func (p *RinPlugin) Initialize(db *gorm.DB) (err error) {
	_ = db.Callback().Create().Before("gorm:before_create").Register(callBackBeforeCreateName, beforeCreate)
	_ = db.Callback().Update().Before("gorm:before_update").Register(callBackBeforeUpdateName, beforeUpdate)

	return
}

func before(db *gorm.DB) {

}

func beforeCreate(db *gorm.DB) {

	if db.Error == nil {
		if field, ok := db.Statement.Schema.FieldsByName["CreatorUserID"]; ok {
			if ss, _ := AuthSession(db); ss != nil {
				setFieldValue(field, db.Statement.ReflectValue, ss.UserID())
			}
		}
		if field, ok := db.Statement.Schema.FieldsByName["UUID"]; ok {
			setFieldValue(field, db.Statement.ReflectValue, utils.MustUUID())
		}
		nowTime := time.Now()
		if field, ok := db.Statement.Schema.FieldsByName["CreatedAt"]; ok {
			setFieldValue(field, db.Statement.ReflectValue, nowTime)
		}
		if field, ok := db.Statement.Schema.FieldsByName["ModifiedAt"]; ok {
			setFieldValue(field, db.Statement.ReflectValue, nowTime)
		}
	}

	before(db)
}

func beforeUpdate(db *gorm.DB) {
	if db.Error == nil {
		if field, ok := db.Statement.Schema.FieldsByName["ModifierUserID"]; ok {
			//field.Set(db.Statement.ReflectValue, 1) //not modifyed
			if ss, _ := AuthSession(db); ss != nil {
				db.Statement.SetColumn(field.Name, ss.UserID())
			}
		}
		if field, ok := db.Statement.Schema.FieldsByName["ModifiedAt"]; ok {
			db.Statement.SetColumn(field.Name, time.Now())
		}
	}

	before(db)
}

func setFieldValue(field *schema.Field, reflectValue reflect.Value, value interface{}) {
	switch reflectValue.Kind() {
	case reflect.Slice, reflect.Array:
		if _, ok := value.(utils.UUID); ok {
			for i := 0; i < reflectValue.Len(); i++ {
				field.Set(reflectValue.Index(i), utils.MustUUID())
			}
		} else {
			for i := 0; i < reflectValue.Len(); i++ {
				field.Set(reflectValue.Index(i), value)
			}
		}
	case reflect.Struct:
		field.Set(reflectValue, value)
	}
}
