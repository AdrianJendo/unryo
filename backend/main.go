package main

import (
	"fmt"
	"net/http"
	"sort"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	ID     int  `json:"id"`
	Name  string  `json:"name"`
}

type UserRequestBody struct {
	Name  string  `json:"name"`
}

var db *gorm.DB

func initDB() {
	// Define the database connection string.
	// Replace with your actual PostgreSQL database information.
	dsn := "host=localhost user=postgres password=postgres dbname=unryo port=5432 sslmode=disable"

	// Open a connection to the database using GORM.
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// Automatically create the "user" table if it doesn't exist.
	db.AutoMigrate(&User{})

	fmt.Println("Connected to the database")
}

func main() {
	// DB
	initDB()

	// Server
	router := gin.Default()

	// https://stackoverflow.com/questions/72155224/golang-gin-middleware-with-cors-not-working
	router.Use(cors.Default())

	g := router.Group("/api/users")
	
	g.GET("", getUsers)
	// g.GET(":id", getUserByID)
	g.POST("", postUsers)
	g.DELETE(":id", deleteUser)
	g.PUT(":id", editUser)

	router.Run(":5001")
}

func getUsers(c *gin.Context) {
	var users []User
	db.Find(&users)
	sort.SliceStable(users, func(i,j int) bool {
		return users[i].ID < users[j].ID
	})
	c.IndentedJSON(http.StatusOK, users)
}

func postUsers(c *gin.Context) {
	var newUser UserRequestBody

	// Call BindJSON to bind the received JSON to
	if err := c.BindJSON(&newUser); err != nil {
			return
	}

	user := User{Name:newUser.Name}
	result := db.Create(&user)
	if result.Error != nil {
		panic(result.Error)
	}
	getUsers(c)
}

func deleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err == nil {
		result := db.Delete(&User{}, id)
		if result.Error != nil {
			panic(result.Error)
		}
		if result.RowsAffected > 0 {
			getUsers(c)
		} else {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
		}
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "error parsing url param"})
	}
}

func editUser(c *gin.Context) {
	var updatedUser UserRequestBody

	// Call BindJSON to bind the received JSON to
	if err := c.BindJSON(&updatedUser); err != nil {
			return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err == nil {
		var user User
		db.First(&user, id)
		user.Name = updatedUser.Name
		result := db.Save(&user)
		if result.Error != nil {
			panic(result.Error)
		}
		if result.RowsAffected > 0 {
			getUsers(c)
		} else {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
		}
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "error parsing url param"})
	}
}