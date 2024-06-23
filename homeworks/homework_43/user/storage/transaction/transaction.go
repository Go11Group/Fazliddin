package storage

import (
	"database/sql"
	"homeworks/homework_43/user/models"
)

type TransactionRepo struct {
	DB *sql.DB
}

func NewTransactionRepo(db *sql.DB) *TransactionRepo {
	return &TransactionRepo{db}
}

func (t *TransactionRepo) Create(transaction models.Transaction) error {
	_, err := t.DB.Exec("insert into transaction(card_id, amount, terminal_id, transaction_type) values($1, $2, $3, $4)",
		transaction.CardId, transaction.Amount, transaction.TerminalId, transaction.TransactionType)
	return err
}

func (t *TransactionRepo) Get(query string) ([]models.Transaction, error) {
	rows, err := t.DB.Query(query)
	if err != nil {
		return nil, err
	}

	transactions := []models.Transaction{}
	for rows.Next() {
		transaction := models.Transaction{}
		err = rows.Scan(&transaction.Id, &transaction.CardId, &transaction.Amount,
			&transaction.TerminalId, &transaction.TransactionType)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}
	return transactions, nil
}

func (t *TransactionRepo) GetById(id string) (models.Transaction, error) {
	transaction := models.Transaction{}
	err := t.DB.QueryRow("select * from transactions where id = $1 and deleted_at=0", id).
		Scan(&transaction.Id, &transaction.CardId, &transaction.Amount,
			&transaction.TerminalId, &transaction.TransactionType)
	return transaction, err
}

func (t *TransactionRepo) Update(transaction models.Transaction, id string) error {
	_, err := t.DB.Exec("update transactions set name=$1, age=$2, phone=$3 where id=$4 and deleted_at=0",
	transaction.CardId, transaction.Amount,
	transaction.TerminalId, transaction.TransactionType, id)
	return err
}

func (t *TransactionRepo) Delete(id int) error {
	_, err := t.DB.Exec(`update transactions set
deleted_at = date_part('epoch', current_timestamp)::INT
where id = $1 and deleted_at = 0`, id)
	return err
}