package main

import (
	"net/http"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type User struct {
	ID     int  `json:"id"`
	Name  string  `json:"name"`
}

type UserRequestBody struct {
	Name  string  `json:"name"`
}

var users = []User{
	{ID: 1, Name: "Michel"},
	{ID: 2, Name: "Adrian"},
	{ID: 3, Name: "Claud"},
}

var uID = 4

func main() {
	router := gin.Default()

	// https://stackoverflow.com/questions/72155224/golang-gin-middleware-with-cors-not-working
	router.Use(cors.Default())

	g := router.Group("/api/users")
	
	g.GET("", getUsers)
	g.GET(":id", getUserByID)
	g.POST("", postUsers)
	g.DELETE(":id", deleteUser)

	router.Run(":5001")
}

func getUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}

func postUsers(c *gin.Context) {
	var newUser UserRequestBody

	// Call BindJSON to bind the received JSON to
	if err := c.BindJSON(&newUser); err != nil {
			return
	}

	users = append(users, User{Name:newUser.Name, ID:uID})
	uID++
	c.IndentedJSON(http.StatusCreated, users)
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
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "error parsing url param"})
	}
}

func deleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err == nil {
		for i, a := range users {
			if a.ID == id {
				// Remove the element at index i from a.
				copy(users[i:], users[i+1:]) // Shift a[i+1:] left one index.
				users = users[:len(users)-1]     // Truncate slice.
				c.IndentedJSON(http.StatusOK, users)
				return
			}
		}
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "error parsing url param"})
	}
}