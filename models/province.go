package models

import "time"

type Province struct {
	ProvinceID int       `gorm:"primaryKey;province_id" json:"province_id"`
	Name       string    `gorm:"name" json:"name"`
	CreatedAt  time.Time `gorm:"created_at" json:"-"`
}

func (Province) TableName() string {
	return "province"
}
