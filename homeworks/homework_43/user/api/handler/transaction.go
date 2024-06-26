package handler

import (
	"fmt"
	"homeworks/homework_43/user/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) TransactionCreate(c *gin.Context) {
	transaction := models.Transaction{}
	err := c.ShouldBindJSON(&transaction)
	if err != nil {
		fmt.Println("Scan transaction ", err)
		return
	}

	err = h.Transaction.Create(transaction)
	if err != nil {
		fmt.Println("Create transaction ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"succesfull": "transaction created"})
}

func (h *Handler) TransactionGet(c *gin.Context) {
	var (
		params = make(map[string]interface{})
		arr    []interface{}
		limit  string
		offset = ""
	)

	query := `select * from transacton where deleted_at=0`
	// filter
	if len(c.Query("card_id")) > 0 {
		card_id := c.Query("card_id")
		params["card_id"] = card_id
		query += " and card_id = :card_id"
	}

	if len(c.Query("name")) > 0 {
		params["name"] = c.Query("name")
		query += " and name = :name"
	}

	if len(c.Query("amount")) > 0 {
		params["amount"] = c.Query("amount")
		query += " and amount = :amount "
	}

	if len(c.Query("transaction_type")) > 0 {
		params["transaction_type"] = c.Query("transaction_type")
		query += " and transaction_type = :transaction_type "
	}

	if len(c.Query("terminal_id")) > 0 {
		params["terminal_id"] = c.Query("terminal_id")
		query += " and terminal_id = :terminal_id "
	}

	if len(c.Query("limit")) > 0 {
		params["limit"] = c.Query("limit")
		limit = ` LIMIT :limit`
	}

	if len(c.Query("offset")) > 0 {
		params["offset"] = c.Query("offset")
		offset = ` OFFSET :offset`
	}

	query = query + limit + offset

	query, arr = ReplaceQueryParams(query, params)

	transacton, err := h.Transaction.Get(query, arr) // get with filter from db
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"ERROR IN GET Transaction": err})
		return
	}
	c.JSON(http.StatusOK, transacton)
}

func (h *Handler) TransactionGetByID(c *gin.Context) {
	id := c.Param("id")
	transaction, err := h.Transaction.GetById(id) // get by id from db
	if err != nil {
		fmt.Println("Error get by ID: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"ERROR IN GET BY ID Transaction": err})
		return
	}
	c.JSON(200, transaction)
}

// Handler struct for Update
func (h *Handler) TransactionUpdate(c *gin.Context) {
	transaction := models.Transaction{}
	err := c.ShouldBindJSON(&transaction)
	if err != nil {
		fmt.Println("Error in read from: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"ERROR IN READ FROM BODY Transaction": err})
		return
	}
	id := c.Param("id")
	err = h.Transaction.Update(transaction, id) //update transaction from db
	if err != nil {
		fmt.Println("ERROR In Update Users: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"ERROR IN UPDATE Transaction": err})
		return
	}
	c.JSON(200, gin.H{"Transaction": "Successfully Updated"})
}

// Handler struct for soft delete
func (h *Handler) TransactionDelete(c *gin.Context) {
	id := c.Param("id")
	err := h.Transaction.Delete(id) //soft delete from db
	if err != nil {
		fmt.Println("ERROR In Delete Users: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"ERROR IN SOFT DELETE Transaction": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Successfully": "Transaction deleted"})
}
