package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type GuestCount struct { 
	FirstName string `json:"fiest_name"`
	LastName string `json:"last_name"`
	Count int `json:"count"`
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
	query  := `
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