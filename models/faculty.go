package models

type Faculty struct {
	ID         int      `gorm:"primaryKey;id" json:"id"`
	Name       string   `gorm:"name" json:"name"`
	Picture    string   `gorm:"picture" json:"picture"`
	ProvinceID int      `gorm:"province_id" json:"province_id"`
	CampusID   []Campus `gorm:"foreignKey:InvoiceID" json:"detail_items"`
}

func (Faculty) TableName() string {
	return "faculty"
}
