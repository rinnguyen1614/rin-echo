package uow

import (
	"errors"
	"reflect"

	"gorm.io/gorm"
)

func Find(db *gorm.DB, dest interface{}) error {
	if err := db.Find(dest).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		dest = reflect.Zero(reflect.TypeOf(dest))
	}

	return nil
}

func First(db *gorm.DB, dest interface{}) error {
	return db.First(dest).Error
}

func Count(db *gorm.DB) int64 {
	var count int64
	db.Count(&count)
	return count
}

func Contains(db *gorm.DB) bool {
	count := Count(db)
	return count > 0
}
