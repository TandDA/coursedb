package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *Handler) getAllFreeRooms(c echo.Context) error {
	rooms, err := h.service.Room.GetAllFreeRooms()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"err": err.Error()})
	}
	return c.JSON(http.StatusOK, rooms)
}
