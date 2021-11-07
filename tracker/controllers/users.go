package controllers

import (
	"local/tracker/db"
	"local/tracker/models"
	"local/tracker/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (u UserController) PostUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	user.SetUA(c.GetHeader("User-Agent"))
	user.IPAddr = c.ClientIP()
	user.ServerTimestamp = util.GetTimeUnixMicro()

	db.WriteUser(user)
	c.Status(200)
}
