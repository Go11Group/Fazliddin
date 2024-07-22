package handler

import (
	"github.com/gin-gonic/gin"
	"homework_62/models"
	"homework_62/service"
	"log"
	"net/http"
)

type Handler struct {
	Service *service.Service
}

func (h Handler) CreataUser(c *gin.Context) {

	var user models.UserReq
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
	}

	err = h.Service.CreateUser(user)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
}

func (h *Handler) CreateBook(c *gin.Context) {
	var book models.Book

	err := c.ShouldBindJSON(&book)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Println(err)
		return
	}

	message, err := h.Service.PostBook(book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Println(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": message})
}

func (h *Handler) GetBooks(c *gin.Context) {
	books, err := h.Service.GetBooks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"books": books})
}

func (h *Handler) GetBook(c *gin.Context) {
	name := c.Param("name")

	books, err := h.Service.GetBook(models.BookName{Name: name})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Println(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"books": books})
}

func (h *Handler) DeleteBook(c *gin.Context) {
	name := c.Param("name")

	message, err := h.Service.DeleteBook(models.BookName{Name: name})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Println(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": message})
}
