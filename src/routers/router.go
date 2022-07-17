package routers

import (
	"net/http"

	"maily/go-backend/src/controllers"

	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
)

func Init(router *gin.Engine) {
	router.GET("/beep", controllers.Beep)
	router.GET("/iptest", controllers.IpAddress)

	router.GET("/test", func(c *gin.Context) {

		println(c.ClientIP())

		for k, vals := range c.Request.Header {
			log.Infof("%s", k)
			log.Infof("%s", vals)
		}

		c.IndentedJSON(http.StatusOK, c.Request.Header)
	})
}
