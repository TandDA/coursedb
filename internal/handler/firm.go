package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (h *Handler) getAllFirmWithBooikingParams(c echo.Context) error {
	query := `
SELECT f.name
FROM booking b
JOIN firm AS f ON f.id = b.firm_id
WHERE ($1::date IS NULL OR $2::date IS NULL OR date_of_entry BETWEEN $1::date AND $2::date)
GROUP BY f.name
HAVING count(*) >= $3
	`

	from := c.QueryParam("from")
	to := c.QueryParam("to")
	up := c.QueryParam("up")
	var intUp int
	if up != "" {
		intUp, _ = strconv.Atoi(up)
	}
	rows, err := h.db.Query(query, from, to, intUp)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"err": err.Error()})
	}
	defer rows.Close()
	firms := make([]string, 0)
	for rows.Next() {
		var str string
		err = rows.Scan(&str)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"err": err.Error()})
		}
		firms = append(firms, str)
	}
	return c.JSON(http.StatusOK, map[string]any{"firms": firms})
}

func (h *Handler) getFirmsBookingCount(c echo.Context) error {
	query := `
		SELECT count(*)
		FROM booking b
		JOIN firm AS f ON f.id = b.firm_id
		WHERE ($1::date IS NULL OR $2::date IS NULL OR date_of_entry BETWEEN $1::date AND $2::date)
		AND f.name = $3
	`

	from := c.QueryParam("from")
	to := c.QueryParam("to")
	firmName := c.QueryParam("name")

	rows := h.db.QueryRow(query, from, to, firmName)
	var count int
	err := rows.Scan(&count)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"err": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]any{"count": count})
}
