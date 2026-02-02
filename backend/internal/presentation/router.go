package presentation

import (
	"ecolife-06/backend/internal/application"
	"ecolife-06/backend/internal/presentation/handlers"
	"ecolife-06/backend/internal/presentation/middleware"
	"github.com/gin-gonic/gin"
)

// SetupRouter настраивает маршруты
func SetupRouter(
	appService *application.ApplicationService,
	containerService *application.ContainerService,
	benefitService *application.BenefitService,
	tariffService *application.TariffService,
	allowedOrigins []string,
) *gin.Engine {
	router := gin.Default()

	// Middleware
	router.Use(middleware.CORS(allowedOrigins))
	router.Use(gin.Recovery())

	// Handlers
	appHandler := handlers.NewApplicationHandler(appService)
	containerHandler := handlers.NewContainerHandler(containerService)
	benefitHandler := handlers.NewBenefitHandler(benefitService)
	tariffHandler := handlers.NewTariffHandler(tariffService)

	// API routes
	api := router.Group("/api")
	{
		// Applications
		api.POST("/applications", appHandler.CreateApplication)
		api.GET("/applications", appHandler.GetAllApplications)
		api.GET("/applications/:id", appHandler.GetApplication)

		// Containers
		api.GET("/containers", containerHandler.GetAllContainers)
		api.GET("/containers/:id", containerHandler.GetContainer)

		// Benefits
		api.GET("/benefits", benefitHandler.GetAllBenefits)

		// Tariffs
		api.GET("/tariffs", tariffHandler.GetAllTariffs)
		api.GET("/tariffs/current", tariffHandler.GetCurrentTariff)
	}

	return router
}

