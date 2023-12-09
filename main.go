package main

import (
	"indonesia-University-API/campus"
	"indonesia-University-API/databases"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	db, err := databases.SetDB()
	if err != nil {
		log.Fatal(err)
	}
	campus.SetDatabase(db)
	router := gin.Default()
	api := router.Group("/api")
	api.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"*", "http://localhost:3000"},
		ExposeHeaders:    []string{"Content-Type"},
		AllowCredentials: true,
	}))
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000", "*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"Content-Type"},
		// ExposeHeaders:    []string{"Content-Type"},
		AllowCredentials: true,
	}))
	api.GET("/campus/:id", campus.GetCampusID)
	api.GET("/campus", campus.GetIndexingCampus)

	// router.Run(":" + os.Getenv("PORT"))
	router.Run(":8080")
}

func init() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
