package blog

import (
	"github.com/gofiber/fiber/v2"
	b_model "github.com/lemmaritzy/blogger/api/models/blog"
	"github.com/lemmaritzy/blogger/internal/entities"
	"strconv"
)

var (
	Data       interface{}
	Record     int64
	Error      error
	StatusCode int
	Message    string
)

func CreateBlog(c *fiber.Ctx) error {
	blog := entities.Blog{}
	errParser := c.BodyParser(&blog)

	if errParser != nil {
		return errParser
	}

	Record, Data = b_model.CreateBlogModel("12", blog)

	if Data == 0 {
		return c.JSON(fiber.Map{
			"data":   Data,
			"status": fiber.StatusInternalServerError,
			"msg":    "an error occurred while inserting data",
		})
	} else {
		return c.JSON(fiber.Map{
			"data":   Data,
			"status": fiber.StatusOK,
			"msg":    "blog inserted successfully",
		})
	}
}

func LoadBlogs(c *fiber.Ctx) error {
	Record, Data = b_model.LoadBlogs()

	if Record == -1 {
		StatusCode = fiber.StatusInternalServerError
		Message = "Sistem hatası"
	} else if Record == 0 {
		StatusCode = fiber.StatusOK
		Message = "Kayıtlar mevcut"
	} else {
		StatusCode = fiber.StatusOK
		Message = "Kayıt yok"
	}
	return c.JSON(fiber.Map{
		"data":    Data,
		"message": Message,
		"status":  StatusCode,
	})
}

func FindBlog(c *fiber.Ctx) error {
	blogid, _ := strconv.ParseInt(c.Params("id"), 10, 64)
	Record, Data = b_model.LoadBlog(blogid)
	if Record == -1 {
		StatusCode = fiber.StatusInternalServerError
		Message = "Sistem hatası"
	} else if Record == 0 {
		StatusCode = fiber.StatusOK
		Message = "Kayıt mevcut"
	} else {
		StatusCode = fiber.StatusOK
		Message = "Kayıt yok"
	}

	return c.JSON(fiber.Map{
		"data": Data,
	})
}

func UpdateBlog(c *fiber.Ctx) error {
	blog := entities.Blog{}

	blogid, _ := strconv.ParseInt(c.Params("id"), 10, 64)
	errParser := c.BodyParser(&blog)
	if errParser != nil {
		return errParser
	}
	Record, Data = b_model.UpdateBlog(blogid, blog)

	return c.JSON(fiber.Map{
		"record": Record,
		"data":   Data,
	})
}
