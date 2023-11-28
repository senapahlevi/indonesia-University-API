package models

type Campus struct {
	CampusID   int        `gorm:"primaryKey;campus_id" json:"campus_id"`
	Name       string     `gorm:"name" json:"name"`
	Picture    string     `gorm:"picture" json:"picture"`
	ProvinceID int        `gorm:"province_id" json:"province_id"`
	Provinces  []Province `gorm:"foreignKey:ProvinceID" json:"provinces"`
}

func (Campus) TableName() string {
	return "campus"
}
