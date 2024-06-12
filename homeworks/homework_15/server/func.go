package server

import (
	"encoding/json"
	"fmt"
	"model/problems"
	"model/user"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

type HandlerRepo struct{
	User *user.UserRepo
	Problem *problems.ProblemRepo
}

func (h *HandlerRepo) UserCreate(w http.ResponseWriter, r *http.Request) {
	u := user.User{}
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil{
		fmt.Println("ERROR JSON : ", err)
	}
	err = h.User.Create(u)
	if err != nil{
		fmt.Println("ERROR JSON: ", err)
		return
	}
}

func (h *HandlerRepo) UserGetByID(w http.ResponseWriter, r *http.Request){
	v := r.URL.Query().Get("id")
	a, err := strconv.Atoi(v)
	if err != nil{
		fmt.Println("ATOI ERROR : ", err)
		return
	}

	user, err := h.User.GetById(a)
	if err != nil{
		fmt.Println("ERROR IN GET_ID : ", err)
		return
	}
	err = json.NewEncoder(w).Encode(user)
	if err != nil{
		fmt.Println("ERROR IN ENCODE : ", err)
	}
}

func (h *HandlerRepo) UserGet(w http.ResponseWriter, r *http.Request){
	users, err := h.User.Get()
	if err != nil{
		fmt.Println("ERROR IN GET : ", err)
	}
	json.NewEncoder(w).Encode(users)
}

func (h *HandlerRepo) UserUpdate(w http.ResponseWriter, r *http.Request){
	u := user.User{}
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil{
		fmt.Println("ERROR JSON : ", err)
	}
	err = h.User.Update(u)
	if err != nil{
		fmt.Println("DATABASE ERROR : ", err)
		return
	}
}

func (h *HandlerRepo) UserDelete(w http.ResponseWriter, r *http.Request){
	v := mux.Vars(r)["id"]
	id, err := strconv.Atoi(v)
	if err != nil{
		fmt.Println("ERROR IN Aoi : ", err)
	}
	h.User.Delete(id)
}



func (h *HandlerRepo) ProblemCreate(w http.ResponseWriter, r *http.Request) {
	p := problems.Problem{}
	err := json.NewDecoder(r.Body).Decode(&p)
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

func (h *HandlerRepo) ProblemGetByID(w http.ResponseWriter, r *http.Request){
	v := mux.Vars(r)["id"]
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
	err = json.NewEncoder(w).Encode(problem)
	if err != nil{
		fmt.Println("ERROR IN ENCODE : ", err)
	}
}

func (h *HandlerRepo) ProblemGet(w http.ResponseWriter, r *http.Request){
	problem, err := h.Problem.Get()
	if err != nil{
		fmt.Println("ERROR IN GET : ", err)
	}
	json.NewEncoder(w).Encode(problem)
}

func (h *HandlerRepo) ProblemUpdate(w http.ResponseWriter, r *http.Request){
	p := problems.Problem{}
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil{
		fmt.Println("ERROR JSON : ", err)
	}
	err = h.Problem.Update(p)
	if err != nil{
		fmt.Println("DATABASE ERROR : ", err)
		return
	}
}

func (h *HandlerRepo) ProblemDelete(w http.ResponseWriter, r *http.Request){
	v := mux.Vars(r)["id"]
	id, err := strconv.Atoi(v)
	if err != nil{
		fmt.Println("ERROR IN Aoi : ", err)
	}
	h.Problem.Delete(id)
}