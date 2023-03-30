package routes

import (
	"github.com/gofiber/fiber/v2"
	blg "github.com/lemmaritzy/blogger/api/controllers/blog"
)

func SetupRouter(app *fiber.App) {

	b := app.Group("/api/blogger")
	b.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"msg": "api router",
		})
	})

	blogs := b.Group("/blogs")
	blogs.Get("/all-blogs", blg.LoadBlogs)
	blogs.Get("/all-blogs/:id?", blg.FindBlog)
	blogs.Post("/create-blog", blg.CreateBlog)
	blogs.Put("/update-blog/:id?", blg.UpdateBlog)

	profile := b.Group("/profile")
	profile.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("profile main")
	})

}
