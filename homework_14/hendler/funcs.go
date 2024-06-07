package hendler

import (
	"encoding/json"
	"fmt"
	"io"
	"model/product"
	"model/user"
	"net/http"
)

type Handler struct{
	User *user.UserRepo
	Product *product.ProductRepo
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request){
	b, err := io.ReadAll(r.Body)
	if err != nil{
		fmt.Println(err)
		w.Write([]byte("Error in reading body"))
	}
	
	users := []user.User{}
	err = json.Unmarshal(b, &users)
	fmt.Println(users)
	if err != nil{
		fmt.Println("json: ", err)
	}	
	for _, i := range users{
		h.User.Create(i)
	}
}

func (h *Handler) Get (w http.ResponseWriter, r *http.Request){
	users, err := h.User.Get()
	if err != nil{
		fmt.Println("Get : ", err)
		return
	}

	data, err := json.Marshal(users)
	if err != nil{
		fmt.Println("Get : ", err)
		return
	}
	w.Write(data)
}