package handler

import "github.com/TandDA/coursedb/internal/service"

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() {
	
}