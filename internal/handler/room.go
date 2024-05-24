package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type RoomDate struct {
	DateOfEntry     string `json:"date_of_entry"`
	DateOfDeparture string `json:"date_of_departure"`
}

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

func (h *Handler) getFreeRoomInfo(c echo.Context) error {
	id := c.QueryParam("id")
	query := `
	SELECT date_of_entry, date_of_departure FROM room r
	JOIN booking b ON r.id = b.room_id
	WHERE r.id = $1
	AND NOW() < date_of_departure
	
	`
	rows, _ := h.db.Query(query, id)
	dates := []RoomDate{}
	for rows.Next() {
		var date RoomDate
		err := rows.Scan(&date.DateOfEntry, &date.DateOfDeparture)
		if err != nil {
			return c.String(400, err.Error())
		}
		dates = append(dates, date)
	}
	return c.JSON(http.StatusOK, dates)
}

func (h *Handler) getPercentage(c echo.Context) error {
	query := `
	SELECT 
  COUNT(*) FILTER (WHERE firm_id IS NOT NULL) * 100.0 / COUNT(*) AS percentage
FROM 
  booking;

	`
	row := h.db.QueryRow(query)
	var perc float64
	row.Scan(&perc)
	return c.JSON(200, perc)
}

func (h *Handler) getFreeRoomInfoWithDate(c echo.Context) error {
	roomId := c.QueryParam("room_id")
	dateOfEntry := c.QueryParam("date_of_entry")
	dateOfDeparture := c.QueryParam("date_of_departure")
	query := `
	SELECT date_of_entry, date_of_departure FROM public.booking
WHERE room_id = $1
AND (
	date_of_entry BETWEEN $2::date AND $3::date
	OR
	date_of_departure BETWEEN $2::date AND $3::date
);
	`
	rows, _ := h.db.Query(query, roomId, dateOfEntry, dateOfDeparture)
	dates := []RoomDate{}
	for rows.Next() {
		var date RoomDate
		err := rows.Scan(&date.DateOfEntry, &date.DateOfDeparture)
		if err != nil {
			return c.String(400, err.Error())
		}
		dates = append(dates, date)
	}
	return c.JSON(http.StatusOK, dates)
}

func (h *Handler) getFreeRoomsOnCertainDate(c echo.Context) error {
	query := `
	SELECT DISTINCT r.* FROM booking AS b
JOIN room AS r ON r.id = b.room_id
WHERE ($1::date NOT BETWEEN date_of_entry AND date_of_departure)
AND (NOW()::date BETWEEN date_of_entry AND date_of_departure)

	`
	to := c.QueryParam("to")
	rows, err := h.db.Query(query, to)
	// TODO 6 request
}
