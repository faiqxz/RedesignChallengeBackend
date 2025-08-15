package routes

import (
	"redesign/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// Grouping routes under /api/v1
	api := app.Group("/api/v1")

	// Upload route
	api.Post("/upload", handlers.UploadFile)

	// News routes
	news := api.Group("/news")
	news.Post("/", handlers.CreateNews)      // Create a new news item
	news.Get("/", handlers.GetAllNews)       // Get all news items
	news.Get("/:id", handlers.GetNewsByID)   // Get a single news item by ID
	news.Patch("/:id", handlers.UpdateNews)  // Update a news item by ID
	news.Delete("/:id", handlers.DeleteNews) // Delete a news item by ID

	// Research Team routes
	researchTeam := api.Group("/research-teams")
	researchTeam.Post("/", handlers.CreateResearchTeam)
	researchTeam.Get("/", handlers.GetAllResearchTeams)
	researchTeam.Get("/:id", handlers.GetResearchTeamByID)
	researchTeam.Patch("/:id", handlers.UpdateResearchTeam)
	researchTeam.Delete("/:id", handlers.DeleteResearchTeam)

	// Downloadable File routes
	downloadableFile := api.Group("/downloadable-files")
	downloadableFile.Post("/", handlers.CreateDownloadableFile)
	downloadableFile.Get("/", handlers.GetAllDownloadableFiles)
	downloadableFile.Get("/:id", handlers.GetDownloadableFileByID)
	downloadableFile.Patch("/:id", handlers.UpdateDownloadableFile)
	downloadableFile.Delete("/:id", handlers.DeleteDownloadableFile)

	// Certification routes
	certification := api.Group("/certifications")
	certification.Post("/", handlers.CreateCertification)
	certification.Get("/", handlers.GetAllCertifications)
	certification.Get("/:id", handlers.GetCertificationByID)
	certification.Patch("/:id", handlers.UpdateCertification)
	certification.Delete("/:id", handlers.DeleteCertification)

	// Gallery routes
	gallery := api.Group("/gallery")
	gallery.Post("/", handlers.CreateGalleryItem)
	gallery.Get("/", handlers.GetAllGalleryItems)
	gallery.Get("/:id", handlers.GetGalleryItemByID)
	gallery.Patch("/:id", handlers.UpdateGalleryItem)
	gallery.Delete("/:id", handlers.DeleteGalleryItem)
}
