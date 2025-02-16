package main

import (
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

    // Initialize services
    mapService, err := maps.NewGoogleMapsClient(cfg.GoogleMaps.APIKey)
    if err != nil {
        log.Fatalf("Failed to create maps client: %v", err)
    }

    // Initialize use cases
    routeService := route.NewService(nil, mapService) // TODO: Add route repository

    // Initialize HTTP handlers
    routeHandler := handler.NewRouteHandler(routeService)

    // Setup Gin router
    router := gin.Default()

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
        }

        drivers := v1.Group("/drivers")
        {
            drivers.GET("/:driverID/routes", routeHandler.GetDriverRoutes)
        }
    }

    // Start server
    addr := fmt.Sprintf(":%d", cfg.Server.Port)
    log.Printf("Server starting on %s", addr)
    if err := router.Run(addr); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}
