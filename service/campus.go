package service

import (
	"fmt"
	"indonesia-University-API/models"
	"strconv"

	"gorm.io/gorm"
)

func ServiceGetCampusID(db *gorm.DB, id int) (models.Campus, error) {
	var campus models.Campus
	// query ke database atau repository
	result := db.Debug().Preload("Provinces").Where("campus_id = ?", id).First(&campus)
	if result.Error != nil {
		return models.Campus{}, result.Error
	}
	return campus, nil
}

func ServiceStringToInteger(id string) (int, error) {
	result, err := strconv.Atoi(id)
	if err != nil {
		fmt.Errorf("not correct")
		return result, err
	}
	return result, nil
}

func ServiceGetIndex(db *gorm.DB, campusName, province string) ([]models.Campus, error) {
	var campus []models.Campus
	var search *gorm.DB

	if campusName != "" {
		search = db.Preload("Provinces").Where("name LIKE ?", "%"+campusName+"%")
	}

	if province != "" {
		search = db.Preload("Provinces").Where("province LIKE ?", "%"+province+"%")
	}
	result := search.Find(&campus)

	if result.Error != nil {
		return nil, result.Error
	}

	return campus, nil
}
