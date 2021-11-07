package controllers

import (
	"local/tracker/db"
	"local/tracker/models"
	"local/tracker/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SessionController struct{}

func (u SessionController) PostSession(c *gin.Context) {
	var session models.Session
	if err := c.ShouldBindJSON(&session); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	session.SetUA(c.GetHeader("User-Agent"))
	session.IPAddr = c.ClientIP()
	session.ServerTimestamp = util.GetTimeUnixMicro()

	db.WriteSession(session)
	c.Status(200)
}
