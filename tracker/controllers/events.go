package controllers

import (
	"fmt"
	"local/tracker/db"
	"local/tracker/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type EventController struct{}

func (e EventController) CustomEvent(c *gin.Context) {
	var event models.Event
	if err := c.ShouldBindQuery(&event); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	var ua models.UserAgent
	ua.New(c.GetHeader("User-Agent"))
	exists := db.WriteUA(ua)
	fmt.Println("User-Agent", exists)

	event.ServerTimestamp = int(time.Now().Unix())
	event.Origin = c.GetHeader("Origin")
	event.Referer = c.GetHeader("Referer")
	event.UserAgentHash = ua.UA_Hash

	db.WriteEvent(event)
	c.Status(200)
}

func (e EventController) PageViewEvent(c *gin.Context) {
	var event models.Event
	if err := c.ShouldBindQuery(&event); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	var ua models.UserAgent
	ua.New(c.GetHeader("User-Agent"))
	exists := db.WriteUA(ua)
	fmt.Println("User-Agent", exists)

	event.Event = "pageview"
	event.ServerTimestamp = int(time.Now().Unix())
	event.Origin = c.GetHeader("Origin")
	event.Referer = c.GetHeader("Referer")
	event.UserAgentHash = ua.UA_Hash

	db.WriteEvent(event)
	c.Status(200)
}

func (e EventController) ClickEvent(c *gin.Context) {
	var event models.Event
	if err := c.ShouldBindQuery(&event); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	var ua models.UserAgent
	ua.New(c.GetHeader("User-Agent"))
	exists := db.WriteUA(ua)
	fmt.Println("User-Agent", exists)

	event.Event = "click"
	event.ServerTimestamp = int(time.Now().Unix())
	event.Origin = c.GetHeader("Origin")
	event.Referer = c.GetHeader("Referer")
	event.UserAgentHash = ua.UA_Hash

	db.WriteEvent(event)
	c.Status(200)
}
