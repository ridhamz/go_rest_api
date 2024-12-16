package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ridhamz/go_rest_api/models"
)

func main() {
	// create the http server
	server := gin.Default()
	// create routes
	server.GET("/events", func(ctx *gin.Context) {
		events := models.GetAllEvents()
		ctx.JSON(http.StatusOK, events)
	})

	server.POST("/events", func(ctx *gin.Context) {
		var event models.Event
		err := ctx.ShouldBindJSON(&event)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "Something went wrong!"})
			return
		}
		event.Save()
		ctx.JSON(http.StatusOK, gin.H{"message": "Event added successfully!"})
	})

	// start the server
	server.Run(":4000")
}
