package storage

import (
	"database/sql"
	"homeworks/homework_43/user/models"
)

type CardRepo struct {
	DB *sql.DB
}

func NewCardRepo(db *sql.DB) *CardRepo {
	return &CardRepo{db}
}

func (u *CardRepo) Create(card models.Card) error {
	_, err := u.DB.Exec("insert into card(number, user_id) values($1, $2)",
		card.Number, card.UserID)
	return err
}

func (u *CardRepo) Get(query string, arr interface{}) ([]models.Card, error) {
	rows, err := u.DB.Query(query, arr)
	if err != nil {
		return nil, err
	}

	cards := []models.Card{}
	for rows.Next() {
		card := models.Card{}
		err = rows.Scan(&card.Id, &card.Number, &card.UserID)
		if err != nil {
			return nil, err
		}
		cards = append(cards, card)
	}
	return cards, nil
}

func (u *CardRepo) GetById(id string) (models.Card, error) {
	card := models.Card{}
	err := u.DB.QueryRow("select * from card where card_id = $1 and deleted_at=0", id).
		Scan(&card.Id, &card.Number, &card.UserID, &card.DeletedAt)
	return card, err
}

func (u *CardRepo) Update(card models.Card, id string) error {
	_, err := u.DB.Exec("update users set number=$1, user_id=$2 where card_id=$3 and deleted_at=0",
		card.Number, card.UserID, id)
	return err
}

func (u *CardRepo) Delete(id string) error {
	_, err := u.DB.Exec(`update card set
deleted_at = date_part('epoch', current_timestamp)::INT
where id = $1 and deleted_at = 0
`, id)
	return err
}