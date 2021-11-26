package domain

type Remark struct {
	Remark string `gorm:"column:remark;size:255;"`
}
