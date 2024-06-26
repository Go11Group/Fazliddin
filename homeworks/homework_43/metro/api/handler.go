package api

import (
	"database/sql"
	"homeworks/homework_43/metro/api/handler"
	storage "homeworks/homework_43/metro/storage/station"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Handler(db *sql.DB) *http.Server {
	
	s := storage.NewStationRepo(db)
	t := storage.NewTerminalRepo(db)

	h := handler.HandlerRepo{s, t}

	gin := gin.Default()

	station := gin.Group("/station")
	terminal := gin.Group("/terminal")
	

	station.POST("", h.StationCreate)
	station.GET("", h.StationdGet)
	station.GET("/:id", h.StationGetByID)
	station.PUT("/:id", h.StationUpdate)
	station.DELETE("/:id", h.StationDelete)

	terminal.POST("", h.TerminalCreate)
	terminal.GET("", h.TerminalGet)
	terminal.GET("/:id", h.TerminalGetByID)
	terminal.PUT("/:id", h.TerminalUpdate)
	terminal.DELETE("/:id", h.TerminalDelete)

	return &http.Server{Addr: ":8081", Handler: gin}
}