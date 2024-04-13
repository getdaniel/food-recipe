package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

func main() {
    // Initialize a Gin router
    r := gin.Default()

    // Define a route
    r.GET("/api/hello", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"message": "Hello from Gin backend!"})
    })

    // Run the Gin server
    r.Run(":8080")
}