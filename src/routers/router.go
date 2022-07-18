package routers

import (
	"net/http"

	"maily/go-backend/src/controllers"
	"maily/go-backend/src/middlewares"

	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
)

func Init(router *gin.Engine) {
	// Suffix api
	public := router.Group("/api")

	// GET
	public.GET("/beep", controllers.Beep)
	public.GET("/iptest", controllers.IpAddress)

	public.GET("/test", func(c *gin.Context) {

		println(c.ClientIP())

		for k, vals := range c.Request.Header {
			log.Infof("%s", k)
			log.Infof("%s", vals)
		}

		c.IndentedJSON(http.StatusOK, c.Request.Header)
	})

	// POST
	public.POST("/register", controllers.Register)
	public.POST("/login", controllers.Login)

	// Suffix admin
	protected := router.Group("/api/admin")
	protected.Use(middlewares.JwtAuthMiddleware())

	// GET
	protected.GET("/user", controllers.CurrentUser)
}
