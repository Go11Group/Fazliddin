package models

type Terminal struct {
	TerminalId string `json:"terminal_id"`
	StationId  string `json:"station_id"`
	DeletedAt  int    `json:"deleted_at"`
}