package handler

import (
	"database/sql"
	"log"

	"github.com/TandDA/coursedb/internal/service"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	service *service.Service
	db      *sql.DB
}

func NewHandler(service *service.Service, db *sql.DB) *Handler {
	return &Handler{service: service, db: db}
}

func (h *Handler) Start() {
	e := echo.New()
	e.POST("/building", h.insertBuilding)
	e.GET("/building", h.getAllBuilding)
	e.PUT("/building", h.updateBuilding)
	e.DELETE("/building", h.deleteBuilding)

	firmGroup := e.Group("/firm")
	firmGroup.GET("/date", h.getAllFirmWithBooikingParams)
	firmGroup.GET("/count", h.getFirmsBookingCount)

	roomGroup := e.Group("/room")
	roomGroup.GET("/free", h.getAllFreeRooms)
	roomGroup.GET("/free-detail", h.getFreeRoomInfo)
	roomGroup.GET("/date", h.getFreeRoomInfoWithDate)
	roomGroup.GET("/percentage", h.getPercentage)
	roomGroup.GET("/todate", h.getFreeRoomsOnCertainDate)

	e.GET("/guest", h.getAllGuests)
	guestGroup := e.Group("/guest")
	guestGroup.GET("/complains", h.getAllGuestsWithComplains)
	guestGroup.GET("/popular", h.getAllPopularGuest)
	guestGroup.GET("/booking", h.getGuestBooking)
	guestGroup.GET("/byroom", h.getGuestsByRoom)

	err := e.Start(":8080")
	if err != nil {
		log.Print(err)
	}
}
