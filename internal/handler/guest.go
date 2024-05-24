package handler

import (
	"net/http"

	"github.com/TandDA/coursedb/internal/model"
	"github.com/labstack/echo/v4"
)

type GuestCount struct {
	FirstName string `json:"fiest_name"`
	LastName  string `json:"last_name"`
	Count     int    `json:"count"`
}

type RoomDateAndId struct {
	DateOfEntry     string `json:"date_of_entry"`
	DateOfDeparture string `json:"date_of_departure"`
	RoomId          string `json:"room_id"`
}

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

func (h *Handler) getAllPopularGuest(c echo.Context) error {
	buildingId := c.QueryParam("building_id")
	query := `
	SELECT g.first_name, g.last_name, COUNT(*) FROM booking AS b
	JOIN room AS r ON b.room_id = r.id
	JOIN guest AS g ON b.guest_id = g.id
	JOIN floor AS f ON r.floor_id = f.id
	WHERE $1 = '' OR building_id = $1
	GROUP BY (g.first_name, g.last_name)
	ORDER BY count DESC
	`

	rows, err := h.db.Query(query, buildingId)
	if err != nil {
		return c.String(400, err.Error())
	}
	var ans []GuestCount
	for rows.Next() {
		var a GuestCount
		err = rows.Scan(&a.FirstName, &a.LastName, &a.Count)
		if err != nil {
			return c.String(400, err.Error())
		}
		ans = append(ans, a)
	}
	return c.JSON(http.StatusOK, ans)
}

func (h *Handler) getGuestBooking(c echo.Context) error {
	query := `
	SELECT b.date_of_entry, b.date_of_departure, b.room_id FROM guest AS g
JOIN booking AS b ON b.guest_id = g.id
WHERE g.id = $1
	`
	id := c.QueryParam("id")
	rows, err := h.db.Query(query, id)
	if err != nil {
		return c.JSON(500, err.Error())
	}
	var ans []RoomDateAndId
	for rows.Next() {
		var a RoomDateAndId
		err = rows.Scan(&a.DateOfEntry, &a.DateOfDeparture, &a.RoomId)
		if err != nil {
			return c.String(400, err.Error())
		}
		ans = append(ans, a)
	}
	return c.JSON(http.StatusOK, ans)
}

func (h *Handler) getGuestsByRoom(c echo.Context) error {
	query := `
	SELECT DISTINCT g.* FROM room AS r
JOIN floor AS f ON r.floor_id = f.id
JOIN booking AS b ON b.room_id = r.id
JOIN guest AS g ON b.guest_id = g.id
WHERE ($1::date IS NULL OR $2::date IS NULL OR b.date_of_entry BETWEEN $1::date AND $2::date) 
AND floor_number = $3 AND number_of_rooms = $4 AND regular_price < $5
	 `

	from := c.QueryParam("from")
	to := c.QueryParam("to")
	floorNumber := c.QueryParam("floor_number")
	numberOfRooms := c.QueryParam("number_of_rooms")
	regularPrice := c.QueryParam("regular_price")

	rows, err := h.db.Query(query, from, to, floorNumber, numberOfRooms, regularPrice)
	if err != nil {
		return c.JSON(500, err.Error())
	}
	var ans []model.Guest
	for rows.Next() {
		var a model.Guest
		err = rows.Scan(&a.Id, &a.FirstName, &a.LastName, &a.DateOfEntry)
		if err != nil {
			return c.String(400, err.Error())
		}
		ans = append(ans, a)
	}
	return c.JSON(http.StatusOK, ans)
}
