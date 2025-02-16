package handler

import (
    "net/http"

    "delivery-planner-bot/backend/internal/domain/entity"
    "delivery-planner-bot/backend/internal/usecase/route"
    "github.com/gin-gonic/gin"
)

type RouteHandler struct {
    routeUseCase route.UseCase
}

func NewRouteHandler(routeUseCase route.UseCase) *RouteHandler {
    return &RouteHandler{
        routeUseCase: routeUseCase,
    }
}

func (h *RouteHandler) CreateRoute(c *gin.Context) {
    var route entity.Route
    if err := c.ShouldBindJSON(&route); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := h.routeUseCase.CreateRoute(c.Request.Context(), &route); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, route)
}

func (h *RouteHandler) GetRoute(c *gin.Context) {
    id := c.Param("id")
    route, err := h.routeUseCase.GetRoute(c.Request.Context(), id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, route)
}

func (h *RouteHandler) UpdateRoute(c *gin.Context) {
    id := c.Param("id")
    var route entity.Route
    if err := c.ShouldBindJSON(&route); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    route.ID = id
    if err := h.routeUseCase.UpdateRoute(c.Request.Context(), &route); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, route)
}

func (h *RouteHandler) DeleteRoute(c *gin.Context) {
    id := c.Param("id")
    if err := h.routeUseCase.DeleteRoute(c.Request.Context(), id); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.Status(http.StatusNoContent)
}

func (h *RouteHandler) ListRoutes(c *gin.Context) {
    routes, err := h.routeUseCase.ListRoutes(c.Request.Context())
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, routes)
}

func (h *RouteHandler) OptimizeRoute(c *gin.Context) {
    id := c.Param("id")
    if err := h.routeUseCase.OptimizeRoute(c.Request.Context(), id); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.Status(http.StatusOK)
}

func (h *RouteHandler) GetDriverRoutes(c *gin.Context) {
    driverID := c.Param("driverID")
    routes, err := h.routeUseCase.GetDriverRoutes(c.Request.Context(), driverID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, routes)
}
