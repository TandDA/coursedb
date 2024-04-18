package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *Handler) getAllGuestsWithComplains(c echo.Context) error {
	complains, err := h.service.Guest.GetAllGuestsWithComplains()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"err": err.Error()})
	}
	return c.JSON(http.StatusOK, complains)
}

func (h *Handler) getAllGuests(c echo.Context) error {
	from := c.QueryParam("from")
	to := c.QueryParam("to")
	all, err := h.service.Guest.GetAll(from, to)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"err": err.Error()})
	}
	return c.JSON(http.StatusOK, all)
}
