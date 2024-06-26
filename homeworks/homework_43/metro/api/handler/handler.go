package handler

import (
	storage "homeworks/homework_43/metro/storage/station"
	"strconv"
	"strings"
)

type HandlerRepo struct {
	Station  *storage.StationRepo
	Terminal *storage.TerminalRepo
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
