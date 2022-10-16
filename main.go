package main

import (
	"fmt"

	"github.com/dchung117/crm_golang_fiber/database"
	"github.com/dchung117/crm_golang_fiber/lead"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// routes
func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/lead", lead.GetLeads)          // get all
	app.Get("/api/v1/lead/:id", lead.GetLead)       // get one
	app.Post("/api/v1/lead", lead.NewLead)          // post one
	app.Delete("/api/v1/lead/:id", lead.DeleteLead) // delete one
}

// initialize database
func initDatabase() {
	// define error variable
	var err error

	// launch database
	database.DBConn, err = gorm.Open("sqlite3", "leads.db")
	if err != nil {
		panic("Failed to connect to database.")
	}
	fmt.Println("Connection to database successful.")

	// migrate models for database
	database.DBConn.AutoMigrate(&lead.Lead{})
	fmt.Println("Database succcessfully migrated.")
}

func main() {
	// define an app
	app := fiber.New()

	// define database
	initDatabase()

	// set up routes for the app
	setupRoutes(app)

	// launch server, begin listening
	app.Listen(3000)

	// defer closing the database connection after main terminates
	defer database.DBConn.Close()
}
