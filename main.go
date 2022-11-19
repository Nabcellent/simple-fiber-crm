package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"simple-fiber-crm/database"
	"simple-fiber-crm/lead"
)

func setUpRoutes(app *fiber.App) {
	app.Get("/api/v1/leads", lead.GetLeads)
	app.Get("/api/v1/leads/:id", lead.GetLead)
	app.Post("/api/v1/leads", lead.CreateLead)
	app.Delete("/api/v1/lead/:id", lead.DeleteLead)
}

func initDB() {
	var err error

	database.DB, err = gorm.Open(sqlite.Open("test.db"))
	if err != nil {
		panic("Failed to connect to DB!")
	}

	fmt.Println("Connected to DB!")

	err = database.DB.AutoMigrate(&lead.Lead{})
	if err != nil {
		panic("Failed to auto-migrate!")
	}

	fmt.Println("Database auto-migrated!")
}

func main() {
	app := fiber.New()

	initDB()

	setUpRoutes(app)

	app.listen(8000)

	defer database.DB.Close()
}
