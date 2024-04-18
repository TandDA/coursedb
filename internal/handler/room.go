package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func (h *Handler) getAllFreeRooms(c echo.Context) error {
	floorNumber := getIntNullableParam(c, "floor_number")
	class := getIntNullableParam(c, "class")
	numberOfRooms := getIntNullableParam(c, "number_of_rooms")
	rooms, err := h.service.Room.GetAllFreeRooms(floorNumber, class, numberOfRooms)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"err": err.Error()})
	}
	return c.JSON(http.StatusOK, rooms)
}

func getIntNullableParam(c echo.Context, name string) int {
	param, _ := strconv.Atoi(c.QueryParam(name))
	if param == 0 {
		return -1
	}
	return param
}
