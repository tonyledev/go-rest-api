package main

import (
	"go/rest/api/db"
	"go/rest/api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	server.Run(":8080") // localhost:8080
}

func getEvents(ctx *gin.Context) {
	events := models.GetAllEvents()
	ctx.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
	}

	event.ID = 1
	event.UserID = 1

	event.Save()
	context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
}
