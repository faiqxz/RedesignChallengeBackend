package handlers

import (
	"github.com/gofiber/fiber/v2"
	"redesign/database"
	"redesign/models"
	"strconv"
)

// CreateCertification creates a new certification
func CreateCertification(c *fiber.Ctx) error {
	cert := new(models.Certification)
	if err := c.BodyParser(cert); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}
	result := database.DB.Create(&cert)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": result.Error.Error()})
	}
	return c.JSON(cert)
}

// GetAllCertifications retrieves all certifications
func GetAllCertifications(c *fiber.Ctx) error {
	var certs []models.Certification
	result := database.DB.Find(&certs)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": result.Error.Error()})
	}
	return c.JSON(certs)
}

// GetCertificationByID retrieves a single certification by ID
func GetCertificationByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	var cert models.Certification
	result := database.DB.First(&cert, id)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Certification not found"})
	}
	return c.JSON(cert)
}

// UpdateCertification partially updates a certification by ID (PATCH)
func UpdateCertification(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	var cert models.Certification
	if err := database.DB.First(&cert, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Certification not found"})
	}

	var updateData map[string]interface{}
	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	if certName, ok := updateData["certificationName"].(string); ok {
		cert.CertificationName = certName
	}
	if org, ok := updateData["issuingOrganization"].(string); ok {
		cert.IssuingOrganization = org
	}
	if desc, ok := updateData["description"].(string); ok {
		cert.Description = desc
	}
	if link, ok := updateData["certificationLink"].(string); ok {
		cert.CertificationLink = link
	}

	result := database.DB.Save(&cert)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": result.Error.Error()})
	}
	return c.JSON(cert)
}

// DeleteCertification deletes a certification by ID
func DeleteCertification(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	result := database.DB.Delete(&models.Certification{}, id)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": result.Error.Error()})
	}

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Certification not found"})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
