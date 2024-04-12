package handler

import (
	"net/http"

	"github.com/TandDA/coursedb/internal/model"
	"github.com/labstack/echo/v4"
)

// type buildingDTO struct {
// 	Class          int    `json:"class"`
// 	NumberOfFloors int    `json:"number_of_floors"`
// 	Address        string `json:"address"`
// }

func (h *Handler) insertBuilding(c echo.Context) error {
	var bld model.Building
	if err := c.Bind(&bld); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"err": err.Error()})
	}
	id, err := h.service.Building.Save(bld)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"err": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"id": id})
}

func (h *Handler) getAllBuilding(c echo.Context) error {
	blds, err := h.service.Building.GetAll()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"err": err.Error()})
	}
	return c.JSON(http.StatusOK, blds)
}

func (h *Handler) updateBuilding(c echo.Context) error {
	var bld model.Building
	if err := c.Bind(&bld); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"err": err.Error()})
	}
	err := h.service.Building.Update(bld)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"err": err.Error()})
	}
	return c.NoContent(http.StatusOK)
}

func (h *Handler) deleteBuilding(c echo.Context) error {
	return nil
}
