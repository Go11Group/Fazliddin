package user

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func Post(user User) error {
	us, err := json.Marshal(user)
	if err != nil {
		return err
	}

	res, err := http.Post("http://localhost:8080/user", "application/json", bytes.NewBuffer(us))
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode == http.StatusCreated || res.StatusCode == http.StatusOK {
		return errors.New("User created sucessfully")
	}

	return errors.New("User not created")
}


func Get() ([]User, error){
	res, err := http.Get("http://localhost:8080/user")
	if err != nil{
		return nil, err
	}

	b, err := io.ReadAll(res.Body)
	if err != nil{
		return nil, err
	}

	users := []User{}
	err = json.Unmarshal(b, &users)

	return users, err
}

func GetByID(id string) (User, error){
	a := "http://localhost:8080/user/" + id
	res, err := http.Get(a)
	if err != nil {
		return User{}, err
	}

	b, err := io.ReadAll(res.Body)
	if err != nil{
		return User{}, nil
	}
	
	user := User{}
	err = json.Unmarshal(b, &user)
	return user, err
}

func Delete(id string) error{
	a :=  "http://localhost:8080/user/" + id
	req, err := http.NewRequest(http.MethodDelete, a, nil)
	if err != nil {
		return err
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		return errors.New("failed to delete user: " + resp.Status)
	}

	return nil
}

func Update(id string) error{
	a := "http://localhost:8080/user/" + id

	user := User{
		FirstName: "kimdur",
		LastName: "Ahmatov",
		Age: 5,
	}
	b, err := json.Marshal(user)
	if err != nil{
		return err
	}
	client := http.Client{}
	
	res, err := http.NewRequest("PUT", a,  bytes.NewBuffer(b))
	if err != nil{
		return err
	}

	_, err = client.Do(res)
	if err != nil{
		return err
	}
	return nil
}