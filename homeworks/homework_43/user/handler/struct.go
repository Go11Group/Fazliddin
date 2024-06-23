package handler

import (
	cr "homeworks/homework_43/user/storage/card"
	tr "homeworks/homework_43/user/storage/transaction"
	us "homeworks/homework_43/user/storage/user"
)

type HandlerRepo struct {
	User        *us.UserRepo
	Card        *cr.CardRepo
	Transaction *tr.TransactionRepo
}
