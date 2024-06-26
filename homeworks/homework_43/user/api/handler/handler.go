package handler

import (
	cr "homeworks/homework_43/user/storage/card"
	tr "homeworks/homework_43/user/storage/transaction"
	us "homeworks/homework_43/user/storage/user"
	"strconv"
	"strings"
)

type Handler struct {
	User        *us.UserRepo
	Card        *cr.CardRepo
	Transaction *tr.TransactionRepo
}

func ReplaceQueryParams(namedQuery string, params map[string]interface{}) (string, []interface{}) {
	var (
		i    int = 1
		args []interface{}
	)
	for k, v := range params {
		if k != "" && strings.Contains(namedQuery, ":"+k) {
			namedQuery = strings.ReplaceAll(namedQuery, ":"+k, "$"+strconv.Itoa(i))
			args = append(args, v)
			i++
		}
	}
	return namedQuery, args
}
