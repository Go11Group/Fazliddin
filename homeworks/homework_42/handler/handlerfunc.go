package handler

import (
	"bytes"
	"encoding/json"
	_ "hash"
	"io"
	"fmt"
	"net/http"
	_ "net/http"
	"strings"

	"homeworks/homework_42/structs"

	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	url := c.Request.URL.String()
	url = "http://localhost:8080" + url

	res, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		fmt.Println(err)
		return
	}

	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {
		c.JSON(http.StatusOK, gin.H{"success": "created"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "not created"})
	}
}

func Get(c *gin.Context) {
	url := "http://localhost:8080"
	check := c.Request.URL.String()
	url += check

	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	switch check {

	case "/user":
		users := []structs.User{}
		err = json.Unmarshal(body, &users)
		if err != nil {
			return
		}
		c.JSON(http.StatusOK, users)

	case "/course":
		course := []structs.Course{}
		err = json.Unmarshal(body, &course)
		if err != nil {
			return
		}
		c.JSON(http.StatusOK, course)

	case "/lesson":
		lesson := []structs.Lesson{}
		err = json.Unmarshal(body, &lesson)
		if err != nil {
			return
		}
		c.JSON(http.StatusOK, lesson)

	case "/enrollment":
		enroll := []structs.User{}
		err = json.Unmarshal(body, &enroll)
		if err != nil {
			return
		}
		c.JSON(http.StatusOK, enroll)
	}
}

func GetByID(c *gin.Context) {
	url := "http://localhost:8080"
	check := c.Request.URL.String()
	url += check
	res, err := http.Get(url)
	if err != nil {
		return
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return
	}
	switch {

	case strings.Contains(check, "/user"):
		users := structs.User{}
		err = json.Unmarshal(body, &users)
		if err != nil {
			return
		}
		c.JSON(http.StatusOK, users)

	case strings.Contains(check, "/course"):
		course := structs.Course{}
		err = json.Unmarshal(body, &course)
		if err != nil {
			return
		}
		c.JSON(http.StatusOK, course)

	case strings.Contains(check, "/lesson"):
		lesson := structs.Lesson{}
		err = json.Unmarshal(body, &lesson)
		if err != nil {
			return
		}
		c.JSON(http.StatusOK, lesson)

	case strings.Contains(check, "/enrollment"):
		enroll := structs.User{}
		err = json.Unmarshal(body, &enroll)
		if err != nil {
			return
		}
		c.JSON(http.StatusOK, enroll)
	}
}

func Update(c *gin.Context){
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	url := c.Request.URL.String()
	url = "http://localhost:8080" + url

	client := http.Client{}

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(body))

	res, err := client.Do(req)
	if err != nil{
		fmt.Print(err)
		return
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {
		c.JSON(http.StatusOK, gin.H{"success": "created"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "not created"})
	}
}

func Delete(c *gin.Context){
	url := "http://localhost:8080"
	url += c.Request.URL.String()

	client := http.Client{}

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil{
		fmt.Println(err)
		return
	}

	res, err := client.Do(req)

	defer res.Body.Close()
	if err != nil{
		fmt.Println(err)
		return
	}

	if res.StatusCode == http.StatusOK {
		c.JSON(http.StatusOK, gin.H{"success": "created"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "not created"})
	}

}