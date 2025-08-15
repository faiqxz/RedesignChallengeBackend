package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"redesign/database"
	"redesign/routes" // Import the routes package
)

func main() {
	// Connect to the database
	database.ConnectDB()

	app := fiber.New()

	// Serve static files from the "public" directory
	app.Static("/", "./public")

	// Setup routes
	routes.SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
