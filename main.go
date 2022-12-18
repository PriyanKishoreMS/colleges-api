package main

import (
	"log"

	"github.com/PriyanKishoreMS/colleges-list-api/config"
	"github.com/PriyanKishoreMS/colleges-list-api/handlers"
	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	config.Connect()

	app.Static("/", "./public")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendFile("./public/index.html")
	})

	app.Get("colleges/", handlers.SearchCollege)
	app.Get("colleges/states", handlers.GetAllStates)
	app.Get("colleges/:state/districts", handlers.GetDistrictsByState)
	app.Get("colleges/:state", handlers.GetAllCollegesInState)
	app.Get("colleges/:state/:district", handlers.GetAllCollegesInDistrict)

	log.Fatal(app.Listen(":3000"))
}
