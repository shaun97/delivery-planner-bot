package main

import (
	"context"
	"fmt"
	"log"

	"delivery-planner-bot/backend/internal/delivery/http/handler"
	"delivery-planner-bot/backend/internal/infrastructure/maps"
	"delivery-planner-bot/backend/internal/usecase/route"
	"delivery-planner-bot/backend/pkg/config"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	ctx := context.Background()

	// Initialize services
	mapService, err := maps.NewGoogleRoutesClient(ctx, cfg.GoogleMaps.APIKey)
	if err != nil {
		log.Fatalf("Failed to create maps client: %v", err)
	}

	geocodingService, err := maps.NewGoogleGeocodingClient(ctx, cfg.GoogleMaps.APIKey)
	if err != nil {
		log.Fatalf("Failed to create geocoding client: %v", err)
	}

	// Initialize use cases
	routeService := route.NewService(nil, mapService, geocodingService) // TODO: Add route repository

	// Initialize HTTP handlers
	routeHandler := handler.NewRouteHandler(routeService)

	// Setup Gin router
	router := gin.Default()
	fmt.Println("test")
	// Register route handlers
	v1 := router.Group("/api/v1")
	{
		routes := v1.Group("/routes")
		{
			routes.POST("/", routeHandler.CreateRoute)
			routes.GET("/", routeHandler.ListRoutes)
			routes.GET("/:id", routeHandler.GetRoute)
			routes.PUT("/:id", routeHandler.UpdateRoute)
			routes.DELETE("/:id", routeHandler.DeleteRoute)
			routes.POST("/:id/optimize", routeHandler.OptimizeRoute)
			routes.POST("/preview", routeHandler.PreviewRoute)
		}

		drivers := v1.Group("/drivers")
		{
			drivers.GET("/:driverID/routes", routeHandler.GetDriverRoutes)
		}
	}

	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// Start server
	addr := fmt.Sprintf(":%d", cfg.Server.Port)
	log.Printf("Server starting on %s", addr)
	if err := router.Run(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
