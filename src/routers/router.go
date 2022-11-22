package routers

import (
	"maily/go-backend/src/controllers"
	"maily/go-backend/src/middlewares"

	"github.com/gin-gonic/gin"
)

func Init(router *gin.Engine) {
	// Suffix api
	public := router.Group("/api")

	// GET
	public.GET("/beep/:trackingId", controllers.Beep)
	public.GET("/iptest", controllers.IpAddress)

	// POST
	public.POST("/register", controllers.Register)
	public.POST("/login", controllers.Login)
	public.POST("/generate", controllers.Generate)

	// Suffix admin
	protected := router.Group("/api/admin")
	protected.Use(middlewares.JwtAuthMiddleware())

	// GET
	protected.GET("/user", controllers.CurrentUser)
}
