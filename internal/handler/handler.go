package handler

import (
	"github.com/TandDA/coursedb/internal/service"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Start() {
	e := echo.New()
	e.POST("/building", h.insertBuilding)
	e.GET("/building", h.getAllBuilding)
	e.PUT("/building", h.updateBuilding)
	e.DELETE("/building", h.deleteBuilding)
	e.Start(":8080")
}
