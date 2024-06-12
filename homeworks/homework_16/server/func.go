package server

import (
	"fmt"
	"model/problems"
	"model/user"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type HandlerRepo struct{
	User *user.UserRepo
	Problem *problems.ProblemRepo
}

func (h *HandlerRepo) UserCreate(c *gin.Context) {
	u := user.User{}
	err := c.ShouldBind(&u)
	if err != nil{
		fmt.Println("ERROR JSON : ", err)
	}
	err = h.User.Create(u)
	if err != nil{
		fmt.Println("ERROR JSON: ", err)
		return
	}
}

func (h *HandlerRepo) UserGetByID(c *gin.Context){
	a := c.Param("id")
	id, err := strconv.Atoi(a)
	if err != nil{
		fmt.Println("ATOI ERROR : ", err)
		return
	}

	user, err := h.User.GetById(id)
	if err != nil{
		fmt.Println("ERROR IN GET_ID : ", err)
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *HandlerRepo) UserGet(c *gin.Context){
	users, err := h.User.Get()
	if err != nil{
		fmt.Println("ERROR IN GET : ", err)
	}
	c.JSON(http.StatusOK, users)
}

func (h *HandlerRepo) UserUpdate(c *gin.Context){
	u := user.User{}
	err := c.ShouldBind(&u)
	if err != nil{
		fmt.Println("ERROR JSON : ", err)
	}
	err = h.User.Update(u)
	if err != nil{
		fmt.Println("DATABASE ERROR : ", err)
		return
	}
}

func (h *HandlerRepo) UserDelete(c *gin.Context){
	a := c.Param("id")

	id, err := strconv.Atoi(a)
	if err != nil{
		fmt.Println("KALLENGA BL  ATOI ERROR : ", err)
		return
	}
	err = h.User.Delete(id)
	if err != nil{
		fmt.Println("ERROR DELETE : ", err)
		return
	}
}



func (h *HandlerRepo) ProblemCreate(c *gin.Context) {
	p := problems.Problem{}
	err := c.ShouldBind(&p)
	if err != nil{
		fmt.Println("ERROR JSON : ", err)
	}
	fmt.Println(p)
	err = h.Problem.Create(p)
	if err != nil{
		fmt.Println("ERROR JSON: ", err)
		return
	}
}

func (h *HandlerRepo) ProblemGetByID(c *gin.Context){
	v := c.Param("id")
	a, err := strconv.Atoi(v)
	if err != nil{
		fmt.Println("ATOI ERROR : ", err)
		return
	}

	problem, err := h.Problem.GetById(a)
	if err != nil{
		fmt.Println("ERROR IN GET_ID : ", err)
		return
	}
	c.JSON(http.StatusOK, problem)
}

func (h *HandlerRepo) ProblemGet(c *gin.Context){
	problem, err := h.Problem.Get()
	if err != nil{
		fmt.Println("ERROR IN GET : ", err)
	}
	c.ShouldBind(problem)
}

func (h *HandlerRepo) ProblemUpdate(c *gin.Context){
	p := problems.Problem{}
	err := c.ShouldBind(&p)
	if err != nil{
		fmt.Println("ERROR JSON : ", err)
	}
	err = h.Problem.Update(p)
	if err != nil{
		fmt.Println("DATABASE ERROR : ", err)
		return
	}
}

func (h *HandlerRepo) ProblemDelete(c *gin.Context){
	v := c.Param("Id")
	id, err := strconv.Atoi(v)
	if err != nil{
		fmt.Println("ERROR IN Aoi : ", err)
	}
	h.Problem.Delete(id)
}