package main

import (
	"net/http"
	"rest-api-go-gin/internal/database"

	"github.com/gin-gonic/gin"
)

func (app *application) createEvent(c *gin.Context) {
	var events database.Event

	err := c.ShouldBindJSON(&events)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = app.models.Event.Insert(&events)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, events)

}
