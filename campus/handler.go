package campus

import (
	"indonesia-University-API/databases"
	"indonesia-University-API/service"
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

// func GetCampusID(c *gin.Context) {
// 	var campus models.Campus
// 	// var response ResponseCampus
// 	id := c.Param("id")
// 	// check id is valid ?
// 	if _, err := strconv.Atoi(id); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "id harus berupa angkas"})
// 		return
// 	}
// 	result := db.Debug().Preload("Provinces").Where("campus_id = ?", id).First(&campus)
// 	if result.Error != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{"data": campus})
// }

func GetCampusID(c *gin.Context) {
	// var campus models.Campus
	// var response ResponseCampus
	id := c.Param("id")
	// check id is valid ?
	// campusID, err := strconv.Atoi(id)
	campusID, err := service.ServiceStringToInteger(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID harus berupa angka"})
		return
	}
	// campus, err := ServiceGetCampusID(campusID)
	campus, err := service.ServiceGetCampusID(db, campusID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": campus})
}

// indexing
func GetIndexingCampus(c *gin.Context) {

	campusName := c.Query("CampusName")

	province := c.Query("province")

	campus, err := service.ServiceGetIndex(db, campusName, province)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if len(campus) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Data tidak ditemukan"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": campus})
}
