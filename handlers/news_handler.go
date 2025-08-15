package handlers

import (
	"github.com/gofiber/fiber/v2"
	"redesign/database"
	"redesign/models"
	"strconv"
)

// CreateNews creates a new news item
func CreateNews(c *fiber.Ctx) error {
	news := new(models.News)
	if err := c.BodyParser(news); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}
	result := database.DB.Create(&news)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": result.Error.Error()})
	}
	return c.JSON(news)
}

// GetAllNews retrieves all news items
func GetAllNews(c *fiber.Ctx) error {
	var news []models.News
	result := database.DB.Find(&news)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": result.Error.Error()})
	}
	return c.JSON(news)
}

// GetNewsByID retrieves a single news item by ID
func GetNewsByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	var news models.News
	result := database.DB.First(&news, id)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "News not found"})
	}
	return c.JSON(news)
}

// UpdateNews partially updates a news item by ID (PATCH)
func UpdateNews(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	var news models.News
	if err := database.DB.First(&news, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "News not found"})
	}

	// Use a map to store the request body
	var updateData map[string]interface{}
	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	// Update fields only if they are provided in the request
	if title, ok := updateData["title"].(string); ok {
		news.Title = title
	}
	if content, ok := updateData["content"].(string); ok {
		news.Content = content
	}
	if author, ok := updateData["author"].(string); ok {
		news.Author = author
	}
	if headerImageURL, ok := updateData["headerImageURL"].(string); ok {
		news.HeaderImageURL = headerImageURL
	}
	if commentCount, ok := updateData["commentCount"].(float64); ok { // JSON numbers are float64
		news.CommentCount = int(commentCount)
	}

	result := database.DB.Save(&news)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": result.Error.Error()})
	}
	return c.JSON(news)
}

// DeleteNews deletes a news item by ID
func DeleteNews(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	result := database.DB.Delete(&models.News{}, id)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": result.Error.Error()})
	}

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "News not found"})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
