package storage

import (
	"database/sql"
	"homeworks/homework_43/metro/models"
)

type StationRepo struct {
	DB *sql.DB
}

func NewStationRepo(db *sql.DB) *StationRepo {
	return &StationRepo{db}
}

func (u *StationRepo) Create(station models.Station) error {
	_, err := u.DB.Exec("insert into station(name) values($1)",
		station.Name)
	return err
}

func (u *StationRepo) Get(query string, arr []interface{}) ([]models.Station, error) {

	rows, err := u.DB.Query(query, arr...)
	if err != nil {
		return nil, err
	}

	cards := []models.Station{}
	for rows.Next() {
		station := models.Station{}
		err = rows.Scan(&station.StatinId, &station.Name, &station.DeletedAt)
		if err != nil {
			return nil, err
		}
		cards = append(cards, station)
	}
	return cards, nil
}

func (u *StationRepo) GetById(id string) (models.Station, error) {
	station := models.Station{}
	err := u.DB.QueryRow("select * from station where station_id = $1 and deleted_at=0", id).
		Scan(&station.StatinId, &station.Name, &station.DeletedAt)
	return station, err
}

func (u *StationRepo) Update(station models.Station, id string) error {
	_, err := u.DB.Exec("update users set name=$1 where station_id=$2 and deleted_at=0",
		station.Name, id)
	return err
}

func (u *StationRepo) Delete(id string) error {
	_, err := u.DB.Exec(`update station set
deleted_at = date_part('epoch', current_timestamp)::INT
where station_id = $1 and deleted_at = 0`, id)
	return err
}