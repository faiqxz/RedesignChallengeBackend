package handlers

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

// UploadFile handles image upload
func UploadFile(c *fiber.Ctx) error {
	// Parse the multipart form:
	file, err := c.FormFile("image")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse image.",
		})
	}

	// Generate a unique filename
	uniqueFileName := fmt.Sprintf("%d-%s", time.Now().UnixNano(), file.Filename)
	dst := fmt.Sprintf("./public/uploads/%s", uniqueFileName)

	// Save the file to the destination
	if err := c.SaveFile(file, dst); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot save file.",
		})
	}

	// Create the public URL for the file
	publicURL := fmt.Sprintf("/uploads/%s", uniqueFileName)

	return c.JSON(fiber.Map{
		"url": publicURL,
	})
}
