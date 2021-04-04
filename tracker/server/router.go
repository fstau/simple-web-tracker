package server

import (
	"database/sql"
	"local/tracker/controllers"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var origins = []string{
	"http://localhost:8000",
	"http://localhost:8080",
}

func NewRouter(db *sql.DB) *gin.Engine {
	router := gin.Default()
	// router.Use(gin.Logger())
	// router.Use(gin.Recovery())

	router.Use(cors.New(cors.Config{
		AllowOrigins:     origins,
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.GET("/status", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "up",
		})
	})

	v1 := router.Group("v1")
	{
		trackingGroup := v1.Group("tracking")
		{
			ec := new(controllers.EventController)
			trackingGroup.GET("/custom", ec.CustomEvent)
			trackingGroup.GET("/pageview", ec.PageViewEvent)
			trackingGroup.GET("/click", ec.ClickEvent)
		}
	}
	return router

}
