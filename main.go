package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	r "github.com/lemmaritzy/blogger/api/routes"
)

func main() {

	app := fiber.New()
	app.Use(cors.New())

	r.SetupRouter(app)

	app.Listen(":3030")
}
