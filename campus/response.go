package campus

import "indonesia-University-API/models"

type ResponseCampus struct {
	CampusID  int               `json:"campus_id"`
	Name      string            `json:"name"`
	Picture   string            `json:"picture"`
	Provinces []models.Province `json:"provinces"`
}
