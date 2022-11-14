package controller

import (
	"fmt"

	"github.com/Saatvik-droid/url-shortner/model"
	"github.com/gofiber/fiber/v2"
)

func CreateUrl(c *fiber.Ctx) error {
	var url model.Url

	err := c.BodyParser(&url)
	if err != nil {
		fmt.Println(err)
	}

	tx := model.DB.Create(&url)
	if tx.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": tx.Error,
		})
	}

	return c.Status(fiber.StatusOK).JSON(url)
}

func Redirect(c *fiber.Ctx) error {
	var Url model.Url
	url := c.Params("redirect")

	tx := model.DB.Where("url = ?", url).First(&Url)
	fmt.Println(Url)
	if tx.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": tx.Error,
		})
	}

	Url.Clicked += 1
	tx = model.DB.Model(&Url).Update("Clicked", Url.Clicked)
	if tx.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": tx.Error,
		})
	}
	return c.Redirect(Url.Redirect, fiber.StatusTemporaryRedirect)
}
