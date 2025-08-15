package handlers

import (
	"github.com/gofiber/fiber/v2"
	"redesign/database"
	"redesign/models"
	"strconv"
)

// CreateResearchTeam creates a new research team
func CreateResearchTeam(c *fiber.Ctx) error {
	team := new(models.ResearchTeam)
	if err := c.BodyParser(team); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}
	result := database.DB.Create(&team)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": result.Error.Error()})
	}
	return c.JSON(team)
}

// GetAllResearchTeams retrieves all research teams
func GetAllResearchTeams(c *fiber.Ctx) error {
	var teams []models.ResearchTeam
	result := database.DB.Find(&teams)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": result.Error.Error()})
	}
	return c.JSON(teams)
}

// GetResearchTeamByID retrieves a single research team by ID
func GetResearchTeamByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	var team models.ResearchTeam
	result := database.DB.First(&team, id)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Research team not found"})
	}
	return c.JSON(team)
}

// UpdateResearchTeam partially updates a research team by ID (PATCH)
func UpdateResearchTeam(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	var team models.ResearchTeam
	if err := database.DB.First(&team, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Research team not found"})
	}

	var updateData map[string]interface{}
	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	if teamName, ok := updateData["teamName"].(string); ok {
		team.TeamName = teamName
	}
	if teamLead, ok := updateData["teamLead"].(string); ok {
		team.TeamLead = teamLead
	}
	if members, ok := updateData["members"].(string); ok {
		team.Members = members
	}
	if researchDescription, ok := updateData["researchDescription"].(string); ok {
		team.ResearchDescription = researchDescription
	}
	if image, ok := updateData["image"].(string); ok {
		team.Image = image
	}

	result := database.DB.Save(&team)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": result.Error.Error()})
	}
	return c.JSON(team)
}

// DeleteResearchTeam deletes a research team by ID
func DeleteResearchTeam(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	result := database.DB.Delete(&models.ResearchTeam{}, id)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": result.Error.Error()})
	}

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Research team not found"})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
