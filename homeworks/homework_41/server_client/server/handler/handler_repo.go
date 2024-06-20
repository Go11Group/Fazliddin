package handler

import "homework_41/server_client/server/postgres"


type HandlerRepo struct {
	User *postgres.UserRepo
}