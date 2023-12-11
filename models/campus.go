package models

type Campus struct {
	CampusID   int        `gorm:"primaryKey;campus_id" json:"campus_id"`
	Name       string     `gorm:"name" json:"name"`
	Picture    string     `gorm:"picture" json:"-"`
	ProvinceID int        `gorm:"province_id" json:"-"`
	Address    string     `gorm:"address" json:"address"`
	Provinces  []Province `gorm:"foreignKey:ProvinceID;references:province_id" json:"provinces"`
}

func (Campus) TableName() string {
	return "campus"
}
