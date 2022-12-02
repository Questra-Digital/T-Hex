package main

import (
	"basic-app/database"
	"basic-app/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	database.Connect()
	routes.Setup(app)

	// var now time.Time
	// db.Raw("SELECT NOW()").Scan(&now)

	// fmt.Println(now)

	app.Listen(":8000")
}
