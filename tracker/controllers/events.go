package controllers

import (
	"local/tracker/db"
	"local/tracker/models"
	"local/tracker/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

type EventController struct{}

func (e EventController) PostEvent(c *gin.Context) {
	var event models.Event
	if err := c.ShouldBindJSON(&event); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	event.ServerTimestamp = util.GetTimeUnixMicro()
	event.Origin = c.GetHeader("Origin")
	event.Referer = c.GetHeader("Referer")

	db.WriteEvent(event)
	c.Status(200)
}
