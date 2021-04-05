package controllers

import (
	"github.com/gin-gonic/gin"
	"local/tracker/db"
	"local/tracker/models"
	"local/tracker/util"
	"net/http"
)

type UserController struct{}

func (u UserController) NewUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindQuery(&user); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	user.SetUA(c.GetHeader("User-Agent"))
	user.IPAddr = c.ClientIP()
	user.ServerTimestamp = util.GetTimeUnixMicro()

	db.WriteUser(user)
	c.Status(200)
}
