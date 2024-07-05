package main

import (
	"github.com/gin-gonic/gin"
)

// User struct represents the structure of a user
type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var users []User

func main() {
	// Initialize Gin router
	router := gin.Default()

	// Define routes
	router.POST("/users", createUser)
	router.GET("/users", getUsers)
	router.GET("/users/:id", getUser)
	router.DELETE("/users/:id", deleteUser)

	// Run the server
	router.Run(":8080")
}

