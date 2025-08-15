package handlers

import (
	"github.com/gofiber/fiber/v2"
	"redesign/database"
	"redesign/models"
	"strconv"
)

// CreateDownloadableFile creates a new downloadable file entry
func CreateDownloadableFile(c *fiber.Ctx) error {
	file := new(models.DownloadableFile)
	if err := c.BodyParser(file); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}
	result := database.DB.Create(&file)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": result.Error.Error()})
	}
	return c.JSON(file)
}

// GetAllDownloadableFiles retrieves all downloadable files
func GetAllDownloadableFiles(c *fiber.Ctx) error {
	var files []models.DownloadableFile
	result := database.DB.Find(&files)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": result.Error.Error()})
	}
	return c.JSON(files)
}

// GetDownloadableFileByID retrieves a single downloadable file by ID
func GetDownloadableFileByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	var file models.DownloadableFile
	result := database.DB.First(&file, id)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "File not found"})
	}
	return c.JSON(file)
}

// UpdateDownloadableFile partially updates a downloadable file by ID (PATCH)
func UpdateDownloadableFile(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	var file models.DownloadableFile
	if err := database.DB.First(&file, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "File not found"})
	}

	var updateData map[string]interface{}
	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	if fileName, ok := updateData["fileName"].(string); ok {
		file.FileName = fileName
	}
	if description, ok := updateData["description"].(string); ok {
		file.Description = description
	}
	if fileURL, ok := updateData["fileURL"].(string); ok {
		file.FileURL = fileURL
	}
	if fileSize, ok := updateData["fileSize"].(float64); ok { // JSON numbers are float64
		file.FileSize = int64(fileSize)
	}
	if fileType, ok := updateData["fileType"].(string); ok {
		file.FileType = fileType
	}

	result := database.DB.Save(&file)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": result.Error.Error()})
	}
	return c.JSON(file)
}

// DeleteDownloadableFile deletes a downloadable file by ID
func DeleteDownloadableFile(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	result := database.DB.Delete(&models.DownloadableFile{}, id)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": result.Error.Error()})
	}

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "File not found"})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
