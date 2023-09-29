package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type user struct {
	ID     int  `json:"id"`
	Name  string  `json:"name"`
}

var users = []user{
	{ID: 1, Name: "Michel"},
	{ID: 2, Name: "Adrian"},
	{ID: 3, Name: "Claud"},
}

func main() {
	router := gin.Default()
	router.GET("/users", getUsers)
	router.GET("/users/:id", getUserByID)
    router.POST("/users", postUsers)

	router.Run("localhost:5001")
}

func getUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}

func postUsers(c *gin.Context) {
	var newUser user

	// Call BindJSON to bind the received JSON to
	if err := c.BindJSON(&newUser); err != nil {
			return
	}

	users = append(users, newUser)
	c.IndentedJSON(http.StatusCreated, newUser)
}

func getUserByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err == nil {
		for _, a := range users {
				if a.ID == id {
						c.IndentedJSON(http.StatusOK, a)
						return
				}
		}
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "error parsing ur param"})
	}
}