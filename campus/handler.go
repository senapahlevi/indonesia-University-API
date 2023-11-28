package campus

import (
	"indonesia-University-API/databases"
	"indonesia-University-API/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB

func SetDatabase(databases *databases.DB) {
	db = databases.DB
}

// func GetCampus(c *gin.Context) {
// 	var campus models.Campus
// 	id := c.Param("id")
// 	result := db.Preload("De")
// }

func GetCampusID(c *gin.Context) {
	var campus models.Campus
	id := c.Param("id")

	result := db.Preload("Provinces").Where("campus_id = ?", id).First(&campus)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"data": campus})
}

// indexing
func GetIndexingCampus(c *gin.Context) {
	var campus []models.Campus
	var search *gorm.DB

	campusName := c.Query("CampusName")
	if campusName != "" {
		search = db.Where("name LIKE ?", "%"+campusName+"%")
	}

	province := c.Query("province")
	if province != "" {
		search = db.Where("province LIKE ?", "%"+province+"%")
	}

	result := search.Find(&campus)
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Data tidak ditemukan"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": campus})
}
