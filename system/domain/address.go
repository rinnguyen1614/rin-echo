package domain

import (
	"github.com/rinnguyen1614/rin-echo-core/domain"
	gormx "github.com/rinnguyen1614/rin-echo-core/gorm/data_types"
)

type Address struct {
	domain.FullAuditedEntity

	UserID    uint   `gorm:"column:user_id;index:idx_addresses_user_id_is_primary,unique,priority:1"`
	IsPrimary bool   `gorm:"column:user_id;index:idx_addresses_user_id_is_primary,unique,priority:2"`
	Name      string `gorm:"column:name;size:255;default:'';not null;"`
	Phone     string `gorm:"column:name;size:16;default:'';not null;"`
	// <FullAddress>
	City         string `gorm:"column:name;size:255;default:'';not null;"`
	CityID       *uint
	District     string `gorm:"column:name;size:255;default:'';not null;"`
	DistrictID   *uint
	State        string `gorm:"column:name;size:255;default:'';not null;"`
	StateID      *uint
	Country      string `gorm:"column:name;size:255;default:'';not null;"`
	CountryID    *uint
	Zipcode      string `gorm:"column:name;size:64;"`
	AddressLine1 string `gorm:"column:address_line_1;size:255;default:'';not null;"`
	AddressLine2 string `gorm:"column:address_line_2;size:255"`

	LocationID *uint
	// </FullAdress>
	Location AddressLocation `gorm:"foreignKey:location_id;references:id"`
}

type Addresses []*Address

type AddressLocation struct {
	domain.Entity

	Confirm  bool `gorm:"column:confirm;default:false;"`
	Location *gormx.Point
}
