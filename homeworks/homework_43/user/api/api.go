package api

import (
	"database/sql"
	hn "homeworks/homework_43/user/api/handler"
	cr "homeworks/homework_43/user/storage/card"
	tr "homeworks/homework_43/user/storage/transaction"
	us "homeworks/homework_43/user/storage/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Handler(db *sql.DB) *http.Server {

	u := us.NewUserRepo(db)
	c := cr.NewCardRepo(db)
	t := tr.NewTransactionRepo(db)
	
	handler := hn.Handler{u, c, t}

	gin := gin.Default()
	user := gin.Group("/user")
	card := gin.Group("/card")
	transaction := gin.Group("/transaction")

	user.POST("", handler.CreateUser)
	user.GET("", handler.UserGet)
	user.GET("/:id", handler.UserGetByID)
	user.PUT("/:id", handler.UserUpdate)
	user.DELETE("/:id", handler.UserDelete)

	card.POST("", handler.CardCreate)
	card.GET("", handler.CardGet)
	card.GET("/:id", handler.CardGetByID)
	card.PUT("/:id", handler.CardUpdate)
	card.DELETE("/:id", handler.CardDelete)

	transaction.POST("", handler.TransactionCreate)
	transaction.GET("", handler.TransactionGet)
	transaction.GET("/:id", handler.TransactionGetByID)
	transaction.PUT("/:id", handler.TransactionUpdate)
	transaction.DELETE("/:id", handler.TransactionDelete)

	return &http.Server{Addr: ":8080", Handler: gin}
}
