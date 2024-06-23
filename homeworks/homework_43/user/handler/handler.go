package handler

import (
	"database/sql"
	cr "homeworks/homework_43/user/storage/card"
	tr "homeworks/homework_43/user/storage/transaction"
	us "homeworks/homework_43/user/storage/user"
	"net/http"
)

func Handler(db *sql.DB) *http.Server {

	u := us.NewUserRepo(db)
	c := cr.NewCardRepo(db)
	t := tr.NewTransactionRepo(db)
	
	handler := HandlerRepo{u, c, t}

}
