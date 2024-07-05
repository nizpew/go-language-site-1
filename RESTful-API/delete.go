package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func deleteUser(c *gin.Context) {
    id := c.Param("id")
    for i, user := range users {
        if user.ID == id {
            users = append(users[:i], users[i+1:]...)
            c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
            return
        }
    }
    c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
}

