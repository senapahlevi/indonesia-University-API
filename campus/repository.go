package campus

import (
	"indonesia-University-API/models"

	"gorm.io/gorm"
)

type Repository interface {
	Save(campus models.Campus) (models.Campus, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(campus models.Campus) (models.Campus, error) {
	err := r.db.Create(&campus).Error
	if err != nil {
		return campus, err
	}
	return campus, nil
}
func (r *repository) FindByID(campus models.Campus) (models.Campus, error) {
	err := r.db.Where(&campus).Error
	if err != nil {
		return campus, err
	}
	return campus, nil
}
