package storage

import (
	"database/sql"
	"homeworks/homework_43/metro/models"
)

type TerminalRepo struct {
	DB *sql.DB
}

func NewTerminalRepo(db *sql.DB) *TerminalRepo {
	return &TerminalRepo{db}
}

func (u *TerminalRepo) Create(terminal models.Terminal) error {
	_, err := u.DB.Exec("insert into terminal(station_id) values($1)",
		terminal.StationId)
	return err
}

func (u *TerminalRepo) Get(query string, arr []interface{}) ([]models.Terminal, error) {
	rows, err := u.DB.Query(query, arr...)
	if err != nil {
		return []models.Terminal{}, err
	}

	cards := []models.Terminal{}
	for rows.Next() {
		terminal := models.Terminal{}
		err = rows.Scan(&terminal.TerminalId, &terminal.StationId, &terminal.DeletedAt)
		if err != nil {
			return nil, err
		}
		cards = append(cards, terminal)
	}
	return cards, nil
}

func (u *TerminalRepo) GetById(id string) (models.Terminal, error) {
	terminal := models.Terminal{}
	err := u.DB.QueryRow("select * from terminal where station_id = $1 and deleted_at=0", id).
		Scan(&terminal.TerminalId, &terminal.StationId, &terminal.DeletedAt)
	return terminal, err
}

func (u *TerminalRepo) Update(terminal models.Terminal, id string) error {
	_, err := u.DB.Exec("update users set name=$1 where station_id=$2 and deleted_at=0",
		terminal.StationId, id)
	return err
}

func (u *TerminalRepo) Delete(id string) error {
	_, err := u.DB.Exec(`update terminal set
deleted_at = date_part('epoch', current_timestamp)::INT
where terminal_id = $1 and deleted_at = 0`, id)
	return err
}