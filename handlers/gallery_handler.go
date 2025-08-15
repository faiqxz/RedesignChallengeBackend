package handlers

import (
	"github.com/gofiber/fiber/v2"
	"redesign/database"
	"redesign/models"
	"strconv"
)

// CreateGalleryItem creates a new gallery item
func CreateGalleryItem(c *fiber.Ctx) error {
	item := new(models.Gallery)
	if err := c.BodyParser(item); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}
	result := database.DB.Create(&item)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": result.Error.Error()})
	}
	return c.JSON(item)
}

// GetAllGalleryItems retrieves all gallery items
func GetAllGalleryItems(c *fiber.Ctx) error {
	var items []models.Gallery
	result := database.DB.Find(&items)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": result.Error.Error()})
	}
	return c.JSON(items)
}

// GetGalleryItemByID retrieves a single gallery item by ID
func GetGalleryItemByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	var item models.Gallery
	result := database.DB.First(&item, id)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Gallery item not found"})
	}
	return c.JSON(item)
}

// UpdateGalleryItem partially updates a gallery item by ID (PATCH)
func UpdateGalleryItem(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	var item models.Gallery
	if err := database.DB.First(&item, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Gallery item not found"})
	}

	var updateData map[string]interface{}
	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	if imageTitle, ok := updateData["imageTitle"].(string); ok {
		item.ImageTitle = imageTitle
	}
	if description, ok := updateData["description"].(string); ok {
		item.Description = description
	}
	if imageURL, ok := updateData["imageURL"].(string); ok {
		item.ImageURL = imageURL
	}

	result := database.DB.Save(&item)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": result.Error.Error()})
	}
	return c.JSON(item)
}

// DeleteGalleryItem deletes a gallery item by ID
func DeleteGalleryItem(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	result := database.DB.Delete(&models.Gallery{}, id)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": result.Error.Error()})
	}

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Gallery item not found"})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
