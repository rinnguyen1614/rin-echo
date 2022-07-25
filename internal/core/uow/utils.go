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

func Count(db *gorm.DB) (int64, error) {
	var count int64
	err := db.Count(&count).Error
	return count, err
}

func Contains(db *gorm.DB) (bool, error) {
	count, err := Count(db)
	return count > 0, err
}
