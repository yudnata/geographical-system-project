package main

import (
	"backend/internal/config"
	"backend/internal/database"
	"backend/internal/middleware"
	"backend/internal/modules/auth"
	"backend/internal/modules/points"
	"backend/internal/router"
	"backend/pkg/upload"
	"github.com/gofiber/fiber/v3"

	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"log"
)

func main() {
	cfg := config.Load()
	db := database.Connect(cfg.DatabaseURL)
	database.Migrate(cfg.DatabaseURL)

	authRepo := auth.NewRepository(db)
	authServ := auth.NewService(authRepo, cfg)
	authHand := auth.NewHandler(authServ)

	cldService, err := upload.NewCloudinaryService()
	if err != nil {
		log.Printf("Warning: Failed to initialize Cloudinary: %v", err)
	}

	pointsRepo := points.NewRepository(db)
	pointsServ := points.NewService(pointsRepo)
	pointsHand := points.NewHandler(pointsServ, cldService)


	authMiddleware := middleware.AuthRequired(cfg)

	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New())

	router.Setup(app, authHand, pointsHand, authMiddleware)

	log.Fatal(app.Listen(":" + cfg.Port))
}
